// internal/services/chat_service.go
package services

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"
	"yug_server/global"
	"yug_server/internal/data/chat/model"
	"yug_server/internal/dto"
	"yug_server/internal/repo"
	"yug_server/utils"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type WsUseCase struct {
	repo               repo.ChatRepo
	rds                *redis.Client
	logger             *zap.Logger
	mu                 sync.Mutex
	userConnectionsMap map[string]*websocket.Conn
}

func NewWsUseCase(repo repo.ChatRepo, rds *redis.Client, logger *zap.Logger) *WsUseCase {
	return &WsUseCase{
		repo:               repo,
		rds:                rds,
		logger:             logger,
		userConnectionsMap: make(map[string]*websocket.Conn),
	}
}

// 发送消息
func (s *WsUseCase) SendMessage(conn *websocket.Conn, messageData []byte) error {
	var msg dto.Message
	err := json.Unmarshal(messageData, &msg)
	if err != nil {
		s.logger.Error("unmarshal message error", zap.Error(err))
		return err
	}

	handler := GetMessageHandler(msg.Type, s.logger)
	err = handler.HandleMessage(&msg)
	if err != nil {
		return err
	}

	msgResponse := dto.NewMessage(
		msg.Type,
		msg.Content,
		msg.URL,
		msg.Size,
		msg.FileName,
		msg.SenderID,
		msg.ReceiverID,
		msg.GroupID,
		msg.Receiver,
	)
	// 发送消息给自己
	err = conn.WriteJSON(msgResponse)
	if err != nil {
		s.logger.Error("消息发送给自己失败", zap.Error(err))
		return err
	}

	// 检查接收者是否在线
	if !s.IsUserOnline(cast.ToUint64(msg.ReceiverID)) {
		s.logger.Info("接收者不在线，存储离线消息")
		// 存储离线消息到数据库
		err = s.StoreOfflineMessage(msg)
		if err != nil {
			s.logger.Error("存储离线消息失败", zap.Error(err))
			return err
		}
		return nil
	}

	// 查找接收者的连接并发送消息
	receiverConn := s.GetReceiverConnection(cast.ToUint64(msg.ReceiverID))
	if receiverConn != nil {
		err = receiverConn.WriteJSON(msgResponse)
		if err != nil {
			s.logger.Error("消息发送给接收者失败", zap.Error(err))
			return err
		}
		s.logger.Info("消息发送给接收者成功", zap.Any("msgResponse", msgResponse))
		// 存储在线消息
		onlineMsg := model.ChatMsg{
			ContentType:    s.getContentTypeInt(msg.Type),
			Content:        msg.Content,
			SenderID:       cast.ToUint64(msg.SenderID),
			ReceiverID:     cast.ToUint64(msg.ReceiverID),
			GroupID:        cast.ToUint64(msg.GroupID),
			Url:            msg.URL,
			FileName:       msg.FileName,
			OfflineMessage: false,
		}
		s.repo.StoreOnlineMessage(context.Background(), onlineMsg)
	} else {
		s.logger.Error("接收者连接不存在")
	}

	return nil
}

// 添加连接
func (s *WsUseCase) AddConnection(userID uint64, conn *websocket.Conn) {
	s.mu.Lock()
	s.userConnectionsMap[cast.ToString(userID)] = conn
	s.mu.Unlock()

	// 将用户ID添加到在线用户集合中
	redisKey := global.ChatRedisOnline
	err := s.rds.SAdd(context.Background(), redisKey, cast.ToString(userID)).Err()
	if err != nil {
		s.logger.Error("添加用户到在线集合失败", zap.Error(err))
	}
}

// 移除连接
func (s *WsUseCase) RemoveConnection(userID uint64) {
	s.mu.Lock()
	delete(s.userConnectionsMap, cast.ToString(userID))
	s.mu.Unlock()

	// 从在线用户集合中移除用户ID
	redisKey := global.ChatRedisOnline
	s.rds.SRem(context.Background(), redisKey, cast.ToString(userID))
}

// 获取接收者的连接
func (s *WsUseCase) GetReceiverConnection(receiverID uint64) *websocket.Conn {
	s.mu.Lock()
	defer s.mu.Unlock()

	conn, exists := s.userConnectionsMap[cast.ToString(receiverID)]
	if !exists {
		s.logger.Error("接收者连接不存在")
		return nil
	}

	return conn
}

// 心跳检测
func (s *WsUseCase) Heartbeat(conn *websocket.Conn, userID uint64) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		// s.logger.Info("发送心跳检测", zap.String("ping", "ping"))
		err := conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			s.logger.Error("心跳检测失败", zap.Error(err))
			s.RemoveConnection(userID)
			return
		}
	}
}

// 判断用户是否在线
func (s *WsUseCase) IsUserOnline(userID uint64) bool {
	redisKey := global.ChatRedisOnline
	isMember, err := s.rds.SIsMember(context.Background(), redisKey, cast.ToString(userID)).Result()
	if err != nil {
		s.logger.Error("获取用户在线状态失败", zap.Error(err))
		return false
	}
	s.logger.Info("用户在线状态", zap.Bool("isOnline", isMember))

	return isMember
}

// 存储离线消息
func (s *WsUseCase) StoreOfflineMessage(msg dto.Message) error {

	offlineMsg := model.ChatMsg{
		ContentType:    s.getContentTypeInt(msg.Type),
		Content:        msg.Content,
		SenderID:       cast.ToUint64(msg.SenderID),
		ReceiverID:     cast.ToUint64(msg.ReceiverID),
		GroupID:        cast.ToUint64(msg.GroupID),
		Url:            msg.URL,
		FileName:       msg.FileName,
		OfflineMessage: true,
	}
	s.logger.Info("存储离线消息", zap.Any("offlineMsg", offlineMsg))

	// 执行插入操作
	err := s.repo.StoreOfflineMessage(context.Background(), offlineMsg)
	if err != nil {
		s.logger.Error("存储离线消息失败", zap.Error(err))
		return err
	}
	s.logger.Info("离线消息存储成功")
	return nil
}

// 获取消息类型对应的整数
func (s *WsUseCase) getContentTypeInt(msgType dto.MessageType) int {
	// 定义消息类型到整数的映射
	messageTypeToInt := map[dto.MessageType]int{
		dto.TextMessageType:  model.ContentTypeText,
		dto.ImageMessageType: model.ContentTypeImage,
		dto.FileMessageType:  model.ContentTypeFile,
		dto.VideoMessageType: model.ContentTypeVideo,
	}
	return messageTypeToInt[msgType]
}

// 定义消息处理接口
type MessageHandler interface {
	HandleMessage(msg *dto.Message) error
}

// 文本消息处理器
type TextMessageHandler struct {
	logger *zap.Logger
}

func (h *TextMessageHandler) HandleMessage(msg *dto.Message) error {
	msg.URL = ""
	msg.FileName = ""
	msg.Size = ""

	if utils.CheckSensitiveWords(msg.Content) {
		h.logger.Warn("消息包含敏感词")
		return errors.New("消息包含敏感词")
	}
	return nil
}

// 图片消息处理器
type ImageMessageHandler struct{}

func (h *ImageMessageHandler) HandleMessage(msg *dto.Message) error {
	msg.Content = "[图片消息]"
	return nil
}

// 文件消息处理器
type FileMessageHandler struct{}

func (h *FileMessageHandler) HandleMessage(msg *dto.Message) error {
	msg.Content = "[文件消息]"
	return nil
}

// 视频消息处理器
type VideoMessageHandler struct{}

func (h *VideoMessageHandler) HandleMessage(msg *dto.Message) error {
	msg.Content = "[视频消息]"
	return nil
}

// 未知消息处理器
type UnknownMessageHandler struct{}

func (h *UnknownMessageHandler) HandleMessage(msg *dto.Message) error {
	msg.Content = "[未知消息]"
	return nil
}

func GetMessageHandler(msgType dto.MessageType, logger *zap.Logger) MessageHandler {
	switch msgType {
	case dto.TextMessageType:
		return &TextMessageHandler{logger: logger}
	case dto.ImageMessageType:
		return &ImageMessageHandler{}
	case dto.FileMessageType:
		return &FileMessageHandler{}
	case dto.VideoMessageType:
		return &VideoMessageHandler{}
	default:
		return &UnknownMessageHandler{}
	}
}
