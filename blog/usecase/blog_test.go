package BlogUsecase

import (
	"testing"

	"github.com/Higins/go_blog2/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSave(t *testing.T) {
	oldBlog := domain.Blog{
		Comment: []domain.Comment(nil),
		Title:   "Test tájtül",
		Body:    "Test bódi",
		Model: gorm.Model{
			ID: 1,
		},
	}
	newBlog := domain.Blog{
		Comment: []domain.Comment(nil),
		Title:   "Test tájtül updated",
		Body:    "Test bódi",
		Model: gorm.Model{
			ID: 1,
		},
	}
	saveBlog := domain.MockBlogRepository{}
	saveBlog.On("Save", newBlog).Return(newBlog, nil)
	saveBlog.On("GetBlogById", 1).Return(oldBlog, nil)
	blog := NewBlogUsecase(&saveBlog)
	err := blog.SaveBlog(domain.BlogApi{Title: "Test tájtül updated", Body: "Test bódi", ID: 1})
	assert.Nil(t, err)

}
func TestSaveNewblog(t *testing.T) {
	oldBlog := domain.Blog{
		Comment: []domain.Comment(nil),
		Title:   "Test tájtül new blog",
		Body:    "Test bódi",
		Model: gorm.Model{
			ID: 1,
		},
	}
	newBlog := domain.Blog{
		Comment: []domain.Comment(nil),
		Title:   "Test tájtül new blog",
		Body:    "Test bódi",
	}
	saveBlog := domain.MockBlogRepository{}
	saveBlog.On("Save", newBlog).Return(newBlog, nil)
	saveBlog.On("GetBlogById", 1).Return(oldBlog, nil)
	blog := NewBlogUsecase(&saveBlog)
	err := blog.SaveBlog(domain.BlogApi{Title: "Test tájtül new blog", Body: "Test bódi", ID: 1})
	assert.Nil(t, err)

}
func TestSaveNewblogAndGetBadBlog(t *testing.T) {
	oldBlog := domain.Blog{
		Comment: []domain.Comment(nil),
		Title:   "Test tájtül bad blog id",
		Body:    "Test bódi",
		Model: gorm.Model{
			ID: 2,
		},
	}
	newBlog := domain.Blog{
		Comment: []domain.Comment(nil),
		Title:   "Test tájtül new blog",
		Body:    "Test bódi",
	}
	saveBlog := domain.MockBlogRepository{}
	saveBlog.On("Save", newBlog).Return(newBlog, nil)
	saveBlog.On("GetBlogById", 1).Return(oldBlog, nil)
	blog := NewBlogUsecase(&saveBlog)
	err := blog.SaveBlog(domain.BlogApi{Title: "Test tájtül bad blog id", Body: "Test bódi", ID: 1})
	assert.Nil(t, err)

}
func TestFindAll(t *testing.T) {
	findAllBlog := new(domain.MockBlogRepository)

	findAllBlog.On("FindAll").Return(domain.Blog{}, nil)
}

func TestGetBlogById(t *testing.T) {
	getBlogById := new(domain.MockBlogRepository)

	getBlogById.On("GetBlogById", 1).Return(domain.Blog{}, nil)

}
