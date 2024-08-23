package model

import "time"

type MessageConfirmRec struct {
	UserId      uint      `json:"user_id"`
	MessageId   []uint    `json:"message_id"`
	ConfirmTime time.Time `json:"confirm_time"`
}
