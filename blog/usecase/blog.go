package blogUsecase

type blogUsecase struct {
	service blogRespository.Blog
}

func NewBlogController(s blogRespository.Blog) {

}

func (p *blogRespository) AddNewBlog(ctx *gin.Context) {

}
