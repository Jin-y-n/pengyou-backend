package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"pengyou/controller"
	"pengyou/docs"
)

func GinRouter(r *gin.Engine) {
	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// websocket connect
	conn := r.Group("/connect")
	{
		conn.GET("/establish", controller.Establish)
		conn.POST("/establish-chat-to", controller.EstablishChatTo)
		conn.POST("/shutdown", controller.Shutdown)
		conn.POST("/cut-chat-from", controller.CutChat)
		conn.POST("/heart-beat", controller.HeartBeat)
	}

	chat := r.Group("/chat")
	{
		chat.POST("/confirm-receive", controller.ReceiveMsgConfirm)
		chat.POST("/leave-msg", controller.LeaveMsg)
		chat.POST("/get-unread-msg", controller.GetUnreadMsg)
	}

	search := r.Group("/query")
	{
		search.POST("/post", controller.SearchPost)
	}

	post := r.Group("/post")
	{
		post.POST("/add", controller.AddPost)
		post.POST("/update", controller.UpdatePost)
		post.POST("/delete", controller.DeletePost)
	}
}
