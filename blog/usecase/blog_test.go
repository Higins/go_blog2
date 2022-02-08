package BlogUsecase

import (
	"testing"

	"github.com/Higins/go_blog2/domain"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	newBlog := domain.Blog{
		Title: "Test tájtül",
		Body:  "Test bódi",
	}
	saveBlog := new(domain.MockBlogRepository)

	saveBlog.On("Save", newBlog).Return(domain.Blog{Title: "Test tájtül", Body: "Test bódi"}, nil)
	saveBlog.On("GetBlogById", 1).Return(true, nil)
	blog := NewBlogUsecase(&domain.MockBlogRepository{})
	blog.SaveBlog(domain.BlogApi{Title: "Test tájtül", Body: "Test bódi"})

	assert := assert.New(t)
	assert.Equal(domain.BlogApi{Title: "Test tájtül", Body: "Test bódi"}, blog, "ok")

}

func TestFindAll(t *testing.T) {
	findAllBlog := new(domain.MockBlogRepository)

	findAllBlog.On("FindAll").Return(domain.Blog{}, nil)
}

func TestGetBlogById(t *testing.T) {
	getBlogById := new(domain.MockBlogRepository)

	getBlogById.On("GetBlogById", 1).Return(domain.Blog{}, nil)

}
