package repo

import "github.com/gin-gonic/gin"

type FileRepo interface {
	Upload(ctx *gin.Context, fileName string) error
	Download(ctx *gin.Context, fileName string) ([]byte, error)
	Delete(ctx *gin.Context, fileName string) error
}
