package request

import (
	"fmt"
	"time"
)

type MessageIn struct {
	SenderId    uint      `json:"sender_id"`
	RecipientId uint      `json:"recipient_id"`
	CreateAt    time.Time `json:"create_at"`

	RequestType int    `json:"type"`
	Content     string `json:"content"`
}

const (
	UnknownMessageType = 0
	TextMessage        = 1
	FileRequestMessage = 2
)

func String(m int) string {
	switch m {
	case 1:
		return "TextMessage"
	case 2:
		return "FileRequestMessage"
	default:
		return fmt.Sprintf("MessageType(%d)", m)
	}
}
