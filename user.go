package main

import (
	"goaAPIKeyUserAuth/app"
	"log"

	"github.com/goadesign/goa"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Info runs the info action.
func (c *UserController) Info(ctx *app.InfoUserContext) error {
	// UserController_Info: start_implement

	// Put your logic here

	// UserController_Info: end_implement
	res := &app.Userdataresponse{}
	return ctx.OK(res)
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	// UserController_Login: start_implement

	// Put your logic here

	log.Println("----------------", ctx.Payload.Email)

	// UserController_Login: end_implement
	res := &app.Loginresponse{}
	return ctx.OK(res)
}

// Logout runs the logout action.
func (c *UserController) Logout(ctx *app.LogoutUserContext) error {
	// UserController_Logout: start_implement

	// Put your logic here

	// UserController_Logout: end_implement
	res := &app.Usercommonresponse{}
	return ctx.OK(res)
}

// Register runs the register action.
func (c *UserController) Register(ctx *app.RegisterUserContext) error {
	// UserController_Register: start_implement

	// Put your logic here

	// UserController_Register: end_implement
	res := &app.Usercommonresponse{}
	return ctx.OK(res)
}
