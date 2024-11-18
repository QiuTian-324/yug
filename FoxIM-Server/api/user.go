package api

import (
	"gin_template/internal/handlers"
	"gin_template/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户相关的路由
func RegisterUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	userGroup.POST("/register", handlers.Register)
	userGroup.POST("/login", handlers.Login)
	// userGroup.POST("/logout", handlers.Logout)
	// userGroup.GET("/info", handlers.GetUserInfo)
	// userGroup.GET("/list", handlers.GetUserList)
	userGroup.POST("/add_friend", middleware.AuthMiddleware(), handlers.AddFriend)
}
