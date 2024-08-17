package storage

import (
	"pengyou/global/config"

	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func InitElasticSearch(cfg *config.Config) {
	var err error
	EsClient, err = elasticsearch.NewClient(elasticsearch.Config{

		Username: cfg.Elasticsearch[0].Username,
		Password: cfg.Elasticsearch[0].Password,
	})
	if err != nil {
		panic(err)
	}
}
