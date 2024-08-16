package entity

import (
	"gorm.io/gorm"
)

type PostHistory struct {
	ID uint64

	SourcePostID uint64
	Post         Post

	Title   string
	Content string
	Author  string

	Status int

	UpdatePerson uint64
	DeletedAt    gorm.DeletedAt
}

func (table *PostHistory) TableName() string {
	return "post_history"
}
