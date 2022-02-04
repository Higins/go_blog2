package domain

import (
	"github.com/stretchr/testify/mock"
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

type MockCommentRepository struct {
	mock.Mock
}

func (m *MockCommentRepository) Save(comment Comment) (Comment, error) {
	args := m.Called(comment)
	return args.Get(0).(Comment), args.Error(1)
}
