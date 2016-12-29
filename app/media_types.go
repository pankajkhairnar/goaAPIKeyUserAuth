//************************************************************************//
// unnamed API: Application Media Types
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

// After login response (default view)
//
// Identifier: application/vnd.loginresponse; view=default
type Loginresponse struct {
	// Response code
	Code string `form:"code" json:"code" xml:"code"`
	// Response Message
	Message string `form:"message" json:"message" xml:"message"`
	// Secret key
	SecretKey string `form:"secret_key" json:"secret_key" xml:"secret_key"`
	// Error Code
	Status int `form:"status" json:"status" xml:"status"`
}

// Validate validates the Loginresponse media type instance.
func (mt *Loginresponse) Validate() (err error) {
	if mt.Code == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "code"))
	}
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}

	if mt.SecretKey == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "secret_key"))
	}
	return
}

// This will be a common response for most the user end points (default view)
//
// Identifier: application/vnd.usercommonresponse; view=default
type Usercommonresponse struct {
	// Response code
	Code string `form:"code" json:"code" xml:"code"`
	// Response Message
	Message string `form:"message" json:"message" xml:"message"`
	// Error Code
	Status int `form:"status" json:"status" xml:"status"`
}

// Validate validates the Usercommonresponse media type instance.
func (mt *Usercommonresponse) Validate() (err error) {
	if mt.Code == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "code"))
	}
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}

	return
}

// This will be response when user/info endpoint called (default view)
//
// Identifier: application/vnd.userdataresponse; view=default
type Userdataresponse struct {
	// Response code
	Code string `form:"code" json:"code" xml:"code"`
	// User Info
	Data *UserPayload `form:"data" json:"data" xml:"data"`
	// Response Message
	Message string `form:"message" json:"message" xml:"message"`
	// Error Code
	Status int `form:"status" json:"status" xml:"status"`
}

// Validate validates the Userdataresponse media type instance.
func (mt *Userdataresponse) Validate() (err error) {
	if mt.Code == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "code"))
	}
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}

	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Data != nil {
		if err2 := mt.Data.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
