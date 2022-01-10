package blogRespository

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:200" json:"title"`
	Body      string    `gorm:"size:3000" json:"body" `
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	db        *gorm.DB
}

func New(db *gorm.DB) *Blog {
	return &Blog{
		db: db,
	}
}

func (b *Blog) Save(post Blog) error {
	return b.db.DB.Create(&post).Error
}
func (b *Blog) FindAll(post Blog, keyword string) (*[]Blog, int64, error) {
	var posts []Blog
	var totalRows int64 = 0

	queryBuider := b.db.DB.Order("created_at desc").Model(&Blog{})

	// Search parameter
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuider = queryBuider.Where(
			b.db.DB.Where("post.title LIKE ? ", queryKeyword))
	}

	err := queryBuider.
		Where(post).
		Find(&posts).
		Count(&totalRows).Error
	return &posts, totalRows, err
}
func (b *Blog) Update(post Blog) error {
	return b.db.DB.Save(&post).Error
}
