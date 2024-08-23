package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"pengyou/global/config"
	"pengyou/model/entity"
	"pengyou/utils/log"
)

var EsClient *elasticsearch.Client

func InitElasticSearch(cfg *config.Config) {
	var err error
	EsClient, err = elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: cfg.Elasticsearch.Nodes,
			//Username:  cfg.Elasticsearch.Username,
			//Password:  cfg.Elasticsearch.Password,
		})
	if err != nil {
		log.Logger.Error("elasticsearch init failed")
		panic(err)
	}
	ping, err := EsClient.Ping()
	if err != nil {
		if ping != nil {
			log.Logger.Info(ping.String())
		}

		log.Logger.Error("elasticsearch connect failed!", zap.Error(err))
		panic(err)
		return
	}

	log.Logger.Info("elasticsearch connect success")

	if !ExistIndex("post") {
		log.Logger.Info("elasticsearch: index not found", zap.String("index", "post"))
		post := entity.Post{}

		res, err := post.ElasticsearchMapping(EsClient)
		if err != nil {
			log.Logger.Error("create index: post failed")
			return
		}
		err = CreateIndexWithBody("post", res)
		if err != nil {
			log.Logger.Error("create index failed")
			panic(err)
			return
		}
		log.Logger.Info("elasticsearch: index create success", zap.String("index", "post"))
	}
}
