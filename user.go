package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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
			sessionKey := req.Header.Get(scheme.Name)
			var sess Session
			sess.Key = sessionKey
			result := sess.Get()

			if result == false {
				fmt.Println("error ------ ")
				goa.LogInfo(ctx, "failed api key auth")
				return ErrUnauthorized("missing auth")
			}
			// user is authenticated, api key is valid
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
	res := &app.Userdataresponse{}
	scheme := app.NewAPIKeySecurity()
	sessionKey := ctx.Request.Header.Get(scheme.Name)

	var sess Session
	sess.Key = sessionKey
	result := sess.Get()
	if result == false {
		res = &app.Userdataresponse{
			Code:    "failure",
			Status:  404,
			Message: "Invalid " + sessionKey,
			Data:    userPayload,
		}
		return ctx.OK(res)
	}

	userID := sess.UserID

	if db.Where("id = ?", userID).First(&user).RecordNotFound() == true {
		// userPayload = &app.UserPayload{}
		res = &app.Userdataresponse{
			Code:    "failure",
			Status:  404,
			Message: "Not found",
			Data:    userPayload,
		}
	} else {
		userPayload = &app.UserPayload{
			UserID:    int(user.ID),
			UserName:  user.Name,
			UserEmail: user.Email,
		}
		res = &app.Userdataresponse{
			Code:    "success",
			Status:  200,
			Message: "",
			Data:    userPayload,
		}
	}

	return ctx.OK(res)
}

// Login runs the login action.
func (c *UserController) Login(ctx *app.LoginUserContext) error {
	//err := ctx.ParseForm()
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if email == "" || password == "" {
		res := &app.Loginresponse{
			Code:      "failure",
			Message:   "Email and Password should not be blanked",
			Status:    400,
			SecretKey: "",
		}
		return ctx.OK(res)
	}
	var user User

	passwordMd5Byte := md5.Sum([]byte(password))
	md5Password := hex.EncodeToString(passwordMd5Byte[:])

	if db.Where("email = ?", email).First(&user).RecordNotFound() == true || user.Password != md5Password {
		res := &app.Loginresponse{
			Code:      "invalid",
			Message:   "Invalid Email or Password",
			Status:    404,
			SecretKey: "",
		}
		return ctx.OK(res)
	}

	var session Session
	session.UserID = user.ID
	session.UserName = user.Name
	sessionToken := session.Register()

	fmt.Println("----------------", user, "-------", sessionToken)

	res := &app.Loginresponse{
		Code:      "success",
		SecretKey: sessionToken,
		Message:   "",
		Status:    200,
	}
	return ctx.OK(res)
}

// Logout runs the logout action.
func (c *UserController) Logout(ctx *app.LogoutUserContext) error {
	// UserController_Logout: start_implement

	// Put your logic here

	// UserController_Logout: end_implement
	//res := &app.Usercommonresponse{}
	//return ctx.OK([]byte("res"))
	scheme := app.NewAPIKeySecurity()
	sessionKey := ctx.Request.Header.Get(scheme.Name)
	var session Session
	id := sessStore[sessionKey].UserID
	db.Where("user_id = ?", id).Delete(&Session{})
	if session.Destroy() == true {
		res := &app.Usercommonresponse{
			Code:    "success",
			Message: "Loggedout successfully",
			Status:  200,
		}
		return ctx.OK(res)
	}
	return nil
}

// Register runs the register action.
func (c *UserController) Register(ctx *app.RegisterUserContext) error {
	// UserController_Register: start_implement

	// Put your logic here

	// UserController_Register: end_implement

	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	re_password := ctx.FormValue("re_password")
	passwordMd5Byte := md5.Sum([]byte(password))
	md5Password := hex.EncodeToString(passwordMd5Byte[:])
	if password != re_password {
		res := &app.Usercommonresponse{
			Code:    "invalid",
			Message: "Re_password and password doesn't match",
			Status:  404,
		}
		return ctx.OK(res)
	}
	user := User{
		Name:     name,
		Email:    email,
		Password: md5Password,
	}
	db.Create(&user)
	res := &app.Usercommonresponse{
		Code:    "success",
		Message: "Registered successfully",
		Status:  200,
	}
	return ctx.OK(res)
}
