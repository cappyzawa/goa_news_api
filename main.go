//go:generate goagen bootstrap -d news-api/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"news-api/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	// Create service
	service := goa.New("news api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	db, err := gorm.Open("mysql", "news:News_920921@/news?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		service.LogError("startup", "err", err)
	}

	// Mount "articles" controller
	c := NewArticlesController(service, db)
	app.MountArticlesController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":18080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
