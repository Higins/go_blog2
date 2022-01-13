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
// Létrehozzuk a router példányt, ami pointert ad vissza (nem interface-t)
func NewRouter(blogUseCase domain.BlogUseCase) *Router {
	return &Router{
		blogUseCase: blogUseCase,
	}
}
// Ezzel a függvénnyel készítjük el a végpontokat
// Gin engine pointert adunk vissza (nézd meg a main.go 34. sorát!)
func (r *Router) InitApi() *gin.Engine {
	server := gin.New()
	server.GET("/", r.GetBlogs)
	server.POST("/saveblog", r.SaveBlog)
	return server
}
// Ez a függvény felelős azért, hogy elmentsük a blogot.
func (r *Router) SaveBlog(c *gin.Context) {
	var blog domain.BlogApi
	// A bejövő hívás body-jában lévő json adatokat bindoljuk a blog változóra
	// Ha nem tudja a rendszer valamiért megcsinálni, hibát dobunk
	err := c.BindJSON(&blog)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Meghívjuk a blog usecase idevágó függvényét
	err = r.blogUseCase.SaveBlog(blog)
	if err != nil {
		fmt.Println(err)
		// Ha hibára fut, hibával térünk vissza. Itt lehet cizellálni, hogy adott hiba típus milyen http hibakódot eredményezzen
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Ha minden OK, 200 a státusz, létrehoztuk/mentettük a blogot
	c.Status(http.StatusOK)
	return
}

func (r *Router) GetBlogs(c *gin.Context) {
	// Kikérjük az összes létező blogot
	blogs, err := r.blogUseCase.AllBlog()
	if err != nil {
		// Ha hibára fut, hibával térünk vissza. Itt lehet cizellálni, hogy adott hiba típus milyen http hibakódot eredményezzen
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Ha minden OK, a státusz 200, és JSON adatokkal térünk vissza
	c.JSON(http.StatusOK, blogs)
	return
}
