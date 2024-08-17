package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model

	UserID        uint      `gorm:"primaryKey" json:"user_id"`
	DisplayName   string    `gorm:"size:50" json:"display_name"`
	AvatarID      string    `gorm:"size:255" json:"avatar_id"`
	Bio           string    `gorm:"size:255" json:"bio"`
	Gender        int8      `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	Location      string    `gorm:"size:100" json:"location"`
	Occupation    string    `gorm:"size:100" json:"occupation"`
	Education     string    `gorm:"size:100" json:"education"`
	School        string    `gorm:"size:100" json:"school"`
	Major         string    `gorm:"size:100" json:"major"`
	Company       string    `gorm:"size:100" json:"company"`
	Position      string    `gorm:"size:100" json:"position"`
	Website       string    `gorm:"size:255" json:"website"`
	CreatedPerson uint      `json:"created_person"`
	UpdatedPerson uint      `json:"updated_person"`
}

func (table *UserProfile) TableName() string {
	return "user_profile"
}
