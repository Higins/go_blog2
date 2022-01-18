package domain

import (
	"gorm.io/gorm"
)

// A blog DB adatstruktúra
type Blog struct {
	gorm.Model
	Title   string
	Body    string
	Comment []Comment
}

// blog API struktúra
type BlogApi struct {
	ID    int
	Title string `json:"title"`
	Body  string `json:"body"`
}

// blog repository interface
type BlogRepository interface {
	Save(post Blog) (Blog, error)
	FindAll() (blogs []Blog, err error)
	GetBlogById(blogId int) (blog Blog, err error)
}

// blog usecase interface
type BlogUseCase interface {
	SaveBlog(blog BlogApi) error
	AllBlog() (blogs []BlogApi, err error)
}
