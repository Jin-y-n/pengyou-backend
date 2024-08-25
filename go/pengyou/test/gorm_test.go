package test

import (
	"fmt"
	"pengyou/model/entity"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TestGorm demonstrates basic GORM operations.
func TestGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open(getDBConnectionString()), &gorm.Config{
		// Add any specific GORM configuration here
	})
	if err != nil {
		t.Errorf("failed to connect database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		t.Errorf("failed to auto migrate: %v", err)
	}

	// Create a new user
	user := &entity.User{
		Username:      "Napbad",
		Password:      "123456",
		Phone:         "123456789",
		Email:         "123456789@qq.com",
		ClientIP:      "127.0.0.1",
		DeviceInfo:    "123456789",
		HeartBeatTime: time.Now(),

		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			ID:        1,
			DeletedAt: gorm.DeletedAt{}, // Ensure the DeletedAt field is initialized
		},
	}

	if err := db.Create(user).Error; err != nil {
		t.Errorf("failed to create user: %v", err)
	}

	// Read a user record
	var userForUse entity.User
	if err := db.Where("id = ?", 1).First(&userForUse).Error; err != nil {
		t.Errorf("failed to read user: %v", err)
	}
	fmt.Println(userForUse)

	// Delete a user record
	// Note: This deletes the user with ID 1, ensure that the user exists before running this test
	if err := db.Delete(&entity.User{}, "ID = ?", 1).Error; err != nil {
		t.Errorf("failed to delete user: %v", err)
	}
}

// getDBConnectionString returns the database connection string.
func getDBConnectionString() string {
	return "root:191019sJs_MySQL@tcp(8.137.96.68:3306)/pengyou?charset=utf8&parseTime=true"
}
