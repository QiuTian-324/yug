package user

import (
	"context"
	"errors"
	"yug_server/internal/data/user/model"
	"yug_server/internal/dto"
	"yug_server/internal/repo"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRepo(db *gorm.DB, logger *zap.Logger) repo.UserRepo {
	return &userRepo{db: db, logger: logger}
}

func (u *userRepo) AddFriend(ctx context.Context, userID uint64, friendID uint64) error {

	var friend = &model.Friend{
		UserID:   userID,
		FriendID: friendID,
		Status:   1,
	}

	if err := u.db.Create(&friend).Error; err != nil {
		u.logger.Error("添加好友失败", zap.Error(err))
		return err
	}

	return nil
}

// 检查用户是否存在
func (u *userRepo) IsUserExist(ctx context.Context, userID uint64) (bool, error) {
	var user model.User
	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		u.logger.Error("用户不存在", zap.Error(err))
		return false, err
	}
	return true, nil
}

// 检查用户是否为好友
func (u *userRepo) IsFriend(ctx context.Context, userID uint64, friendID uint64) (bool, error) {
	var friend model.Friend
	if err := u.db.Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).First(&friend).Error; err != nil {
		u.logger.Error("好友关系不存在", zap.Error(err))
		return false, err
	}
	return true, nil
}

// 注册用户
func (u *userRepo) Register(ctx context.Context, dto *dto.RegisterRequest) error {
	user := &model.User{
		Username: dto.Username,
		Password: dto.Password,
		Email:    dto.Email,
		Phone:    dto.Phone,
	}

	// 检查用户是否存在
	if err := u.db.Where("username = ?", dto.Username).First(user).Error; err != nil {
		u.logger.Error("用户已存在", zap.Error(err))
		return err
	}

	err := u.db.Create(user).Error
	if err != nil {
		u.logger.Error("注册失败", zap.Error(err))
		return err
	}

	return nil
}

// 登录
func (u *userRepo) Login(ctx context.Context, username, password string) (*model.User, error) {
	var user model.User
	err := u.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("登录失败，用户不存在或密码错误", zap.Error(err))
		return nil, err
	}
	return &user, err
}

// 通过用户名查询用户
func (u *userRepo) QueryUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := u.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("查询用户失败，用户不存在", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

// 通过邮箱查询用户
func (u *userRepo) QueryUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("查询用户失败，用户不存在", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

// 通过手机号查询用户
func (u *userRepo) QueryUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := u.db.Where("phone = ?", phone).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("查询用户失败，用户不存在", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

// 获取好友列表
func (u *userRepo) GetFriends(ctx context.Context, userID uint64) ([]model.User, error) {
	var user model.User

	// 使用 Preload 预加载 Friends 关系
	if err := u.db.Preload("Friends").First(&user, userID).Error; err != nil {
		u.logger.Error("获取好友列表失败", zap.Error(err))
		return nil, err
	}

	// 返回预加载的好友列表
	return user.Friends, nil
}
