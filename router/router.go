package router

import (
	"fmt"
	"github.com/Higins/go_blog2/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	blogUseCase domain.BlogUseCase
}

func NewRouter(blogUseCase domain.BlogUseCase) *Router {
	return &Router{
		blogUseCase: blogUseCase,
	}
}

func (r *Router) InitApi() *gin.Engine {
	server := gin.New()
	server.GET("/", r.GetBlogs)
	server.POST("/saveblog", r.SaveBlog)
	return server
}

func (r *Router) SaveBlog(c *gin.Context) {
	var blog domain.BlogApi
	err := c.BindJSON(&blog)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = r.blogUseCase.SaveBlog(blog)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusOK)
	return
}

func (r *Router) GetBlogs(c *gin.Context) {
	blogs, err := r.blogUseCase.AllBlog()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, blogs)
	return
}
