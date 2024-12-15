package file

import (
	"yug_server/internal/repo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FileRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewFileRepo(db *gorm.DB, logger *zap.Logger) repo.FileRepo {
	return &FileRepo{db: db, log: logger}
}

func (r *FileRepo) Upload(ctx *gin.Context, fileName string) error {
	return nil
}

func (r *FileRepo) Delete(ctx *gin.Context, fileName string) error {
	return nil
}

func (r *FileRepo) Preview(ctx *gin.Context, fileName string) ([]byte, error) {
	return nil, nil
}

func (r *FileRepo) Download(ctx *gin.Context, fileName string) ([]byte, error) {
	return nil, nil
}
