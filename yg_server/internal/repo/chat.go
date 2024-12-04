package repo

import (
	"context"
	"yug_server/internal/data/chat/model"
)

type ChatRepo interface {
	StoreOfflineMessage(ctx context.Context, msg model.ChatMsg) error
	StoreOnlineMessage(ctx context.Context, msg model.ChatMsg) error
}
