package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"pengyou/constant"
	"pengyou/model/entity"
	"pengyou/utils/common"
	"pengyou/utils/log"
	"strconv"
	"strings"
	"time"
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

	// Define the update request
	updateReq := esapi.UpdateRequest{
		Index:      constant.PostIndex,         // Use the same index name as when you created it
		DocumentID: strconv.Itoa(int(post.ID)), // Assuming ID is an auto-incrementing integer
		Body:       strings.NewReader(string(doc)),
		Refresh:    "true",
	}

	// Execute the request
	res, err := updateReq.Do(context.Background(), EsClient)
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
		return fmt.Errorf("error updating document: %s", res.Status())
	}

	fmt.Println("Document updated successfully")
	return nil
}

type Post struct {
	gorm.Model
	Author          uint       `gorm:"not null" json:"author"`
	Title           string     `gorm:"type:varchar(255)" json:"title"`
	Content         string     `gorm:"type:text" json:"content"`
	Status          uint8      `gorm:"default:1" json:"status"`
	CreatedPerson   *uint      `gorm:"index;references:User(id);on_delete:set_null" json:"created_person"`
	ModifiedPerson  *uint      `gorm:"index;references:User(id);on_delete:set_null" json:"modified_person"`
	DeleteAt        *time.Time `gorm:"index" json:"delete_at"`
	ModifiedByAdmin uint8      `gorm:"default:0" json:"modified_by_admin"`
}

// IndexPostQuery searches for posts in Elasticsearch.
func IndexPostQuery(post *entity.Post) ([]*entity.Post, error) {
	searchSource := strings.NewReader(`
    {
        "query": {
            "match_all": {}
        }
    }`)

	request := esapi.SearchRequest{
		Index: []string{constant.PostIndex},
		Body:  searchSource,
	}

	response, err := request.Do(context.Background(), EsClient)
	if err != nil {
		log.Logger.Error("get post failed: es, err :", zap.Error(err))
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Logger.Warn("read-closer closed failed")
		}
	}(response.Body)

	// Check the response status code
	if response.IsError() {
		log.Logger.Error("error in the response:", zap.String("status", strconv.Itoa(response.StatusCode)))
		return nil, fmt.Errorf("error in the response: %d", response.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Logger.Error("error reading the response body:", zap.Error(err))
		return nil, err
	}

	// Parse the JSON response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Logger.Error("error parsing the JSON response:", zap.Error(err))
		return nil, err
	}

	// Extract hits
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	// Convert hits to Post structs
	var posts []*entity.Post
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		post := &entity.Post{
			Author:          uint(source["author"].(float64)),
			Title:           source["title"].(string),
			Content:         source["content"].(string),
			Status:          uint8(source["status"].(float64)),
			CreatedPerson:   uintPtr(source["created_person"].(float64)),
			ModifiedPerson:  uintPtr(source["modified_person"].(float64)),
			DeleteAt:        timePtr(source["delete_at"].(string)),
			ModifiedByAdmin: uint8(source["modified_by_admin"].(float64)),
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// uintPtr converts a float64 to a pointer to uint.
func uintPtr(f float64) *uint {
	i := uint(f)
	return &i
}

// timePtr converts a string to a pointer to time.Time.
func timePtr(s string) *time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil
	}
	return &t
}

func IndexPostDelete(post int) error {
	resQuery, err1 := IndexPostQuery(&entity.Post{
		Model: gorm.Model{ID: uint(post)},
	})

	if err1 != nil {
		return err1
	}

	postFull := resQuery[0]

	if !common.CheckUserIdDefault(postFull.Author) {
		contextDefault, b := common.GetTokenFromContextDefault()
		if !b {
			return errors.New("token is not exist")
		}
		return errors.New(contextDefault + ` are not the author`)
	}

	// Define the delete request
	deleteReq := esapi.DeleteRequest{
		Index:      constant.PostIndex, // Use the same index name as when you created it
		DocumentID: strconv.Itoa(post),
		Refresh:    "true",
	}

	// Execute the request
	res, err := deleteReq.Do(context.Background(), EsClient)
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
		return fmt.Errorf("error deleting document: %s", res.Status())
	}

	fmt.Println("Document deleted successfully")
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
