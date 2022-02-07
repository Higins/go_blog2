package commentRepository

import (
	"testing"

	"github.com/Higins/go_blog2/domain"
)

func TestSave(t *testing.T) {
	newComment := domain.Comment{
		Text:   "test kommentttt",
		BlogId: 1,
	}
	findAllBlog := new(domain.MockCommentRepository)

	findAllBlog.On("Save", newComment).Return(domain.Comment{Text: "test kommentttt", BlogId: 1}, nil)
}
