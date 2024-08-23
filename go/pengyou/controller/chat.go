package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pengyou/constant"
	"pengyou/model"
	"pengyou/model/common/response"
	"pengyou/model/entity"
	"pengyou/service"
	db "pengyou/storage/database"
	string2 "pengyou/utils/check/string"
	"pengyou/utils/log"
	"strconv"
	"time"
)

// LeaveMsg
//	@Summary		Send a new message
//	@Description	Sends a new message between users.
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Param			message	body		entity.MessageSend	true	"The message details"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response	"Invalid input"
//	@Failure		500		{object}	response.Response	"Internal server error"
//	@Router			/chat/leave-msg [post]
func LeaveMsg(c *gin.Context) {

	msg := &entity.MessageSend{
		ID:          0,
		SenderId:    0,
		RecipientId: 0,
		Type:        0,
		Content:     "",
		SentAt:      time.Time{},
		DeleteAt:    gorm.DeletedAt{},
		IsRead:      0,
	}

	err := c.ShouldBindJSON(msg)
	if err != nil {
		log.Logger.Error("convert to object failed: ", zap.Error(err))
		response.FailWithMessage("input invalid", c)
		return
	}

	if msg.SentAt.Add(5 * time.Second).Before(time.Now()) {
		response.FailWithMessage(constant.ConnectTimeOut, c)
		return
	}

	if msg.Type == 0 || msg.SenderId == 0 || msg.RecipientId == 0 {
		response.FailWithMessage("input invalid", c)
		return
	}

	service.LeaveMsg(msg, c)
}

// ReceiveMsgConfirm confirm that message is received
//	@Summary		Confirm receipt of a message
//	@Description	Confirms that a message has been received.
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Param			confirmation	body		model.MessageConfirmRec	true	"Confirmation details"
//	@Success		200				{object}	response.Response
//	@Failure		400				{object}	response.Response	"Invalid input"
//	@Failure		500				{object}	response.Response	"Internal server error"
//	@Router			/chat/confirm-receive [post]
func ReceiveMsgConfirm(c *gin.Context) {
	mc := &model.MessageConfirmRec{}
	err := c.ShouldBindJSON(mc)

	if err != nil {
		response.FailWithMessage(constant.InvalidParams, c)
		return
	}

	if mc.ConfirmTime.Add(5 * time.Second).Before(time.Now()) {
		response.FailWithMessage(constant.ConnectTimeOut, c)
		return
	}

	for _, id := range mc.MessageId {
		err = db.StoreMsgConfirm(id)
		err := db.StoreMsgRec(id, mc.ConfirmTime)
		if err != nil {
			response.FailWithMessage(constant.ServerError, c)
			return
		}
	}
}

// GetUnreadMsg
//	@Summary		Get unread messages
//	@Description	Retrieves unread messages for a user.
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Param			userId	formData	string	true	"User ID"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response	"Invalid input"
//	@Failure		500		{object}	response.Response	"Internal server error"
//	@Router			/chat/get-unread-msg [post]
func GetUnreadMsg(c *gin.Context) {
	id, success := c.GetPostForm(constant.UserId)

	if !success {
		response.FailWithMessage(constant.InvalidParams, c)
	}

	if string2.IsNumberString(&id) {
		response.FailWithMessage(constant.InvalidParams, c)
		return
	}

	userId, _ := strconv.Atoi(id)

	service.GetMsgUnRead(uint(userId), c)
}
