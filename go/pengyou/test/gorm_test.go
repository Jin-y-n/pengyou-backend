package test

import (
	"fmt"
	"pengyou/model/entity"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TestGorm demonstrates basic GORM operations.
func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open(getDBConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		panic("failed to auto migrate")
	}

	// Create a new user
	user := &entity.User{
		Username:      "Napbad",
		Password:      "123456",
		Phone:         "123456789",
		Email:         "123456789@qq.com",
		ClientIP:      "127.0.0.1",
		DeviceInfo:    "123456789",
		LoginTime:     time.Now(),
		HeartBeatTime: time.Now(),
		IsLogout:      0,
		LogOutTime:    time.Now(),

		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ID:        1,
			DeletedAt: gorm.DeletedAt{},
		},
	}

	if err := db.Create(user).Error; err != nil {
		panic("failed to create user")
	}

	// Read a user record
	var userForUse entity.User
	if err := db.Where("id = ?", 1).First(&userForUse).Error; err != nil {
		panic("failed to read user")
	}
	fmt.Println(userForUse)

	// Delete a user record
	if err := db.Delete(&entity.User{}, "ID = ?", 1).Error; err != nil {
		panic("failed to delete user")
	}
}

// getDBConnectionString returns the database connection string.
func getDBConnectionString() string {
	return "root:191019sJs_MySQL@tcp(8.137.96.68:3306)/pengyou?charset=utf8&parseTime=true"
}
