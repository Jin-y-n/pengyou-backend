package entity

import (
	"time"

	"gorm.io/gorm"
)

type MessageReceive struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`

	MessageSenderId uint
	RecipientId     uint

	ReadAt   time.Time
	DeleteAt gorm.DeletedAt

	Type int
}

func (table *MessageReceive) TableName() string {
	return "message_receive"
}
