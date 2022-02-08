package commentUsecase

import (
	"testing"

	"github.com/Higins/go_blog2/domain"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	newComment := domain.Comment{
		Text:   "test kommentttt",
		BlogId: 1,
	}
	findAllBlog := new(domain.MockCommentRepository)

	findAllBlog.On("Save", newComment).Return(domain.Comment{Text: "test kommentttt", BlogId: 1}, nil)
	comment := NewCommentUsecase(&domain.MockCommentRepository{})
	comment.SaveComment(domain.CommentApi{Text: "test kommentttt",BlogId: 1})

	assert := assert.New(t)
	assert.Equal(domain.CommentApi{Text: "test kommentttt",BlogId: 1},comment,"ok")
}
