package main

import (
	"fmt"
	blogRespository "github.com/Higins/go_blog2/blog/repository"
	blogUsecase "github.com/Higins/go_blog2/blog/usecase"
	"github.com/Higins/go_blog2/domain"
	"github.com/Higins/go_blog2/router"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	gormDB, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	err = gormDB.AutoMigrate(domain.Blog{})
	fmt.Println(err)
	blogRouter := InitServiceWithDependencies(gormDB)
	// create error group to detect failed services
	errorGroup := make(chan error, 1)
	go func() {
		errorGroup <- blogRouter.InitApi().Run(":8080")
	}()
	if err := <-errorGroup; err != nil {
		// fatal trigger exit 1
		log.Fatal(err)
	}
}

func InitServiceWithDependencies(gormDB *gorm.DB) *router.Router {
	blogRespo := blogRespository.NewBlogRepository(gormDB)
	blogUseCase := blogUsecase.NewBlogUsecase(blogRespo)
	blogRouter := router.NewRouter(blogUseCase)
	return blogRouter
}
