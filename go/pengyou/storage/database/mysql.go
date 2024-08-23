package db

import (
	"log"
	"pengyou/global/config"
	plog "pengyou/utils/log"

	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

var logger = plog.Logger

var GormDB *gorm.DB

func InitMySQL(cfg *config.Config) *gorm.DB {

	glogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		glogger.Config{
			SlowThreshold:             time.Millisecond,
			LogLevel:                  glogger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		})

	mysqlConfig := cfg.MySQL

	mysqlConfigString := mysqlConfig.User +
		":" +
		mysqlConfig.Password +
		"@tcp(" +
		mysqlConfig.Host +
		":" +
		strconv.Itoa(mysqlConfig.Port) +
		")/" +
		mysqlConfig.DB +
		"?" +
		mysqlConfig.Conf

	GormDB, err := gorm.Open(mysql.Open(mysqlConfigString), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := GormDB.DB()

	if err != nil {
		panic("failed to connect database")
	}
	err = sqlDB.Ping()
	if err != nil {
		panic("failed to connect database")
		return nil
	}

	plog.Logger.Info("mysql connect success -> " +
		mysqlConfig.Host + ":" +
		strconv.Itoa(mysqlConfig.Port))

	return GormDB
}
