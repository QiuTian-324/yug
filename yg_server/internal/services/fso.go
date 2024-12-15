package services

import (
	"fmt"
	"mime/multipart"
	"yug_server/global"
	"yug_server/internal/repo"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type FileUseCase struct {
	repo   repo.FileRepo
	rds    *redis.Client
	logger *zap.Logger
}

func NewFileUseCase(repo repo.FileRepo, rds *redis.Client, logger *zap.Logger) *FileUseCase {
	return &FileUseCase{
		repo:   repo,
		rds:    rds,
		logger: logger,
	}
}

type Uploader interface {
	Upload(ctx *gin.Context, files *multipart.Form) error
	Download(ctx *gin.Context, fileName string) (fileData []byte, err error)
	Delete(ctx *gin.Context, fileName string) error
	Preview(ctx *gin.Context, fileName string) (fileData []byte, err error)
}

type LocalUploader struct {
	config global.LocalConfig
}

func (u *LocalUploader) Upload(ctx *gin.Context, files *multipart.Form) error {
	// 实现本地上传逻辑
	return nil
}

func (u *LocalUploader) Download(ctx *gin.Context, fileName string) ([]byte, error) {
	return nil, nil
}

func (u *LocalUploader) Delete(ctx *gin.Context, fileName string) error {
	return nil
}

func (u *LocalUploader) Preview(ctx *gin.Context, fileName string) ([]byte, error) {
	return nil, nil
}

type MinioUploader struct {
	config global.MinioConfig
}

func (u *MinioUploader) Upload(ctx *gin.Context, files *multipart.Form) error {
	// 实现 Minio 上传逻辑
	return nil
}

func (u *MinioUploader) Download(ctx *gin.Context, fileName string) ([]byte, error) {
	return nil, nil
}

func (u *MinioUploader) Delete(ctx *gin.Context, fileName string) error {
	return nil
}

func (u *MinioUploader) Preview(ctx *gin.Context, fileName string) ([]byte, error) {
	return nil, nil
}

type UploaderType string

const (
	LocalUploaderType UploaderType = "local"
	MinioUploaderType UploaderType = "minio"
)

func (f *FileUseCase) UploadFile(ctx *gin.Context, uploaderType UploaderType, localConfig global.LocalConfig, minioConfig global.MinioConfig, files *multipart.Form) error {
	switch uploaderType {
	case LocalUploaderType:
		uploader := &LocalUploader{config: localConfig}
		return uploader.Upload(ctx, files)
	case MinioUploaderType:
		uploader := &MinioUploader{config: minioConfig}
		return uploader.Upload(ctx, files)
	default:
		return fmt.Errorf("unsupported uploader type: %s", uploaderType)
	}
}
