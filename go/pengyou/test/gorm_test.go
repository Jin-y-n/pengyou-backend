package test

import (
	"pengyou/model/entity"

	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func testGorm() {
	db, err := gorm.Open(mysql.Open("root:191019sJs_MySQL@tcp(8.137.96.68:3306)/pengyou?charset=utf8&parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&entity.User{})

	user := &entity.User{}

	user.Username = "Napbad"
	user.Password = "123456"
	user.Phone = "123456789"
	user.Email = "123456789@qq.com"
	user.ID = 1
	user.ClientIP = "127.0.0.1"
	user.DeviceInfo = "123456789"
	user.LoginTime = time.Now()
	user.HeartBeatTime = time.Now()
	user.IsLogout = 0
	user.LogOutTime = time.Now()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = gorm.DeletedAt{}

	// Create
	db.Create(user)

	// Read
	var userForUse entity.User
	db.Where("id = ?", 1).First(&userForUse) // 根据整型主键查找
	fmt.Println()
	fmt.Println(userForUse)
	// db.First(&userForUse, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 userForUse 的 price 更新为 200
	// db.Model(&userForUse).Update("Price", 200)
	// Update - 更新多个字段
	// db.Model(&userForUse).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&userForUse).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 userForUse
	db.Delete(&entity.User{}, "123456789")
}
