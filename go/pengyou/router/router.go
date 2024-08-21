package router

import (
	"pengyou/controller"
	"pengyou/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func GinRouter() *gin.Engine {
	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// websocket connect
	conn := r.Group("/connect")
	{
		conn.POST("/establish", controller.Establish)
		conn.POST("/establish-chat-to", controller.EstablishChatTo)
		conn.POST("/shutdown", controller.Shutdown)
		conn.POST("/cut-chat", controller.CutChat)
		conn.POST("/heart-beat", controller.HeartBeat)
	}

	chat := r.Group("/chat")
	{
		chat.POST("/saf", nil)
	}

	search := r.Group("/query")
	{
		search.POST("/post", controller.SearchPost)
	}

	return r
}
