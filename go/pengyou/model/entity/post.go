package entity

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Title   string
	Content string

	Status int

	Author       string
	UpdatePerson uint64
}

func (table *Post) TableName() string {
	return "post"
}
