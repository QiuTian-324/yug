package services

import (
	"errors"
	"gin_template/internal/data/user"
	"gin_template/internal/dto"
	"gin_template/pkg"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, dto *dto.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		pkg.Error("生成密码哈希失败", err)
		return errors.New("注册失败")
	}
	user := &user.User{
		Username: dto.Username,
		Password: string(hashedPassword),
		Email:    dto.Email,
		Phone:    dto.Phone,
	}

	err = user.Add(db)
	if err != nil {
		pkg.Error("添加用户失败", err)
		return errors.New("注册失败")
	}
	return nil
}

func Login(db *gorm.DB, username, password string) (*user.User, error) {

	// 查询用户
	user, err := user.GetByUsername(db, username)
	if err != nil {
		pkg.Error("用户不存在", err)
		return nil, errors.New("用户不存在")
	}

	return user, nil
}

func AddFriend(db *gorm.DB, userID uint, dto *dto.AddFriendRequest) error {

	// 根据用户名查找好友
	friendInfo, err := user.FindUserByUserName(db, dto.UserName)
	if err != nil {
		pkg.Error("好友不存在", err)
		return errors.New("好友不存在")
	}

	isFriend, _ := user.IsFriend(db, userID, friendInfo.ID)

	if isFriend {
		pkg.Info("已经是好友了")
		return errors.New("你们已经是好友了")
	}

	if friendInfo.ID == userID {
		pkg.Error("不能添加自己为好友", nil)
		return errors.New("不能添加自己为好友")
	}

	err = user.AddFriend(db, userID, friendInfo.ID)
	if err != nil {
		pkg.Error("添加好友失败", err)
		return errors.New("添加好友失败")
	}
	return nil
}
