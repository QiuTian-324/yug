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

// SetOnline 设置用户在线状态
func (u *User) SetOnline(db *gorm.DB) error {
	return db.Model(u).Where("id = ?", u.ID).Update("online", u.Online).Error
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
	err := db.
		Where("username = ?", username).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 校验是否是否为好友
func IsFriend(db *gorm.DB, userID, friendID uint) (bool, error) {
	var friend Friend
	err := db.
		Where("(user_id = ? AND friend_id = ?) OR (friend_id AND user_id)", userID, friendID, userID, friendID).
		First(&friend).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

// 获取好友列表
func GetFriends(db *gorm.DB, userID uint) ([]User, error) {
	var friends []Friend
	if err := db.
		Preload("Friend").
		Where("user_id = ?", userID).
		Find(&friends).
		Error; err != nil {
		return nil, err
	}

	var reverseFriends []Friend
	if err := db.
		Preload("User").
		Where("friend_id = ?", userID).
		Find(&reverseFriends).
		Error; err != nil {
		return nil, err
	}

	// 合并两个结果
	allFriendsMap := make(map[uint]User)
	for _, friend := range friends {
		allFriendsMap[friend.FriendID] = friend.Friend
	}

	for _, friend := range reverseFriends {
		allFriendsMap[friend.UserID] = friend.User
	}

	// 将 map 转换为切片
	var friendUsers []User
	for _, user := range allFriendsMap {
		friendUsers = append(friendUsers, user)
	}

	return friendUsers, nil
}
