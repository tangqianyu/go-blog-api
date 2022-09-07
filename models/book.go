package models

import (
	"time"
)

//Book  Model
type Book struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:200" json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TableName method that returns tablename of Post model
func (book *Book) TableName() string {
	return "book"
}

//ResponseMap -> response map of Post
func (book *Book) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = book.ID
	resp["name"] = book.Name
	resp["created_at"] = book.CreatedAt
	resp["updated_at"] = book.UpdatedAt
	return resp

}
