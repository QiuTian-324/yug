package api

import (
	"yug_server/internal/handlers"
	"yug_server/internal/middleware"
	"yug_server/pkg"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.RouterGroup) {
	chatHandler := pkg.GetHandler("chatHandler").(*handlers.ChatHandler)

	chatGroup := router.Group("/chat", middleware.AuthMiddleware())
	// websocket连接
	chatGroup.GET("/ws", chatHandler.Ws)
	// 获取离线消息
	// chatGroup.GET("/offline-messages", chatHandler.GetOfflineMessages)

	// 获取会话列表
	chatGroup.GET("/session_list", chatHandler.GetSessionList)
	// 获取单聊聊天记录
	chatGroup.GET("/chat_record", chatHandler.GetChatRecord)
	// 获取群聊聊天记录
	// chatGroup.GET("/group_chat_record", chatHandler.GetGroupChatRecord)
}
