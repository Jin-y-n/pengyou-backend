package entity

import (
	"time"

	"gorm.io/gorm"
)

type MessageSend struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

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
