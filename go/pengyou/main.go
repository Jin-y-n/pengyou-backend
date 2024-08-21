package main

import (
	"go.uber.org/zap"
	"pengyou/global"
	"pengyou/global/config"
	"pengyou/utils/log"
)

func main() {

	global.Init()

	r := global.GinEngine

	err := r.Run(config.Cfg.Server.Port)
	if err != nil {

		log.Error("gin start failed: ", zap.Error(err))

		return
	}
}
