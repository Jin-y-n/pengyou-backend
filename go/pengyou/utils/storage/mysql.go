package storage

import (
	"pengyou/global/config"
	plog "pengyou/utils/log"

	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func InitMySQL(cfg *config.Config) *gorm.DB {

	logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

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
	sqlDB.Ping()

	plog.Logger.Info("mysql connect success -> " +
		mysqlConfig.Host + ":" +
		strconv.Itoa(mysqlConfig.Port))

	return GormDB
}
