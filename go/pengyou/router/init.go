package router

import (
	"pengyou/constant"
	"pengyou/global/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(cfg *config.Config) *gin.Engine {
	eng := gin.Default()

	corsCfg := cors.New(
		cors.Config{
			AllowAllOrigins: true, // allow all origin
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowHeaders:    []string{"user_id", "Origin", "Content-Length", "Content-Type", constant.Authorization, "Upgrade", "upgrade "}, // 只需添加 user_id
		})

	eng.Use(corsCfg)
	//eng.Use(jwtMiddleware())
	GinRouter(eng)

	return eng
}
