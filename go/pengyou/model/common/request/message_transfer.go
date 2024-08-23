package request

import (
	"fmt"
	"pengyou/constant"
	"time"
)

type MessageTransfer struct {
	SenderId    uint      `json:"sender_id"`
	RecipientId uint      `json:"recipient_id"`
	CreateAt    time.Time `json:"create_at"`

	Type    int    `json:"type"`
	Content string `json:"content"`
}

// String returns the string representation of the message type.
func String(m int) string {
	switch m {
	case constant.UnknownMessageType:
		return "UnknownMessageType"
	case constant.TextMessage:
		return "TextMessage"
	case constant.FileRequestMessage:
		return "FileRequestMessage"
	case constant.FriendRequestMessage:
		return "FriendRequestMessage"
	case constant.EstablishChatMessage:
		return "EstablishChatMessage"
	case constant.CutChatMessage:
		return "CutChatMessage"
	default:
		return fmt.Sprintf("MessageType(%d)", m)
	}
}
