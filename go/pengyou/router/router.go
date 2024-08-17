package router

import (
	"pengyou/docs"
	"pengyou/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func ServiceRouter() *gin.Engine {
	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// test
	test := r.Group("/test")
	{
		test.GET("/test", service.Test)
	}

	// websocker connect
	conn := r.Group("/conn")
	{
		conn.GET("/establish/websocket", service.EstablishWsConn)
	}

	//
	// chat := r.Group("/chat")
	// {
	// 	chat.POST("/start", service.HandleMessage)
	// }

	post := r.Group("/post")
	{
		post.POST("/upload", service.PostUpload)
	}

	return r
}
