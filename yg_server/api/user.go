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
	userGroup.Use(middleware.AuthMiddleware())
	userGroup.GET("/list", handlers.GetUserSeesionList)
	userGroup.POST("/add_friend", handlers.AddFriend)
	// 获取好友列表
	userGroup.GET("/friends", handlers.GetFriends)
	userGroup.POST("/logout", handlers.Logout)
	// userGroup.GET("/info", handlers.GetUserInfo)
}
