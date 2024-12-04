package model

import "yug_server/internal/data"

// UserSetting 表的结构定义
type UserSetting struct {
	data.BaseModel
	UserID               int64  `gorm:"type:bigint;not null" json:"user_id"`                 // 用户 ID
	ReceiveNotifications int    `gorm:"type:tinyint;default:1" json:"receive_notifications"` // 是否接收通知（1: 接收, 0: 不接收）
	AllowStrangers       int    `gorm:"type:tinyint;default:0" json:"allow_strangers"`       // 是否允许陌生人添加好友（1: 允许, 0: 禁止）
	Theme                string `gorm:"type:varchar(50);default:'light'" json:"theme"`       // 界面主题（light/dark）
}

// 表名
func (UserSetting) TableName() string {
	return "user_settings"
}
