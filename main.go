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
	// DB kapcsolat példány létrehozása
	gormDB, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		// Ha nem sikerül, nem futhat a rendszer, ezért pánik
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	// Ez egy automigrálás, amivel létrehozzuk, ha nem létezik, az adatbázis táblát
	err = gormDB.AutoMigrate(domain.Blog{})
	fmt.Println(err)
	// Létrehozzuk a router példányt
	blogRouter := InitServiceWithDependencies(gormDB)
	// create error group to detect failed services
	errorGroup := make(chan error, 1)
	// Egy go routine-ban futtatjuk a routert
	go func() {
		// Futtatjuk a routert a 8080-as porton, ha bármi hiba adódik, amit a router (gin engine) elkap, küldjük az error channelbe
		errorGroup <- blogRouter.InitApi().Run(":8080")
	}()
	// Bármiféle hiba kerül az error channelbe, logoljuk
	if err := <-errorGroup; err != nil {
		// fatal trigger exit 1
		log.Fatal(err)
	}
}

func InitServiceWithDependencies(gormDB *gorm.DB) *router.Router {
	// Elkészítjük a blog repository-t, injektáljuk a db kapcsolatot
	blogRespo := blogRespository.NewBlogRepository(gormDB)
	// Elkészítjük a blog usecase-t, injektáljuk a blog repository-t
	blogUseCase := blogUsecase.NewBlogUsecase(blogRespo)
	// Blog usecase-t injektáljuk a routerbe
	blogRouter := router.NewRouter(blogUseCase)
	return blogRouter
}
