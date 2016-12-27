//go:generate goagen bootstrap -d goaAPIKeyUserAuth/design

package main

import (
	"goaAPIKeyUserAuth/app"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
)

func main() {
	// Create service
	service := goa.New("")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// middleware
	app.UseAPIKeyMiddleware(service, NewAPIKeyMiddleware())

	// Mount "user" controller
	c := NewUserController(service)
	app.MountUserController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
