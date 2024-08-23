package entity

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user entity.
type User struct {
	gorm.Model

	Username      string    `gorm:"size:50" json:"username"`
	Password      string    `gorm:"size:64" json:"password"`
	Email         string    `gorm:"size:50" json:"email"`
	Phone         string    `gorm:"size:20" json:"phone"`
	LoginTime     time.Time `json:"login_time"`
	Status        int8      `json:"status"`
	HeartBeatTime time.Time `json:"heart_beat_time"`
	ClientIP      string    `gorm:"size:50" json:"client_ip"`
	IsLogout      int8      `json:"is_logout"`
	LogOutTime    time.Time `json:"log_out_time"`
	DeviceInfo    string    `gorm:"size:255" json:"device_info"`
	CreatedPerson uint      `json:"created_person"`
	UpdatedPerson uint      `json:"updated_person"`
}

func (table *User) TableName() string {
	return "user"
}
