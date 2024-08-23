package response

import "time"

type PostQueryView struct {
	ID       uint      `json:"id" form:"id"`
	Title    string    `json:"title" form:"title"`
	Content  string    `json:"content" form:"content"`
	Author   string    `json:"author" form:"author"`
	CreateAt time.Time `json:"create_at" form:"create_at"`
}
