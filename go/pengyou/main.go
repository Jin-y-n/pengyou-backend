package main

import (
	"pengyou/global"
	"pengyou/global/config"
)

func main() {

	global.Init()

	r := global.GinEngine

	r.Run(config.Cfg.Server.Port)
}
