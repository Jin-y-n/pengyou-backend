package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"go.uber.org/zap"
	"io"
	"pengyou/constant"
	"pengyou/model/entity"
	"pengyou/utils/common"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestElasticSearch demonstrates basic Elasticsearch operations.
func TestElasticSearch(t *testing.T) {
	// Set up the Elasticsearch client
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		t.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	ping, err := esClient.Ping()
	if err != nil {
		if ping != nil {
			t.Error(ping.String())
		}

		t.Error("elasticsearch connect failed!", zap.Error(err))
		panic(err)
		return
	}

	// Prepare the current time
	now := time.Now()
	future := &time.Time{}

	generator, err := common.NewSnowflakeIDGenerator(1)
	if err != nil {
		fmt.Println("init snowflake failed")
		return
	}

	// Create a new post
	post := &entity.Post{
		Author:          1,
		Title:           "test",
		Content:         "this is a test post",
		Status:          1,
		CreatedPerson:   UintPtr(1),
		ModifiedPerson:  UintPtr(1),
		DeleteAt:        future,
		ModifiedByAdmin: 0,
		ID:              uint(generator.NextID()),
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	// Index the document
	//indexName := "post"
	fmt.Println("-------------save----------------")
	fmt.Println(post)
	marshal, err := json.Marshal(post)
	if err != nil {
		return
	}
	fmt.Println("-------------save----------------")

	request := esapi.IndexRequest{
		Index:      constant.PostIndex,
		DocumentID: strconv.Itoa(int(post.ID)), // Assuming ID is an auto-incrementing integer
		Body:       strings.NewReader(string(marshal)),
		Refresh:    "true",
	}
	res, err := request.Do(context.Background(), esClient)
	if err != nil {
		return
	}

	// Check the response
	if res.IsError() {
		t.Errorf("Received an error when indexing: %s", res.String())
	}

	// Retrieve the document

	query :=
		`
{
  "query": {
    "match_phrase": {
      "title": "test"
    }
  }
}
`
	getRes, err := esClient.Search(
		esClient.Search.WithIndex(constant.PostIndex),
		esClient.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		t.Fatalf("Error getting document: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(getRes.Body)

	// Check the response
	if getRes.IsError() {
		t.Errorf("Received an error when getting: %s", getRes.String())
	}
	//
	//
	//var resBytes []byte
	//read, err := getRes.Body.Read(resBytes)
	//if err != nil {
	//	t.Errorf("read failed!")
	//	return
	//}

	fmt.Println("------------res-------------")
	fmt.Println(getRes.String())
	fmt.Println("------------res-------------")
	//// Parse the response into a Post object
	//var retrievedPost entity.Post
	//err = esapi.UnmarshalGetResult(getRes, &retrievedPost)
	//if err != nil {
	//	t.Errorf("Error unmarshalling the result: %s", err)
	//}

	//getRes.Body.Read()

	// Assert that the retrieved post matches the original post
	//assert.Equal(t, post.Title, retrievedPost.Title)
	//assert.Equal(t, post.Content, retrievedPost.Content)
}

// UintPtr returns a pointer to an unsigned integer.
func UintPtr(i uint) *uint {
	return &i
}

func IntPtr(i int) *int {
	var iPtr = i
	return &iPtr
}

func UintPtrFromInt(i int) *uint {
	var iPtr = uint(i)
	return &iPtr
}
