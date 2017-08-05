//go:generate goagen bootstrap -d news-api/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"news-api/app"
)

func main() {
	// Create service
	service := goa.New("news api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "articles" controller
	c := NewArticlesController(service)
	app.MountArticlesController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServeTLS(":18080", "cert.pem", "key.pem"); err != nil {
		service.LogError("startup", "err", err)
	}

}
