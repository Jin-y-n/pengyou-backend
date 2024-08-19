package controller

import (
	"pengyou/constant"
	"pengyou/model/common/response"
	"pengyou/service"
	"pengyou/utils/check"
	"pengyou/utils/log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// this function is listening to the heartbeat of the user
func HeartBeat(c *gin.Context) {
	header := c.GetHeader(constant.USER_ID)

	if check.IsNumberString(&header) {
		id, err := strconv.ParseInt(header, 10, 64)
		if err != nil {
			response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
			return
		}
		service.HeartBeat(c, uint(id))
	}
}

// this is for connect user to server
func Establish(c *gin.Context) {

	userIdStr := c.GetHeader(constant.USER_ID)

	if !check.IsNumberString(&userIdStr) {
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	service.EstablishWsConn(c, uint(userId))
}

// this is for shutdown the connection from the user
func Shutdown(c *gin.Context) {

	userIdStr := c.GetHeader(constant.USER_ID)

	if !check.IsNumberString(&userIdStr) {

		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)

		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	service.ShutdownWsConn(c, uint(userId))
}

// establish chat to the target user
func EstablishChatTo(c *gin.Context) {
	userIdStr := c.GetHeader(constant.USER_ID)
	objectIdStr, succ := c.GetPostForm(constant.CHATTER_ID)

	if !check.IsNumberString(&userIdStr) || !succ || !check.IsNumberString(&objectIdStr) {
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	objectId, _ := strconv.ParseUint(objectIdStr, 10, 64)

	service.EstablishChatTo(c, uint(userId), uint(objectId))
}

// cut chat with the target user
func CutChat(c *gin.Context) {

	userIdStr := c.GetHeader(constant.USER_ID)
	objectIdStr, succ := c.GetPostForm(constant.CHATTER_ID)

	if !check.IsNumberString(&userIdStr) {
		log.Error("userId is wrong")
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
	}
	if !succ || !check.IsNumberString(&objectIdStr) {
		log.Error("chatterId is wrong")
		response.FailWithMessage(constant.REQUEST_ARGUMENT_ERROR, c)
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	objectId, _ := strconv.ParseUint(objectIdStr, 10, 64)

	service.CutChat(c, uint(userId), uint(objectId))
}
