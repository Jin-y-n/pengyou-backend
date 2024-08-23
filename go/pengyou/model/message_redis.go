package model

import "time"

// MessageRedis -> Type of message that stored in redis

type MessageRedis struct {
	ID          uint      `json:"id"`
	SenderId    uint      `json:"sender_id"`
	RecipientId uint      `json:"recipient_id"`
	Type        int       `json:"type"`
	SentAt      time.Time `json:"sent_at"`
	ReceiveAt   time.Time `json:"receive_at"`

	Content string `json:"content"`
}
