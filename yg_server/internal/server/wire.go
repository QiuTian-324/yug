package server

import (
	"yug_server/internal/data/user"
	"yug_server/internal/handlers"
	"yug_server/internal/services"

	"github.com/google/wire"
)

var ChatHandlerSet = wire.NewSet(
	services.NewWsService,
	handlers.NewChatHandler,
)

var UserHandlerSet = wire.NewSet(
	user.NewUserRepo,
	services.NewUserUseCase,
	handlers.NewUserHandler,
)
