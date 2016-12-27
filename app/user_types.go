//************************************************************************//
// unnamed API: Application User Types
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

import "github.com/goadesign/goa"

// userPayload user type.
type userPayload struct {
	// University Name
	UserName *string `form:"UserName,omitempty" json:"UserName,omitempty" xml:"UserName,omitempty"`
	// University Id
	UserID *int `form:"userId,omitempty" json:"userId,omitempty" xml:"userId,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "userId"))
	}
	if ut.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "UserName"))
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.UserName != nil {
		pub.UserName = *ut.UserName
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// UserPayload user type.
type UserPayload struct {
	// University Name
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// University Id
	UserID int `form:"userId" json:"userId" xml:"userId"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {

	if ut.UserName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "UserName"))
	}
	return
}
