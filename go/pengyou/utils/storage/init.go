package storage

import "pengyou/global/config"

func Init(cfg *config.Config) {
	// init redis
	InitRedis(&cfg.Redis)

	// init mysql
	InitMySQL(cfg)

	InitFile(cfg)

	InitElasticSearch(cfg)
}
