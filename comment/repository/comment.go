package commentRepository

import (
	"fmt"

	"github.com/Higins/go_blog2/domain"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cp *commentRepository) Save(comment domain.Comment) (domain.Comment, error) {
	var err error
	if comment.ID > 0 {
		err = cp.db.Create(&comment).Error
	} else {
		err = cp.db.Save(&comment).Error
	}
	if err != nil {
		fmt.Println(err)
		return domain.Comment{}, err
	}
	return comment, nil
}
