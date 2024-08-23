package controller

import (
	"github.com/gin-gonic/gin"
	"pengyou/constant"
	"pengyou/model/common/request"
	"pengyou/model/common/response"
	"pengyou/service"
)

func SearchPost(c *gin.Context) {

	pqi := &request.PostQueryInput{}

	err := c.ShouldBindJSON(pqi)
	if err != nil {
		response.FailWithMessage(constant.InvalidParams, c)
		return
	}

	service.SearchPost(pqi, c)
}
