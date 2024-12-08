package repo

import (
	"context"
	"yug_server/internal/data/chat/model"
)

type ChatRepo interface {
	StoreOfflineMessage(ctx context.Context, msg model.ChatMsg) error
	StoreOnlineMessage(ctx context.Context, msg model.ChatMsg) error
	GetOfflineMessages(ctx context.Context, userID uint64) ([]model.ChatMsg, error)
	UpdateMessageStatus(ctx context.Context, messageID string, status string) error
}
