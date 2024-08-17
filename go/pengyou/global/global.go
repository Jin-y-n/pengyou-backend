package global

import (
	"fmt"

	"pengyou/global/config"
	"pengyou/router"
	"pengyou/service"
	"pengyou/utils/log"
	"pengyou/utils/storage"

	"github.com/gin-contrib/cors"
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

	// init redis
	storage.InitRedis(&globalConfig.Redis)

	// init mysql
	storage.InitMySQL(globalConfig)

	storage.InitFile(globalConfig)

	service.Init(globalConfig)

	config.Cfg = globalConfig

	GinEngine = router.ServiceRouter()

	corsCfg := cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		})

	GinEngine.Use(corsCfg)
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
