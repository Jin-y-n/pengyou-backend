package service

import (
	"github.com/gorilla/websocket"
	"net/http"
	"pengyou/global/config"
	"pengyou/utils/log"
)

// Init service package
func Init(cfg *config.Config) {
	upGrade = websocket.Upgrader{
		ReadBufferSize:  cfg.Files.ReadBufSize,
		WriteBufferSize: cfg.Files.WriteBufSize,
		CheckOrigin: func(r *http.Request) bool {
			log.Logger.Info("check origin")
			return true
		},
	}
	log.Logger.Info("init service: websocket.upGrader")

	CheckUserConnect()
	log.Logger.Info("begin removing disconnected user...")
}
