package blogUsecase

import (
	"github.com/Higins/go_blog2/domain"
)

type blogUsecase struct {
	blogRepository domain.BlogRepository
}

func NewBlogUsecase(blog domain.BlogRepository) domain.BlogUseCase {
	return &blogUsecase{
		blogRepository: blog,
	}
}

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

func (u *blogUsecase) AllBlog() (blogs []domain.BlogApi, err error) {
	allBlogs, err := u.blogRepository.FindAll()
	if err != nil {
		return nil, err
	}
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