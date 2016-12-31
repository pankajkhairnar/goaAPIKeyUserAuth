//************************************************************************//
// unnamed API: Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=goaAPIKeyUserAuth/design
// --out=$(GOPATH)/src/goaAPIKeyUserAuth
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Info(*InfoUserContext) error
	Login(*LoginUserContext) error
	Logout(*LogoutUserContext) error
	Register(*RegisterUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewInfoUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Info(rctx)
	}
	h = handleSecurity("api_key", h)
	service.Mux.Handle("GET", "/user/info", ctrl.MuxHandler("Info", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Info", "route", "GET /user/info", "security", "api_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Login(rctx)
	}
	service.Mux.Handle("POST", "/user/login", ctrl.MuxHandler("Login", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Login", "route", "POST /user/login")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLogoutUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Logout(rctx)
	}
	h = handleSecurity("api_key", h)
	service.Mux.Handle("GET", "/user/logout", ctrl.MuxHandler("Logout", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Logout", "route", "GET /user/logout", "security", "api_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRegisterUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Register(rctx)
	}
	service.Mux.Handle("POST", "/user/register", ctrl.MuxHandler("Register", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Register", "route", "POST /user/register")
}
