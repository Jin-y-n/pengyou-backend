package entity

import (
	"time"

	"gorm.io/gorm"
)

type MessageSend struct {
	ID          uint `gorm:"primaryKey;autoIncrement" json:"id"`
	SenderId    uint
	RecipientId uint
	Type        int
	Content     string `json:"content"`

	SentAt   time.Time
	DeleteAt gorm.DeletedAt
	IsRead   int
}

func (table *MessageSend) TableName() string {
	return "message_send"
}

// NewMessageSend creates a new instance of MessageSend and fills it with data.
func NewMessageSend(senderId, recipientId uint, messageType int, content string) *MessageSend {
	return &MessageSend{
		SenderId:    senderId,
		RecipientId: recipientId,
		Type:        messageType,
		Content:     content,
		SentAt:      time.Now(),
		IsRead:      0, // Assuming 0 means unread and 1 means read
	}
}
