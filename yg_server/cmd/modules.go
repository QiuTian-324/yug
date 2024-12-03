package cmd

import (
	"yug_server/pkg"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Module interface {
	Initialize(db *gorm.DB, logger *zap.Logger) error
}

type UserModule struct{}

func (m *UserModule) Initialize(db *gorm.DB, logger *zap.Logger) error {
	userHandler := InitializeUserHandler(db, logger)
	pkg.RegisterHandler("userHandler", userHandler)
	return nil
}

type ChatModule struct{}

func (m *ChatModule) Initialize(db *gorm.DB, logger *zap.Logger) error {
	chatHandler := InitializeChatHandler(db, logger)
	pkg.RegisterHandler("chatHandler", chatHandler)
	return nil
}
