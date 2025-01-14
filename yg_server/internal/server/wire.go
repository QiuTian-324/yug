package server

import (
	"yug_server/internal/data/chat"
	file "yug_server/internal/data/fso"
	"yug_server/internal/data/user"
	"yug_server/internal/handlers"
	"yug_server/internal/services"

	"github.com/google/wire"
)

var ChatHandlerSet = wire.NewSet(
	chat.NewChatRepo,
	services.NewWsUseCase,
	handlers.NewChatHandler,
)

var UserHandlerSet = wire.NewSet(
	user.NewUserRepo,
	services.NewUserUseCase,
	handlers.NewUserHandler,
)

var FileHandlerSet = wire.NewSet(
	file.NewFileRepo,
	services.NewFileUseCase,
	handlers.NewFileHandler,
)
