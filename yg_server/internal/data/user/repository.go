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

func (u *userRepo) Login(ctx context.Context, username, password string) (*model.User, error) {
	var user model.User
	err := u.db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("登录失败，用户不存在或密码错误", zap.Error(err))
		return nil, err
	}
	return &user, err
}

// // GetAll 获取所有用户
// func GetAll(db *gorm.DB) ([]User, error) {
// 	var users []User
// 	err := db.Where("is_deleted = ?", false).Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// // GetByID 根据 ID 获取用户
// func GetByID(db *gorm.DB, id uint64) (*User, error) {
// 	var user User
// 	err := db.First(&user, "id = ? AND is_deleted = ?", id, false).Error
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, nil
// 	}
// 	return &user, err
// }

// // Update 更新用户信息
// func (u *User) Update(db *gorm.DB) error {
// 	u.UpdatedAt = time.Now()
// 	return db.Save(u).Error
// }

// GetByUsername 根据用户名获取用户
// func (u *userRepo) GetByUsername(username string) (*User, error) {
// 	var user User
// 	err := u.db.First(&user, "username = ? AND is_deleted = ?", username, false).Error
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		pkg.Error("用户不存在", err)
// 		return nil, err
// 	}
// 	return &user, err
// }

// // SetOnline 设置用户在线状态
// func (u *User) SetOnline(db *gorm.DB) error {
// 	return db.Model(u).Where("id = ?", u.ID).Update("online", u.Online).Error
// }

// // 获取用户及其好友
// func GetUserWithFriends(db *gorm.DB, userID int64) (*User, error) {
// 	var user User
// 	if err := db.Preload("Friends").First(&user, userID).Error; err != nil {
// 		pkg.Error("获取用户失败", err)
// 		return nil, err
// 	}
// 	return &user, nil
// }

// // 添加好友
// func AddFriend(db *gorm.DB, userID, friendID uint64) error {
// 	friend := Friend{
// 		UserID:   userID,
// 		FriendID: friendID,
// 		Status:   0, // 待接受
// 	}
// 	if err := db.Create(&friend).Error; err != nil {
// 		pkg.Error("添加好友失败", err)
// 		return err
// 	}
// 	return nil
// }

// func FindUserByUserName(db *gorm.DB, username string) (*User, error) {
// 	var user User
// 	err := db.
// 		Where("username = ?", username).
// 		First(&user).
// 		Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// // 校验是否为好友
// func IsFriend(db *gorm.DB, userID, friendID uint64) (bool, error) {
// 	var friend Friend
// 	err := db.
// 		Where("(user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)", userID, friendID, friendID, userID).
// 		First(&friend).
// 		Error
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return false, nil
// 	}
// 	return err == nil, err
// }

// // 获取好友列表
// func GetFriends(db *gorm.DB, userID uint64) ([]User, error) {
// 	var friends []Friend
// 	if err := db.
// 		Preload("Friend").
// 		Where("user_id = ?", userID).
// 		Find(&friends).
// 		Error; err != nil {
// 		return nil, err
// 	}

// 	var reverseFriends []Friend
// 	if err := db.
// 		Preload("User").
// 		Where("friend_id = ?", userID).
// 		Find(&reverseFriends).
// 		Error; err != nil {
// 		return nil, err
// 	}

// 	// 合并两个结果
// 	allFriendsMap := make(map[uint64]User)
// 	for _, friend := range friends {
// 		allFriendsMap[friend.FriendID] = friend.Friend
// 	}

// 	for _, friend := range reverseFriends {
// 		allFriendsMap[friend.UserID] = friend.User
// 	}

// 	// 将 map 转换为切片
// 	var friendUsers []User
// 	for _, user := range allFriendsMap {
// 		friendUsers = append(friendUsers, user)
// 	}

// 	return friendUsers, nil
// }
