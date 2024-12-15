package libs

import (
	chatModel "yug_server/internal/data/chat/model"
	userModel "yug_server/internal/data/user/model"
	"yug_server/pkg"

	"gorm.io/gorm"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&userModel.UserSetting{},
		&userModel.User{},
		&userModel.Friend{},
		&chatModel.Session{},
		&chatModel.ChatMsg{},
	)
	if err != nil {
		return err
	}
	pkg.Info("数据库表迁移成功")
	return nil
}
