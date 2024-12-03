package chat

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ChatMsgRepositoryInterface interface {
	Insert(msg *ChatMsg) error
	GetById(id uint64) (*ChatMsg, error)
}

type ChatMsgRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewChatMsgRepository(db *gorm.DB, logger *zap.Logger) ChatMsgRepositoryInterface {
	return &ChatMsgRepository{DB: db, Log: logger}
}

func (r *ChatMsgRepository) Insert(msg *ChatMsg) error {
	return r.DB.Create(msg).Error
}

func (r *ChatMsgRepository) GetById(id uint64) (*ChatMsg, error) {
	var msg ChatMsg
	if err := r.DB.Where("id = ?", id).First(&msg).Error; err != nil {
		return nil, err
	}
	return &msg, nil
}
