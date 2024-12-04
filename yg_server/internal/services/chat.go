// internal/services/chat_service.go
package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"yug_server/global"
	"yug_server/internal/data/chat/model"
	"yug_server/internal/dto"
	"yug_server/internal/libs"
	"yug_server/internal/repo"
	"yug_server/utils"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type WsUseCase struct {
	repo   repo.ChatRepo
	rds    *redis.Client
	logger *zap.Logger
}

func NewWsUseCase(repo repo.ChatRepo, rds *redis.Client, logger *zap.Logger) *WsUseCase {
	return &WsUseCase{repo: repo, rds: rds, logger: logger}
}

// 发送消息
func (s *WsUseCase) SendMessage(conn *websocket.Conn, messageData []byte) error {
	var msg dto.Message
	err := json.Unmarshal(messageData, &msg)
	if err != nil {
		s.logger.Error("unmarshal message error", zap.Error(err))
		return err
	}

	// 根据消息类型处理内容
	switch msg.Type {
	case dto.TextMessageType:
		msg.URL = ""
		msg.FileName = ""
		msg.Size = ""
		
		if utils.CheckSensitiveWords(msg.Content) {
			s.logger.Warn("消息包含敏感词")
			return errors.New("消息包含敏感词")
		}
	case dto.ImageMessageType:
		msg.Content = "[图片消息]"
	case dto.FileMessageType:
		msg.Content = "[文件消息]"
	case dto.VideoMessageType:
		msg.Content = "[视频消息]"
	default:
		msg.Content = "[未知消息]"
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
	libs.Mu.Lock()
	libs.UserConnectionsMap[cast.ToString(userID)] = conn
	libs.Mu.Unlock()

	redisKey := fmt.Sprintf("%s:%s", global.ChatRedisOnline, cast.ToString(userID))
	err := s.rds.Set(context.Background(), redisKey, "true", time.Hour*24).Err()
	if err != nil {
		s.logger.Error("设置用户在线状态失败", zap.Error(err))
	}
}

// 移除连接
func (s *WsUseCase) RemoveConnection(userID uint64) {
	libs.Mu.Lock()
	delete(libs.UserConnectionsMap, cast.ToString(userID))
	libs.Mu.Unlock()

	// 从 Redis 中删除用户在线状态
	redisKey := fmt.Sprintf("%s:%s", global.ChatRedisOnline, cast.ToString(userID))
	s.rds.Del(context.Background(), redisKey)
}

// 获取接收者的连接
func (s *WsUseCase) GetReceiverConnection(receiverID uint64) *websocket.Conn {
	libs.Mu.Lock()
	defer libs.Mu.Unlock()
	return libs.UserConnectionsMap[cast.ToString(receiverID)]
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
func (s *WsUseCase) IsUserOnline(userID uint64) bool {
	redisKey := fmt.Sprintf("%s:%s", global.ChatRedisOnline, cast.ToString(userID))
	onlineStatus, err := s.rds.Get(context.Background(), redisKey).Result()
	if err != nil {
		s.logger.Error("获取用户在线状态失败", zap.Error(err))
		return false
	}
	s.logger.Info("用户在线状态", zap.String("onlineStatus", onlineStatus))

	return onlineStatus == "true"
}

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
