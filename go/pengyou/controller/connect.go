package controller

import (
	"pengyou/constant"
	"pengyou/model/common/response"
	"pengyou/service"
	"pengyou/utils/check"
	"pengyou/utils/log"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Establish(c *gin.Context) {

	userIdStr := c.GetHeader(constant.USER_ID)

	if check.IsBlank(&userIdStr) {
		log.Error("userId is blank")
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		log.Error("userId is not a number", zap.Error(err))
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	service.EstablishWsConn(c, uint(userId))
}

func Shutdown(c *gin.Context) {

	userIdStr := c.GetHeader(constant.USER_ID)

	if check.IsBlank(&userIdStr) {
		log.Error("userId is blank")
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		log.Error("userId is not a number", zap.Error(err))
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	service.ShutdownWsConn(c, uint(userId))

}

func CutChat(c *gin.Context) {

	userIdStr := c.GetHeader(constant.USER_ID)
	if check.IsBlank(&userIdStr) {
		log.Error("userId is blank")
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		log.Error("userId is not a number", zap.Error(err))
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	objectIdStr, succ := c.GetPostForm(constant.CHATTER_ID)
	if !succ {
		log.Error("chatterId is blank")
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	objectId, err := strconv.ParseUint(objectIdStr, 10, 64)
	if err != nil {
		log.Error("chatterId is not a number", zap.Error(err))
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	service.CutChat(c, uint(userId), uint(objectId))
}
