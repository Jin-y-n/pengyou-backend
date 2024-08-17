package service

import (
	"net/http"
	"pengyou/global/config"
	"pengyou/utils/log"

	"github.com/gorilla/websocket"
)

// init service package
func Init(cfg *config.Config) {
	upGrade = websocket.Upgrader{
		ReadBufferSize:  cfg.Files.ReadBufSize,
		WriteBufferSize: cfg.Files.WriteBufSize,
		CheckOrigin: func(r *http.Request) bool {
			log.Info("check origin")
			return true
		},
	}
	log.Info("init service: websocket.upGrader")
}
