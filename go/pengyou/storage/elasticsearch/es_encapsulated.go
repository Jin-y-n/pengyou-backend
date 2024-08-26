package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"go.uber.org/zap"
	"io"
	"pengyou/constant"
	esmodel "pengyou/model/elsaticsearch"
	"pengyou/model/entity"
	"pengyou/utils/log"
	"strconv"
	"strings"
)

func Store() {

}

func IndexPostAdd(post *entity.Post) error {
	// Convert the Post object to a JSON-compatible format
	doc, err := json.Marshal(post)
	if err != nil {
		return err
	}

	// Index the document
	indexReq := esapi.IndexRequest{
		Index:      constant.PostIndex,
		DocumentID: strconv.Itoa(int(post.ID)), // Assuming ID is an auto-incrementing integer
		Body:       strings.NewReader(string(doc)),
		Refresh:    "true",
	}

	// Execute the request
	res, err := indexReq.Do(context.Background(), EsClient)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Logger.Error("close reader failed: es, err :", zap.Error(err))
		}
	}(res.Body)

	// Check the response status code
	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}

	fmt.Println("Document indexed successfully")
	return nil
}

func IndexPostUpdate(post *entity.Post) error {
	// Convert the Post object to a JSON-compatible format
	doc, err := json.Marshal(post)
	if err != nil {
		return err
	}

	update, err := EsClient.Update(
		constant.PostIndex,
		strconv.Itoa(int(post.ID)),
		strings.NewReader(string(doc)))
	if err != nil {
		log.Logger.Error("update failed", zap.Error(err))
		return err
	}

	if update.IsError() {
		log.Logger.Error("update failed ",
			zap.String("code: ", strconv.Itoa(update.StatusCode)))
		return err
	}

	fmt.Println("Document updated successfully")
	return nil
}

func PostQueryById(id uint) (*entity.Post, error) {
	return QueryByID(id, constant.PostIndex)
}

// PostQueryByTitleContent searches for posts in Elasticsearch.
func PostQueryByTitleContent(post *entity.Post) ([]*entity.Post, error) {
	query := fmt.Sprintf(`
{
        "query": {
            "match_phrase": {
				"title": %s,
				"content": %s
			}
        }
}
`, post.Title, post.Content)

	response, err := Query(constant.PostIndex, query)
	if err != nil {
		return nil, err
	}
	esPosts := response.Hits.Hits

	var posts []*entity.Post

	for _, v := range esPosts {
		posts = append(posts, &v.Source)
	}

	return posts, nil
}

func PostQueryByTitle(post *entity.Post) ([]*entity.Post, error) {
	query := fmt.Sprintf(`
{
        "query": {
            "match_phrase": {
				"title": %s,
			}
        }
}
`, post.Title)

	response, err := Query(constant.PostIndex, query)
	if err != nil {
		return nil, err
	}
	esPosts := response.Hits.Hits

	var posts []*entity.Post

	for _, v := range esPosts {
		posts = append(posts, &v.Source)
	}

	return posts, nil
}

func PostQueryByContent(post *entity.Post) ([]*entity.Post, error) {
	query := fmt.Sprintf(`
{
        "query": {
            "match_phrase": {
				"content": %s
			}
        }
}
`, post.Content)

	response, err := Query(constant.PostIndex, query)
	if err != nil {
		return nil, err
	}
	esPosts := response.Hits.Hits

	var posts []*entity.Post

	for _, v := range esPosts {
		posts = append(posts, &v.Source)
	}

	return posts, nil
}

// QueryByID query by id
func QueryByID(id uint, index string) (*entity.Post, error) {
	query := fmt.Sprintf(`
	{
	  "query": {
		"match": {
		  "ID": %d
		}
	  }
	}
	`, id)

	response, err := Query(constant.PostIndex, query)
	if err != nil {
		return nil, err
	}
	hits := response.Hits.Hits

	if len(hits) != 1 {
		return nil, errors.New("not found")
	}

	return &hits[0].Source, nil
}

func IndexPostDelete(post, author int) error {
	queryByID, err := QueryByID(uint(post), constant.PostIndex)

	if err != nil {
		log.Logger.Error("query by id failed", zap.Error(err))
		return err
	}

	if !(queryByID.Author == uint(author)) {
		log.Logger.Error("no Authorization to delete")
		return errors.New("no Authorization to delete")
	}

	response, err := EsClient.Delete(constant.PostIndex, strconv.Itoa(post))
	if err != nil {
		log.Logger.Error("delete failed: "+strconv.Itoa(post),
			zap.Error(err))
		return err
	}
	if response.IsError() {
		return errors.New("delete failed")
	}

	return nil
}

func AddDoc(index, id, doc string) error {
	_, err := EsClient.Index(index, nil, nil)

	log.Logger.Info("add doc")

	if err != nil {
		log.Logger.Error("add doc failed")
		return err
	}

	return err

}

func ExistIndex(indexName string) bool {
	resp, err := EsClient.Indices.
		Exists([]string{indexName})
	if err != nil {
		fmt.Printf("check index failed, err:%v\n", err)
		return false
	}

	log.Logger.Info("check index: ",
		zap.String("index", indexName),
		zap.String("status", resp.Status()))

	if resp.Status() != "200" {
		return false
	}

	return true
}

// CreateIndex create index
func CreateIndex(indexName string) error {
	resp, err := EsClient.Indices.
		Create(indexName)

	if err != nil {
		fmt.Printf("create index failed, err:%v\n", err)
		return err
	}
	fmt.Printf("index:%#v\n", resp.Body)
	return nil
}

// CreateIndexWithBody create index
func CreateIndexWithBody(indexName string, body map[string]interface{}) error {
	resp, err := EsClient.Indices.
		Create(indexName)
	if resp == nil {
		log.Logger.Error("resp is nil")
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Logger.Error("close body failed")
		}
	}(resp.Body)

	if err != nil {
		fmt.Printf("create index failed, err:%v\n", err)
		return err
	}

	req := esapi.IndicesPutMappingRequest{
		Index: []string{indexName},
		Body:  esutil.NewJSONReader(body),
	}

	do, err := req.Do(context.Background(), EsClient)
	if err != nil {
		log.Logger.Error("create index error", zap.Error(err))
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Logger.Error("close body failed")
		}
	}(do.Body)

	if do.IsError() {
		log.Logger.Error("create index error", zap.Error(err))
		return err
	}

	fmt.Printf("index:%#v\n", resp.Body)
	return nil
}

// DeleteIndex delete index
func DeleteIndex(indexName string) error {
	_, err := EsClient.Indices.
		Delete([]string{indexName})
	if err != nil {
		fmt.Printf("delete index failed,err:%v\n", err)
		return err
	}
	fmt.Printf("delete index successed,indexName:%s", indexName)
	return nil
}

func Query(index string, body string) (*esmodel.Response, error) {
	getRes, err := EsClient.Search(
		EsClient.Search.WithIndex(constant.PostIndex),
		EsClient.Search.WithBody(strings.NewReader(body)),
	)

	if err != nil {
		log.Logger.Error("get post failed: es, err :", zap.Error(err))
		return nil, err
	}

	if getRes.IsError() {
		log.Logger.Error("error in the response:", zap.String("status", strconv.Itoa(getRes.StatusCode)))
		return nil, fmt.Errorf("error in the response: %d", getRes.StatusCode)
	}

	response := &esmodel.Response{}
	err = json.Unmarshal([]byte(getRes.String()), &response)
	if err != nil {
		log.Logger.Error("unmarshal failed: es, err :", zap.Error(err))
		return nil, err
	}

	return response, nil
}
