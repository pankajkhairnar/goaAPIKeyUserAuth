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

package client

import "github.com/goadesign/goa"

// User data payload
type userPayload struct {
	// User Email
	UserEmail *string `form:"UserEmail,omitempty" json:"UserEmail,omitempty" xml:"UserEmail,omitempty"`
	// User Id
	UserID *int `form:"UserId,omitempty" json:"UserId,omitempty" xml:"UserId,omitempty"`
	// User Name
	UserName *string `form:"UserName,omitempty" json:"UserName,omitempty" xml:"UserName,omitempty"`
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "UserId"))
	}
	if ut.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "UserName"))
	}
	if ut.UserEmail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "UserEmail"))
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.UserEmail != nil {
		pub.UserEmail = *ut.UserEmail
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	if ut.UserName != nil {
		pub.UserName = *ut.UserName
	}
	return &pub
}

// User data payload
type UserPayload struct {
	// User Email
	UserEmail string `form:"UserEmail" json:"UserEmail" xml:"UserEmail"`
	// User Id
	UserID int `form:"UserId" json:"UserId" xml:"UserId"`
	// User Name
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {

	if ut.UserName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "UserName"))
	}
	if ut.UserEmail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "UserEmail"))
	}
	return
}
