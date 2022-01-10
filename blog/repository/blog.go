package blogRespository

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
	db    *gorm.DB
}

func New(db *gorm.DB) *Blog {
	return &Blog{
		db: db,
	}
}

func (b *Blog) Save(post Blog) error {
	return b.db.DB.Create(&post).Error
}
func (b *Blog) FindAll() *Blog {
	blog := make([]Blog, 0)
	b.db.Find(&blog)
	return blog
}
func (b *Blog) Update(post Blog) error {
	return b.db.DB.Save(&post).Error
}
