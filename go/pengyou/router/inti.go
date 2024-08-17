package router

import (
	"pengyou/global/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(cfg *config.Config) *gin.Engine {
	eng := GinRouter()

	corsCfg := cors.New(
		cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		})

	eng.Use(corsCfg)

	return eng
}
