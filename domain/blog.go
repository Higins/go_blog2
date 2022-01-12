package domain

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title string
	Body string
}

type BlogApi struct {
	ID int
	Title string `json:"title"`
	Body string `json:"body"`
}

type BlogRepository interface {
	Save (post Blog) (Blog, error)
	FindAll() (blogs []Blog, err error)
	GetBlogById(blogId int) (blog Blog, err error)
}

type BlogUseCase interface {
	SaveBlog(blog BlogApi) error
	AllBlog() (blogs []BlogApi, err error)
}
