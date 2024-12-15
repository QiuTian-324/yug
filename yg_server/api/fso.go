package api

import (
	"yug_server/internal/handlers"
	"yug_server/pkg"

	"github.com/gin-gonic/gin"
)

func RegisterFsoRoutes(router *gin.RouterGroup) {
	fileHandler := pkg.GetHandler("fileHandler").(*handlers.FileHandler)
	fileGroup := router.Group("/fso")
	fileGroup.POST("/upload", fileHandler.Upload)
}
