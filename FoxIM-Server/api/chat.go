package api

import (
	"gin_template/internal/handlers"
	"gin_template/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.RouterGroup) {
	chatGroup := router.Group("/chat", middleware.AuthMiddleware())
	// websocket连接
	chatGroup.GET("/ws", handlers.Ws)
}
