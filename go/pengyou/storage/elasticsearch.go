package storage

import (
	"pengyou/global/config"

	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func InitElasticSearch(cfg *config.Config) {
	var err error
	EsClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: cfg.Elasticsearch.Nodes,
		Username:  cfg.Elasticsearch.Username,
		Password:  cfg.Elasticsearch.Password,
	})
	if err != nil {
		panic(err)
	}
}
