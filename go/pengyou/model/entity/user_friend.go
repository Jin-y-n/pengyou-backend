package entity

import (
	"time"
)

type UserFriend struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint      `json:"user_id"`
	FriendID      uint      `json:"friend_id"`
	Status        int8      `json:"status"`
	RequestDate   time.Time `json:"request_date"`
	AcceptedDate  time.Time `json:"accepted_date"`
	RequirePerson uint      `json:"require_person"`
	Relationship  int16     `json:"relationship"`
	DeleteAt      time.Time `json:"delete_at"`
}
