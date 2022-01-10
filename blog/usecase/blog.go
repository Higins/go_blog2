package blogUsecase

import (
	"encoding/json"
)

struct BlogUsecase struct {
	gin *gin.Context
	blog *BlogRespository
}

func NewBlogUsecase(gin *gin.Context,blog *BlogRespository) *NewBlogUsecase {
	return &NewBlogUsecase{
		gin:gin,
		blog:blog,
	}
}

func (u *BlogUsecase) Save(gin *gin.Context) {
	saveblogstruct := &u.blog.Blog{
		Title: c.Param("title"),
		Body: c.Param("body"),
	} 
	u.blog.Save(saveblogstruct)

}
func (u *BlogUsecase) Update(gin *gin.Context) {
	saveblogstruct := &u.blog.Blog{
		Title: c.Param("title"),
		Body: c.Param("body"),
	} 
	u.blog.Update(saveblogstruct)

}
func (u *BlogUsecase) AllBlog() {
	return c.JSON(u.blog.AllBlog())
}