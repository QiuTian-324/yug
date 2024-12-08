package chat

import (
	"context"
	"yug_server/internal/data/chat/model"
	"yug_server/internal/repo"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type chatRepo struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *zap.Logger
}

func NewChatRepo(db *gorm.DB, rds *redis.Client, logger *zap.Logger) repo.ChatRepo {
	return &chatRepo{db: db, rds: rds, logger: logger}
}

// 存储离线消息
func (r *chatRepo) StoreOfflineMessage(ctx context.Context, msg model.ChatMsg) error {
	return r.db.Create(&msg).Error
}

// 存储在线消息
func (r *chatRepo) StoreOnlineMessage(ctx context.Context, msg model.ChatMsg) error {
	return r.db.Create(&msg).Error
}

// 获取离线消息
func (r *chatRepo) GetOfflineMessages(ctx context.Context, userID uint64) ([]model.ChatMsg, error) {
	var messages []model.ChatMsg
	err := r.db.Where("receiver_id = ? AND status = ?", userID, "1").Find(&messages).Error
	return messages, err
}

// 更新消息状态
func (r *chatRepo) UpdateMessageStatus(ctx context.Context, messageID string, status string) error {
	return r.db.Model(&model.ChatMsg{}).Where("id = ?", messageID).Update("status", status).Error
}
