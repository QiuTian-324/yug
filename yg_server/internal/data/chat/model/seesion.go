package model

import (
	"time"
	"yug_server/internal/data"
)

type Session struct {
	data.BaseModel
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	FriendID  uint      `gorm:"index;not null" json:"friend_id"`
	UnreadNum int       `gorm:"default:0" json:"unread_num"`
	LastMsg   string    `gorm:"type:varchar(255);not null" json:"last_msg"`
	LastMsgAt time.Time `gorm:"type:datetime;not null" json:"last_msg_at"`
}

func (Session) TableName() string {
	return "chat_session"
}
