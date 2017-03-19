package models

import (
	"time"
)

type Post struct {
	ID          int       `json:"id" gorm:"primary_key;column:id"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Content     string    `json:"content" gorm:"column:content"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at"`

	User User `gorm:"ForeignKey:user_id"`
}

func (p *Post) TableName() string {
	return "post"
}

func (p Post) CreateDate() string {
	return p.CreatedAt.Format("2006-01-02 03:04")
}
