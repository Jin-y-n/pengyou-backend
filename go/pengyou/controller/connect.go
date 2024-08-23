package controller

import (
	"pengyou/constant"
	"pengyou/model/common/response"
	"pengyou/service"
	"pengyou/utils/check/string"
	"pengyou/utils/common"
	"pengyou/utils/log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HeartBeat this function is listening to the heartbeat of the user
//
//	@Summary		Heartbeat for user connection
//	@Description	Checks the heartbeat of a user to maintain the connection.
//	@Tags			connections
//	@Accept			json
//	@Produce		json
//	@Param			userId	header		string	true	"User ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response	"Invalid input"
//	@Failure		500		{object}	response.Response	"Internal server error"
//	@Router			/connect/heart-beat [POST]
func HeartBeat(c *gin.Context) {
	header := c.GetHeader(constant.UserId)

	if string.IsNumberString(&header) {
		userId, err := strconv.ParseInt(header, 10, 64)
		if err != nil {
			response.FailWithMessage(constant.RequestArgumentError, c)
			return
		}

		if !common.CheckUserIdDefault(uint(userId)) {
			response.FailWithMessage(constant.RequestArgumentError, c)
			return
		}

		service.HeartBeat(c, uint(userId))
	} else {
		log.Logger.Error("user id is not a number")
		response.FailWithMessage(constant.RequestArgumentError, c)
	}
}

// Establish this is for connect user to server
//
//	@Summary		Establish a new connection
//	@Description	Establishes a new WebSocket connection for a user.
//	@Tags			connections
//	@Accept			json
//	@Produce		json
//	@Param			userId	header		string	true	"User ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response	"Invalid input"
//	@Failure		500		{object}	response.Response	"Internal server error"
//	@Router			/connect/establish [post]
func Establish(c *gin.Context) {

	userIdStr := c.GetHeader(constant.UserId)

	if !string.IsNumberString(&userIdStr) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	if !common.CheckUserIdDefault(uint(userId)) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	service.EstablishWsConn(c, uint(userId))
}

// Shutdown this is for shutdown the connection from the user
//
//	@Summary		Shutdown an existing connection
//	@Description	Shuts down an existing WebSocket connection for a user.
//	@Tags			connections
//	@Accept			json
//	@Produce		json
//	@Param			userId	header		string	true	"User ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response	"Invalid input"
//	@Failure		500		{object}	response.Response	"Internal server error"
//	@Router			/connect/shutdown [post]
func Shutdown(c *gin.Context) {

	userIdStr := c.GetHeader(constant.UserId)

	if !string.IsNumberString(&userIdStr) {

		response.FailWithMessage(constant.RequestArgumentError, c)

		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	if !common.CheckUserIdDefault(uint(userId)) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	service.ShutdownWsConn(c, uint(userId))
}

// EstablishChatTo establish chat to the target user
//
//	@Summary		Establish a chat session
//	@Description	Establishes a chat session with another user.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			userId		header		string	true	"User ID"
//	@Param			chatterId	formData	string	true	"Chatter ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response	"Invalid input"
//	@Failure		500			{object}	response.Response	"Internal server error"
//	@Router			/connect/establish-chat-to [post]
func EstablishChatTo(c *gin.Context) {
	userIdStr := c.GetHeader(constant.UserId)
	objectIdStr, success := c.GetPostForm(constant.ChatterId)

	if !string.IsNumberString(&userIdStr) || !success || !string.IsNumberString(&objectIdStr) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	objectId, _ := strconv.ParseUint(objectIdStr, 10, 64)

	if !common.CheckUserIdDefault(uint(userId)) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	service.EstablishChatTo(c, uint(userId), uint(objectId))
}

// CutChat cut chat with the target user
//
//	@Summary		Cut off a chat session
//	@Description	Cuts off a chat session with another user.
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			userId		header		string	true	"User ID"
//	@Param			chatterId	formData	string	true	"Chatter ID"
//	@Success		200			{object}	response.Response
//	@Failure		400			{object}	response.Response	"Invalid input"
//	@Failure		500			{object}	response.Response	"Internal server error"
//	@Router			/connect/cut-chat-from [post]
func CutChat(c *gin.Context) {

	userIdStr := c.GetHeader(constant.UserId)
	objectIdStr, success := c.GetPostForm(constant.ChatterId)

	if !string.IsNumberString(&userIdStr) {
		log.Logger.Error("userId is wrong")
		response.FailWithMessage(constant.RequestArgumentError, c)
	}
	if !success || !string.IsNumberString(&objectIdStr) {
		log.Logger.Error("chatterId is wrong")
		response.FailWithMessage(constant.RequestArgumentError, c)
	}

	userId, _ := strconv.ParseUint(userIdStr, 10, 64)
	objectId, _ := strconv.ParseUint(objectIdStr, 10, 64)

	if !common.CheckUserIdDefault(uint(userId)) {
		response.FailWithMessage(constant.RequestArgumentError, c)
		return
	}

	service.CutChat(c, uint(userId), uint(objectId))
}
