package ws

import (
	"encoding/json"
	"pengyou/model/common/request"
	"pengyou/utils/log"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func Close(wsCon *websocket.Conn) {
	if wsCon == nil {
		log.Logger.Warn("websocket connection is nil")
		return
	}
	err := wsCon.Close()

	if err != nil {
		log.Logger.Error("close websocket error", zap.Error(err))
	}
}

func SendTextMessage(wsCon *websocket.Conn, sender, recipient uint, message string) bool {

	return SendMessage(
		wsCon,
		sender,
		recipient,
		message,
		websocket.TextMessage,
	)
}

func SendMessage(wsCon *websocket.Conn, sender, recipient uint, message string, messageType int) bool {
	log.Logger.Info("send message")

	msg := request.MessageTransfer{
		SenderId:    sender,
		RecipientId: recipient,
		Content:     message,
		CreateAt:    time.Now(),
		Type:        messageType,
	}

	res, err := json.Marshal(msg)
	if err != nil {
		log.Logger.Error("marshal message error", zap.Error(err))
		return false
	}

	err = wsCon.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		log.Logger.Error("send message error", zap.Error(err))
		return false
	}

	err = wsCon.WriteMessage(websocket.TextMessage, res)
	if err != nil {
		log.Logger.Error("send message error", zap.Error(err))
		return false
	}

	return true
}
