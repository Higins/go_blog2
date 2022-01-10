package blogUsecase

import (
	"github.com/Higins/go_blog2/*"
)

type blogRespository struct {
	service blogRespository.Blog
}

func NewBlogController(s blogRespository.Blog) blogRespository {
	return blogRespository{
		service: s,
	}
}

func (p *blogRespository) AddNewBlog(ctx *gin.Context) {
    var post models.Post
    ctx.ShouldBindJSON(&post)

    if post.Title == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Title is required")
        return
    }
    if post.Body == "" {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Body is required")
        return
    }
    err := p.service.Save(post)
    if err != nil {
        util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create post")
        return
    }
    util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Post")
}