package commentUsecase

import "github.com/Higins/go_blog2/domain"

type commentUsecase struct {
	commentRepository domain.CommentRepository
}

func NewCommentUsecase(comment domain.CommentRepository) domain.CommentUsecase {
	return &commentUsecase{
		commentRepository: comment,
	}
}

func (cu *commentUsecase) SaveComment(comment domain.CommentApi) error {
	var commentDb domain.Comment

	commentDb.Text = comment.Text
	commentDb.BlogId = comment.BlogId
	_, err := cu.commentRepository.Save(commentDb)
	if err != nil {
		return err
	}
	return nil
}
