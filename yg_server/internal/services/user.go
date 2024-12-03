package services

import (
	"context"

	"yug_server/internal/data/user/model"
	"yug_server/internal/dto"
	"yug_server/internal/repo"

	"go.uber.org/zap"
)

type UserUseCase struct {
	repo   repo.UserRepo
	logger *zap.Logger
}

func NewUserUseCase(repo repo.UserRepo, logger *zap.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, logger: logger}
}

func (uc *UserUseCase) Register(ctx context.Context, dto *dto.RegisterRequest) error {
	err := uc.repo.Register(ctx, dto)
	if err != nil {
		uc.logger.Error("注册失败", zap.Error(err))
		return err
	}
	return nil
}

func (uc *UserUseCase) Login(ctx context.Context, username, password string) (*model.User, error) {

	// 查询用户
	userInfo, err := uc.repo.Login(ctx, username, password)
	if err != nil {
		uc.logger.Error("用户不存在", zap.Error(err))
		return nil, err
	}

	return userInfo, nil
}

// // 第三方登录

// func LoginByThirdParty(db *gorm.DB, username string) (*user.User, error) {
// 	// 查询用户
// 	userInfo, err := user.GetByUsername(db, username)
// 	if err != nil {
// 		pkg.Error("用户不存在", err)
// 		return nil, errors.New("当前账号未注册")
// 	}

// 	// 更改当前用户的在线状态
// 	user := new(user.User)
// 	user.ID = userInfo.ID
// 	user.Online = 1
// 	err = user.SetOnline(db)
// 	if err != nil {
// 		pkg.Error("设置在线状态失败", err)
// 		return nil, errors.New("登陆失败")
// 	}

// 	return userInfo, nil
// }

// func (uc *UserUseCase) AddFriend(ctx context.Context, userID uint64, friendID uint64) error {

// 	// 根据用户名查找好友
// 	err := uc.repo.AddFriend(ctx, userID, friendID)
// 	if err != nil {
// 		uc.logger.Error("添加好友失败", zap.Error(err))
// 		return err
// 	}
// 	// isFriend, _ := uc.repo.IsFriend(ctx, userID, friendInfo.ID)

// 	// if isFriend {
// 	// 	pkg.Info("已经是好友了")
// 	// 	return errors.New("你们已经是好友了")
// 	// }

// 	// if friendInfo.ID == userID {
// 	// 	pkg.Error("不能添加自己为好友", nil)
// 	// 	return errors.New("不能添加自己为好友")
// 	// }

// 	// err = user.AddFriend(db, userID, friendInfo.ID)
// 	// if err != nil {
// 	// 	pkg.Error("添加好友失败", err)
// 	// 	return errors.New("添加好友失败")
// 	// }
// 	return nil
// }

// func (uc *UserUseCase) GetFriends(ctx context.Context, userID uint64) ([]dto.FriendListResponse, error) {
// 	friends, err := uc.repo.GetFriends(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var friendList []dto.FriendListResponse

// 	for _, friend := range friends {
// 		friendList = append(friendList, dto.FriendListResponse{
// 			UserID:    friend.ID,
// 			Username:  friend.Username,
// 			Nickname:  friend.Nickname,
// 			Email:     friend.Email,
// 			Phone:     friend.Phone,
// 			AvatarUrl: friend.AvatarUrl,
// 			Bio:       friend.Bio,
// 			Online:    friend.Online,
// 		})
// 	}

// 	return friendList, nil
// }
