package handlers

import (
	"yug_server/global"
	"yug_server/internal/libs"
	"yug_server/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type FileHandler struct {
	fc     *services.FileUseCase
	rds    *redis.Client
	logger *zap.Logger
	fsoCfg *global.FileStorageConfig
}

func NewFileHandler(fc *services.FileUseCase, rds *redis.Client, logger *zap.Logger) *FileHandler {
	fsoCfg := new(global.FileStorageConfig)
	return &FileHandler{fc: fc, rds: rds, logger: logger, fsoCfg: fsoCfg}
}

func (h *FileHandler) Upload(ctx *gin.Context) {
	// 获取多文件
	files, err := ctx.MultipartForm()
	if err != nil {
		h.logger.Error("获取文件失败", zap.Error(err))
		libs.Failed(ctx, err.Error(), nil)
		return
	}

	err = h.fc.UploadFile(ctx, services.UploaderType(h.fsoCfg.Type), h.fsoCfg.Local, h.fsoCfg.Minio, files)
	if err != nil {
		h.logger.Error("上传失败", zap.Error(err))
		libs.Failed(ctx, err.Error(), nil)
		return
	}

	libs.OK(ctx, "上传成功", nil)
}
