package blogUsecase

import (
	"github.com/Higins/go_blog2/domain"
)

type blogUsecase struct {
	blogRepository domain.BlogRepository
}
// blog usecase példány létrehozása (constructor) függvény, ami interface-t ad vissza
// Ha van különbség a blogUsecase struct-ra ráaggatott publikus függvények és az interface között
// Az IDE is sipákol!
func NewBlogUsecase(blog domain.BlogRepository) domain.BlogUseCase {
	return &blogUsecase{
		blogRepository: blog,
	}
}
// Az új blog létrehozása és egy blog szerkesztését érdemes egy függvénybe tenni
func (u *blogUsecase) SaveBlog(blog domain.BlogApi) error {
	var blogDb domain.Blog
	if blog.ID > 0 {
		blogDb, err := u.blogRepository.GetBlogById(blog.ID)
		if err != nil || blogDb.ID <= 0 {
			return err
		}
	}
	blogDb.Title = blog.Title
	blogDb.Body = blog.Body
	_, err := u.blogRepository.Save(blogDb)
	if err != nil {
		return err
	}
	return nil
}
// Itt API-s blog típust adunk vissza
func (u *blogUsecase) AllBlog() (blogs []domain.BlogApi, err error) {
	allBlogs, err := u.blogRepository.FindAll()
	if err != nil {
		return nil, err
	}
	// Átkonvertáljuk a DB modelt API modellé
	for _, b := range allBlogs {
		blogApi := domain.BlogApi{
			ID: int(b.ID),
			Body: b.Body,
			Title: b.Title,
		}
		blogs = append(blogs, blogApi)
	}
	return blogs, nil
}