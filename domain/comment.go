package domain

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text   string
	BlogId int
}

type CommentApi struct {
	ID     int
	Text   string `json:"text"`
	BlogId int    `json:"blogid"`
}

type CommentRepository interface {
	Save(comment Comment) (Comment, error)
}

type CommentUsecase interface {
	SaveComment(comment CommentApi) error
}
