package main

import (
	"goaAPIKeyUserAuth/app"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
}

// User : strut for user model
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

// NewAPIKeyMiddleware creates a middleware that checks for the presence of an authorization header
// and validates its content.
func NewAPIKeyMiddleware() goa.Middleware {
	// Instantiate API Key security scheme details generated from design
	scheme := app.NewAPIKeySecurity()

	// Middleware
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// Retrieve and log header specified by scheme
			key := req.Header.Get(scheme.Name)

			// A real app would do something more interesting here
			if len(key) == 0 || key == "Bearer" {
				goa.LogInfo(ctx, "failed api key auth")
				return ErrUnauthorized("missing auth")
			}
			// Proceed.
			goa.LogInfo(ctx, "auth", "apikey", "key", key, "scheme.Name", scheme.Name)
			return h(ctx, rw, req)
		}
	}
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Info runs the info action.
func (c *UserController) Info(ctx *app.InfoUserContext) error {
	var user User
	var userPayload *app.UserPayload

	user.ID = uint(ctx.ID)
	db.First(&user)

	userPayload = &app.UserPayload{
		UserID:   ctx.ID,
		UserName: user.Name,
	}

	res := &app.Goaapikeyuserauth{
		Code:    "success",
		Status:  200,
		Message: "",
		Data:    userPayload,
	}

	return ctx.OK(res)
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	res := &app.Goaapikeyuserauth{}
	return ctx.OK(res)
}

// Register runs the register action.
func (c *UserController) Register(ctx *app.RegisterUserContext) error {

	res := &app.Goaapikeyuserauth{}
	return ctx.OK(res)
}
