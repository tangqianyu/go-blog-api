package models

import (
	"gorm.io/gorm"
	"time"
)

//Post  Model
type Post struct {
	gorm.Model
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:200" json:"title"`
	BookID    int64     `json:"book_id"`
	Book      Book      `json:"book"`
	Body      string    `gorm:"size:3000" json:"body" `
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TableName method that returns tablename of Post model
func (post *Post) TableName() string {
	return "post"
}

//ResponseMap -> response map of Post
func (post *Post) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = post.ID
	resp["title"] = post.Title
	resp["book_name"] = post.Book.Name
	resp["body"] = post.Body
	resp["created_at"] = post.CreatedAt
	resp["updated_at"] = post.UpdatedAt
	return resp

}
