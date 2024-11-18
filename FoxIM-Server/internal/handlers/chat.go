package handlers

import (
	"encoding/json"
	"gin_template/global"
	"gin_template/internal/dto"
	"gin_template/pkg"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Ws(ctx *gin.Context) {
	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := global.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		pkg.Error("websocket upgrade error", err)
		return
	}

	defer conn.Close()

	for {
		messageType, messageData, err := conn.ReadMessage()
		if err != nil {
			return
		}

		pkg.Info("receive message", zap.String("message", string(messageData)))
		pkg.Info("message type", zap.Int("type", messageType))

		// 解析消息
		var msg dto.Message
		err = json.Unmarshal(messageData, &msg)
		if err != nil {
			pkg.Error("unmarshal message error", err)
			continue
		}

		// 根据消息类型处理消息
		switch msg.Type {
		case dto.TextMessageType:
			// 构造响应消息
			msgResponse := dto.NewMessage(
				msg.Type,
				msg.Content,
				msg.URL,
				msg.FileName,
				msg.SenderID,
				msg.ReceiverID,
				msg.GroupID,
			)
			err = conn.WriteJSON(msgResponse)
			if err != nil {
				pkg.Error("消息发送失败", err)
				return
			}
		default:
			conn.WriteJSON("消息发送失败")
		}

	}

}
