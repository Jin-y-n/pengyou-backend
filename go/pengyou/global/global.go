package global

import (
	"fmt"

	"pengyou/global/config"
	"pengyou/router"
	"pengyou/service"
	"pengyou/utils/log"
	"pengyou/utils/storage"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	globalConfig *config.Config
	Viper        *viper.Viper
	GinEngine    *gin.Engine
)

func Init() {
	// init globalConfig
	initViper()

	if err := Viper.Unmarshal(&globalConfig); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %w", err))
	}

	// init Logger
	log.NewZapLogger(&globalConfig.Zap)

	service.Init(globalConfig)
	storage.Init(globalConfig)
	GinEngine = router.Init(globalConfig)

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
