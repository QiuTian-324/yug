package user

import (
	"errors"
	"gin_template/pkg"
	"time"

	"gorm.io/gorm"
)

// Add 新增用户
func (u *User) Add(db *gorm.DB) error {
	return db.Create(u).Error
}

// GetAll 获取所有用户
func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID 根据 ID 获取用户
func GetByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, "id = ? AND is_deleted = ?", id, false).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// Update 更新用户信息
func (u *User) Update(db *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return db.Save(u).Error
}

// GetByUsername 根据用户名获取用户
func GetByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.First(&user, "username = ? AND is_deleted = ?", username, false).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pkg.Error("用户不存在", err)
		return nil, err
	}
	return &user, err
}

// 获取用户及其好友
func GetUserWithFriends(db *gorm.DB, userID int64) (*User, error) {
	var user User
	if err := db.Preload("Friends").First(&user, userID).Error; err != nil {
		pkg.Error("获取用户失败", err)
		return nil, err
	}
	return &user, nil
}

// 添加好友
func AddFriend(db *gorm.DB, userID, friendID uint) error {
	friend := Friend{
		UserID:   userID,
		FriendID: friendID,
		Status:   0, // 待接受
	}
	if err := db.Create(&friend).Error; err != nil {
		pkg.Error("添加好友失败", err)
		return err
	}
	return nil
}

func FindUserByUserName(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
