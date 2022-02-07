package blogRespository

import (
	"testing"

	"github.com/Higins/go_blog2/domain"
)

func TestSave(t *testing.T) {
	newBlog := domain.Blog{
		Title: "Test tájtül",
		Body:  "Test bódi",
	}
	saveBlog := new(domain.MockBlogRepository)

	saveBlog.On("Save", newBlog).Return(domain.Blog{Title: "Test tájtül", Body: "Test bódi"}, nil)
}

func TestFindAll(t *testing.T) {
	findAllBlog := new(domain.MockBlogRepository)

	findAllBlog.On("FindAll").Return(domain.Blog{}, nil)
}

func TestGetBlogById(t *testing.T) {
	getBlogById := new(domain.MockBlogRepository)

	getBlogById.On("GetBlogById", 1).Return(domain.Blog{}, nil)

}
