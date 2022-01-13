package blogRespository

import (
	"fmt"
	"github.com/Higins/go_blog2/domain"
	"gorm.io/gorm"
)

type blogRepository struct {
	db    *gorm.DB
}
// blog repository példány létrehozása (constructor) függvény, ami interface-t ad vissza
// Ha van különbség a blogRepository struct-ra ráaggatott publikus függvények és az interface között
// Az IDE is sipákol!
func NewBlogRepository(db *gorm.DB) domain.BlogRepository {
	return &blogRepository{
		db: db,
	}
}
// Repository mindig a DB adattípusát adja vissza, ezzel kezd valamit a usecase réteg
func (b *blogRepository) Save(post domain.Blog) (domain.Blog, error) {
	var err error
	// Ez így egyértelmű a gorm.io-nak, hogy mit kell kezdenie az adott adattal
	// Van, hogy ha csak Save-et használsz, hibára fut, amikor újat akarsz létrehozni
	// Egy Where feltételre szokott panaszkodni
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
