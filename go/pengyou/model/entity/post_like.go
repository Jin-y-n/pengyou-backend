package entity

import (
	"gorm.io/gorm"
)

type PostLike struct {
	ID uint64

	Status int

	UpdatePerson uint64
	DeletedAt    gorm.DeletedAt
}

func (table *PostLike) TableName() string {
	return "post_like"
}
