package entity

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
	"time"
)

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

func (table *Post) TableName() string {
	return "post"
}

// ElasticsearchMapping creates an Elasticsearch mapping for the Post struct.
func (*Post) ElasticsearchMapping(esClient *elasticsearch.Client) (map[string]interface{}, error) {
	mapping := map[string]interface{}{
		"properties": map[string]interface{}{
			"author": map[string]interface{}{
				"type": "long",
			},
			"title": map[string]interface{}{
				"type": "text",
			},
			"content": map[string]interface{}{
				"type": "text",
			},
			"status": map[string]interface{}{
				"type": "byte",
			},
			"created_person": map[string]interface{}{
				"type": "long",
			},
			"modified_person": map[string]interface{}{
				"type": "long",
			},
			"delete_at": map[string]interface{}{
				"type": "date",
			},
			"modified_by_admin": map[string]interface{}{
				"type": "byte",
			},
			"id": map[string]interface{}{
				"type": "long",
			},
			"created_at": map[string]interface{}{
				"type": "date",
			},
			"updated_at": map[string]interface{}{
				"type": "date",
			},
			"deleted_at": map[string]interface{}{
				"type": "date",
			},
		},
	}

	return mapping, nil
}
