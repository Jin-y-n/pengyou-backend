package request

import "time"

type PostQueryInput struct {
	ID              uint       `json:"id" form:"id"`                           // ID of the post
	PageInfo        PageInfo   `json:"pageInfo" form:"pageInfo"`               // Pagination information
	Title           string     `json:"title" form:"title"`                     // Title of the post
	Content         string     `json:"content" form:"content"`                 // Content of the post
	Author          string     `json:"author" form:"author"`                   // Author of the post
	Status          uint8      `json:"status" form:"status"`                   // Status of the post
	CreatedPerson   *uint      `json:"createdPerson" form:"createdPerson"`     // Created person ID
	ModifiedPerson  *uint      `json:"modifiedPerson" form:"modifiedPerson"`   // Modified person ID
	DeleteAt        *time.Time `json:"deleteAt" form:"deleteAt"`               // Deletion timestamp
	ModifiedByAdmin uint8      `json:"modifiedByAdmin" form:"modifiedByAdmin"` // Whether modified by admin
}

type PostCreateInput struct {
	Title           string     `json:"title" form:"title"`                     // Title of the post
	Content         string     `json:"content" form:"content"`                 // Content of the post
	Author          string     `json:"author" form:"author"`                   // Author of the post
	Status          uint8      `json:"status" form:"status"`                   // Status of the post
	CreatedPerson   *uint      `json:"createdPerson" form:"createdPerson"`     // Created person ID
	ModifiedPerson  *uint      `json:"modifiedPerson" form:"modifiedPerson"`   // Modified person ID
	DeleteAt        *time.Time `json:"deleteAt" form:"deleteAt"`               // Deletion timestamp
	ModifiedByAdmin uint8      `json:"modifiedByAdmin" form:"modifiedByAdmin"` // Whether modified by admin
}

type PostUpdateInput struct {
	Id              uint       `json:"id" form:"id"`                           // ID of the post
	Title           string     `json:"title" form:"title"`                     // Title of the post
	Content         string     `json:"content" form:"content"`                 // Content of the post
	Author          string     `json:"author" form:"author"`                   // Author of the post
	Status          uint8      `json:"status" form:"status"`                   // Status of the post
	CreatedPerson   *uint      `json:"createdPerson" form:"createdPerson"`     // Created person ID
	ModifiedPerson  *uint      `json:"modifiedPerson" form:"modifiedPerson"`   // Modified person ID
	DeleteAt        *time.Time `json:"deleteAt" form:"deleteAt"`               // Deletion timestamp
	ModifiedByAdmin uint8      `json:"modifiedByAdmin" form:"modifiedByAdmin"` // Whether modified by admin
}
