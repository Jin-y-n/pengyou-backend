package storage

import (
	"go.uber.org/zap"
	"pengyou/global/config"
	"pengyou/utils/log"

	"github.com/elastic/go-elasticsearch/v8"
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
		log.Error("elasticsearch init failed")
		panic(err)
	}
	ping, err := EsClient.Ping()
	if err != nil {
		if ping != nil {
			log.Info(ping.String())
		}

		log.Error("elasticsearch connect failed!", zap.Error(err))
		return
	}

	log.Info("elasticsearch connect success")

}

func AddDoc(index, id, doc string) error {
	_, err := EsClient.Index(index, nil, nil)

	log.Info("add doc")

	if err != nil {
		log.Error("add doc failed")
		return err
	}

	return err

}
