package main

import (
	"pengyou/global"
	"pengyou/global/config"
	"pengyou/router"
)

func main() {

	global.Init()

	r := router.ServiceRouter()

	r.Run(config.Cfg.Server.Port)
}
