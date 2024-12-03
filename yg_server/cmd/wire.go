//go:build wireinject
// +build wireinject

package cmd

import (
	"yug_server/internal/handlers"
	"yug_server/internal/server"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitializeChatHandler(db *gorm.DB, logger *zap.Logger) *handlers.ChatHandler {
	wire.Build(server.ChatHandlerSet)
	return &handlers.ChatHandler{}
}

func InitializeUserHandler(db *gorm.DB, logger *zap.Logger) *handlers.UserHandler {
	wire.Build(server.UserHandlerSet)
	return &handlers.UserHandler{}
}
