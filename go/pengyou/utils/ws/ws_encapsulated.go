package ws

import (
	"pengyou/utils/log"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func Close(wsCon *websocket.Conn) {
	if wsCon == nil {
		log.Warn("websocket connection is nil")
		return
	}
	err := wsCon.Close()

	if err != nil {
		log.Error("close websocket error", zap.Error(err))
	}
}

func SendTextMessage(wsCon *websocket.Conn, message string) {
	err := wsCon.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Error("send message error", zap.Error(err))
	}
}
