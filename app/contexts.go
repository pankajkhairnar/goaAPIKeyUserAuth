//************************************************************************//
// unnamed API: Application Contexts
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
	"strconv"
	"unicode/utf8"
)

// InfoUserContext provides the user info action context.
type InfoUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewInfoUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller info action.
func NewInfoUserContext(ctx context.Context, service *goa.Service) (*InfoUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := InfoUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *InfoUserContext) OK(r *Goaapikeyuserauth) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goaapikeyuserauth")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *InfoUserContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// LoginUserContext provides the user login action context.
type LoginUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Email    *string
	Password *string
}

// NewLoginUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller login action.
func NewLoginUserContext(ctx context.Context, service *goa.Service) (*LoginUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := LoginUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramEmail := req.Params["email"]
	if len(paramEmail) > 0 {
		rawEmail := paramEmail[0]
		rctx.Email = &rawEmail
	}
	paramPassword := req.Params["password"]
	if len(paramPassword) > 0 {
		rawPassword := paramPassword[0]
		rctx.Password = &rawPassword
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginUserContext) OK(r *Goaapikeyuserauth) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goaapikeyuserauth")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// RegisterUserContext provides the user register action context.
type RegisterUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Email      *string
	Name       *string
	Password   *string
	RePassword *string
}

// NewRegisterUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller register action.
func NewRegisterUserContext(ctx context.Context, service *goa.Service) (*RegisterUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RegisterUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramEmail := req.Params["email"]
	if len(paramEmail) > 0 {
		rawEmail := paramEmail[0]
		rctx.Email = &rawEmail
		if rctx.Email != nil {
			if utf8.RuneCountInString(*rctx.Email) < 5 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`email`, *rctx.Email, utf8.RuneCountInString(*rctx.Email), 5, true))
			}
		}
		if rctx.Email != nil {
			if utf8.RuneCountInString(*rctx.Email) > 50 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`email`, *rctx.Email, utf8.RuneCountInString(*rctx.Email), 50, false))
			}
		}
	}
	paramName := req.Params["name"]
	if len(paramName) > 0 {
		rawName := paramName[0]
		rctx.Name = &rawName
		if rctx.Name != nil {
			if ok := goa.ValidatePattern(`^[ a-zA-Z]+$`, *rctx.Name); !ok {
				err = goa.MergeErrors(err, goa.InvalidPatternError(`name`, *rctx.Name, `^[ a-zA-Z]+$`))
			}
		}
		if rctx.Name != nil {
			if utf8.RuneCountInString(*rctx.Name) < 2 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, *rctx.Name, utf8.RuneCountInString(*rctx.Name), 2, true))
			}
		}
		if rctx.Name != nil {
			if utf8.RuneCountInString(*rctx.Name) > 50 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, *rctx.Name, utf8.RuneCountInString(*rctx.Name), 50, false))
			}
		}
	}
	paramPassword := req.Params["password"]
	if len(paramPassword) > 0 {
		rawPassword := paramPassword[0]
		rctx.Password = &rawPassword
		if rctx.Password != nil {
			if utf8.RuneCountInString(*rctx.Password) < 5 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`password`, *rctx.Password, utf8.RuneCountInString(*rctx.Password), 5, true))
			}
		}
		if rctx.Password != nil {
			if utf8.RuneCountInString(*rctx.Password) > 20 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`password`, *rctx.Password, utf8.RuneCountInString(*rctx.Password), 20, false))
			}
		}
	}
	paramRePassword := req.Params["re_password"]
	if len(paramRePassword) > 0 {
		rawRePassword := paramRePassword[0]
		rctx.RePassword = &rawRePassword
		if rctx.RePassword != nil {
			if utf8.RuneCountInString(*rctx.RePassword) < 5 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`re_password`, *rctx.RePassword, utf8.RuneCountInString(*rctx.RePassword), 5, true))
			}
		}
		if rctx.RePassword != nil {
			if utf8.RuneCountInString(*rctx.RePassword) > 20 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`re_password`, *rctx.RePassword, utf8.RuneCountInString(*rctx.RePassword), 20, false))
			}
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *RegisterUserContext) OK(r *Goaapikeyuserauth) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goaapikeyuserauth")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
