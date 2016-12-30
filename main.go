//go:generate goagen bootstrap -d goaAPIKeyUserAuth/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"goaAPIKeyUserAuth/app"
)

func main() {
	// Create service
	service := goa.New("")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "user" controller
	c := NewUserController(service)
	app.MountUserController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
