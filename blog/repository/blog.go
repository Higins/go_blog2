package blogRespository

import (
	"fmt"
	"github.com/Higins/go_blog2/domain"
	"gorm.io/gorm"
)

type blogRepository struct {
	db    *gorm.DB
}

func NewBlogRepository(db *gorm.DB) domain.BlogRepository {
	return &blogRepository{
		db: db,
	}
}

func (b *blogRepository) Save(post domain.Blog) (domain.Blog, error) {
	var err error
	if post.ID > 0 {
		err = b.db.Create(&post).Error
	} else {
		err = b.db.Save(&post).Error
	}
	if err != nil {
		fmt.Println(err)
		return domain.Blog{}, err
	}
	return post, nil
}
func (b *blogRepository) FindAll() (blogs []domain.Blog, err error) {
	err = b.db.Find(&blogs).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return blogs, nil
}

func (b *blogRepository) GetBlogById(blogId int) (blog domain.Blog, err error) {
	err = b.db.Where("id = ?", blogId).First(&blog).Error
	if err != nil {
		fmt.Println(err)
		return domain.Blog{}, err
	}
	return blog, nil
}
