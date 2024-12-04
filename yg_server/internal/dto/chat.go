package dto

import "time"

// MessageType 定义消息类型
type MessageType string

const (
	TextMessageType  MessageType = "text"
	ImageMessageType MessageType = "image"
	FileMessageType  MessageType = "file"
	VideoMessageType MessageType = "video"
)

// Message 定义通用的消息结构体
type Message struct {
	Type       MessageType `json:"type"`        // 消息类型
	Content    string      `json:"content"`     // 消息内容
	URL        string      `json:"url"`         // 图片/文件/视频的 URL
	FileName   string      `json:"file_name"`   // 文件名称
	Size       string      `json:"size"`        // 文件大小
	SenderID   string      `json:"sender_id"`   // 发送人ID
	ReceiverID string      `json:"receiver_id"` // 接收人ID
	GroupID    string      `json:"group_id"`    // 群ID
	SendTime   time.Time   `json:"send_time"`   // 发送时间
}

// NewMessage 创建新的通用消息
func NewMessage(
	messageType MessageType,
	content string,
	url string,
	fileName string,
	size string,
	senderID string,
	receiverID string,
	groupID string,
) *Message {
	return &Message{
		Type:       messageType,
		Content:    content,
		URL:        url,
		FileName:   fileName,
		Size:       size,
		SenderID:   senderID,
		ReceiverID: receiverID,
		GroupID:    groupID,
		SendTime:   time.Now(),
	}
}
