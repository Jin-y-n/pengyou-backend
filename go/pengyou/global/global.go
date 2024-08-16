package global

import (
	"fmt"

	"pengyou/global/config"
	"pengyou/utils/log"
	"pengyou/utils/storage"

	"github.com/spf13/viper"
)

var (
	globalConfig *config.Config
	Viper        *viper.Viper
)

func Init() {
	// init globalConfig
	initViper()

	if err := Viper.Unmarshal(&globalConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}

	// init Logger
	log.NewZapLogger(&globalConfig.Zap)

	// init redis
	storage.InitRedis(&globalConfig.Redis)

	// init mysql
	storage.InitMySQL(globalConfig)

	storage.InitFile(globalConfig)

	config.Cfg = globalConfig
}

func initViper() {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w ", err))
	}

	Viper = v
}
