package controller

import (
	"pengyou/constant"
	"pengyou/model/common/response"
	"pengyou/service"
	"pengyou/utils/check/string"
	"pengyou/utils/log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HeartBeat this function is listening to the heartbeat of the user
func HeartBeat(c *gin.Context) {
	header := c.GetHeader(constant.UserId)

	if string.IsNumberString(&header) {
		id, err := strconv.ParseInt(header, 10, 64)
		if err != nil {
			response.FailWithMessage(constant.RequestArgumentError, c)
			return
		}
		service.HeartBeat(c, uint(id))
	} else {
		log.Error("user id is not a number")
		response.FailWithMessage(constant.RequestArgumentError, c)
	}
}

// Establish this is for connect user to server
func Establish(c *gin.Context) {

	userIdStr := c.GetHeader(constant.UserId)

	if !string.IsNumberString(&userIdStr) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	service.EstablishWsConn(c, uint(userId))
}

// Shutdown this is for shutdown the connection from the user
func Shutdown(c *gin.Context) {

	userIdStr := c.GetHeader(constant.UserId)

	if !string.IsNumberString(&userIdStr) {

		response.FailWithMessage(constant.RequestArgumentError, c)

		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	service.ShutdownWsConn(c, uint(userId))
}

// EstablishChatTo establish chat to the target user
func EstablishChatTo(c *gin.Context) {
	userIdStr := c.GetHeader(constant.UserId)
	objectIdStr, success := c.GetPostForm(constant.ChatterId)

	if !string.IsNumberString(&userIdStr) || !success || !string.IsNumberString(&objectIdStr) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	objectId, _ := strconv.ParseUint(objectIdStr, 10, 64)

	service.EstablishChatTo(c, uint(userId), uint(objectId))
}

// CutChat cut chat with the target user
func CutChat(c *gin.Context) {

	userIdStr := c.GetHeader(constant.UserId)
	objectIdStr, success := c.GetPostForm(constant.ChatterId)

	if !string.IsNumberString(&userIdStr) {
		log.Error("userId is wrong")
		response.FailWithMessage(constant.RequestArgumentError, c)
	}
	if !success || !string.IsNumberString(&objectIdStr) {
		log.Error("chatterId is wrong")
		response.FailWithMessage(constant.RequestArgumentError, c)
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	objectId, _ := strconv.ParseUint(objectIdStr, 10, 64)

	service.CutChat(c, uint(userId), uint(objectId))
}
