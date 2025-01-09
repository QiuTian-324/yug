package handlers

import (
	"context"
	"fmt"
	"time"
	"yug_server/global"
	"yug_server/internal/data/chat/model"
	"yug_server/internal/libs"
	"yug_server/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type ChatHandler struct {
	wc     *services.WsUseCase
	rds    *redis.Client
	logger *zap.Logger
}

func NewChatHandler(wc *services.WsUseCase, rds *redis.Client, logger *zap.Logger) *ChatHandler {
	return &ChatHandler{wc: wc, rds: rds, logger: logger}
}

func (h *ChatHandler) Ws(ctx *gin.Context) {
	userID := ctx.MustGet("id").(uint64)
	conn, err := global.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		h.logger.Error("websocket upgrade error", zap.Error(err))
		return
	}

	defer func() {
		conn.Close()
		h.wc.RemoveConnection(userID)
	}()

	// 添加连接
	h.wc.AddConnection(userID, conn)

	// 启动心跳检测
	go h.wc.Heartbeat(conn, userID)

	for {
		_, messageData, err := conn.ReadMessage()
		if err != nil {
			return
		}

		h.logger.Info("receive message", zap.String("message", string(messageData)))

		err = h.wc.SendMessage(conn, messageData)
		if err != nil {
			return
		}
	}
}

// 获取会话列表
func (h *ChatHandler) GetSessionList(ctx *gin.Context) {
	userID := ctx.MustGet("id").(uint64)

	friendSetKey := fmt.Sprintf("user:%d:friends", userID)

	// 获取用户的所有好友ID
	friendIDs, err := h.rds.SMembers(context.Background(), friendSetKey).Result()

	if err != nil {
		h.logger.Error("获取好友集合失败", zap.Error(err))
		libs.Failed(ctx, "获取好友集合失败", nil)
		return
	}

	type list struct {
		UserID    uint64 `json:"user_id"`
		FriendID  uint64 `json:"friend_id"`
		Nickname  string `json:"nickname"`
		Avatar    string `json:"avatar"`
		UnreadNum int    `json:"unread_num"`
		LastMsg   string `json:"last_msg"`
		LastMsgAt string `json:"last_msg_at"`
	}

	var sessions = make([]list, 0)
	for _, friendIDStr := range friendIDs {
		friendID := cast.ToUint(friendIDStr)
		redisKey := fmt.Sprintf("session:%d:%d", userID, friendID)

		// 从 Redis 哈希表中获取会话数据
		result, err := h.rds.HGetAll(context.Background(), redisKey).Result()
		if err != nil {
			h.logger.Error("从Redis获取会话数据失败", zap.Error(err))
			continue
		}

		if len(result) == 0 {
			continue
		}
		lastMsgAt, _ := time.Parse(time.RFC3339, result["LastMsgAt"])
		// 定义一个返回响应
		var session list
		session.UserID = cast.ToUint64(result["UserID"])
		session.FriendID = cast.ToUint64(result["FriendID"])
		session.Nickname = result["FriendName"]
		session.Avatar = result["FriendAvatar"]
		session.UnreadNum = cast.ToInt(result["UnreadNum"])
		session.LastMsg = result["LastMsg"]
		session.LastMsgAt = lastMsgAt.Format(time.RFC3339)

		sessions = append(sessions, session)
	}

	libs.OK(ctx, "获取会话列表成功", map[string]interface{}{
		"list": sessions,
	})
}

// 获取聊天记录
func (h *ChatHandler) GetChatRecord(ctx *gin.Context) {
	userID := ctx.MustGet("id").(uint64)
	friendID := ctx.Query("friend_id")

	// 获取chat_msg表中的数据
	chatMsgs := []model.ChatMsg{}
	global.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", userID, friendID, friendID, userID).Find(&chatMsgs)

	libs.OK(ctx, "获取聊天记录成功", map[string]interface{}{
		"list": chatMsgs,
	})
}
