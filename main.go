package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
	err    error
)

func main() {
	var gormDB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connect database")
	}
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})
	fmt.Println("Connection Opened to Database")

	InitServiceWithDependencies(gormDB)
	r.Run(":8080")
}

func InitServiceWithDependencies(gormDB *gorm.DB) error {

	return nil
}
