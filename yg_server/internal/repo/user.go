package repo

import (
	"context"
	"yug_server/internal/data/user/model"
	"yug_server/internal/dto"
)

type UserRepo interface {
	Register(ctx context.Context, dto *dto.RegisterRequest) error
	Login(ctx context.Context, username, password string) (*model.User, error)
	// GetUserSessionList(ctx context.Context, userID uint64) ([]*model.User, error)
	// AddFriend(ctx context.Context, userID uint64, friendID uint64) error
	// GetFriends(ctx context.Context, userID uint64) ([]*model.User, error)
}
