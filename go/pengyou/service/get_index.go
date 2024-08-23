package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(
		http.StatusOK, "name: Napbad")
}

//	@BasePath	/api/v1

// Test PingExample godoc
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	json{"name": "Napbad", "func": "test"}
//	@Router			/test [get]
func Test(c *gin.Context) {
	c.JSON(
		http.StatusOK, "name: Napbad, func: test")
}
