package router

import (
	"pengyou/global/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(cfg *config.Config) *gin.Engine {
	eng := gin.Default()

	corsCfg := cors.New(
		cors.Config{
			AllowAllOrigins: true, // 允许所有来源
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowHeaders:    []string{"user_id", "Origin", "Content-Length", "Content-Type"}, // 只需添加 user_id
		})

	eng.Use(corsCfg)
	//eng.Use(jwtMiddleware())
	GinRouter(eng)

	return eng
}
