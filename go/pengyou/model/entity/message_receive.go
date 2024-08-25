package entity

import (
	"time"

	"gorm.io/gorm"
)

type MessageReceive struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	MessageSendId uint
	RecipientId   uint

	ReadAt   time.Time
	DeleteAt gorm.DeletedAt
}

func (table *MessageReceive) TableName() string {
	return "message_receive"
}
