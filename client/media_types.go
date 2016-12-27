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

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// The common media type to all request responses for user service (default view)
//
// Identifier: application/vnd.goaapikeyuserauth; view=default
type Goaapikeyuserauth struct {
	// Response code
	Code string `form:"code" json:"code" xml:"code"`
	// User Info
	Data *UserPayload `form:"data" json:"data" xml:"data"`
	// Response Message
	Message string `form:"message" json:"message" xml:"message"`
	// Error Code
	Status int `form:"status" json:"status" xml:"status"`
}

// Validate validates the Goaapikeyuserauth media type instance.
func (mt *Goaapikeyuserauth) Validate() (err error) {
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

// DecodeGoaapikeyuserauth decodes the Goaapikeyuserauth instance encoded in resp body.
func (c *Client) DecodeGoaapikeyuserauth(resp *http.Response) (*Goaapikeyuserauth, error) {
	var decoded Goaapikeyuserauth
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

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

// DecodeLoginresponse decodes the Loginresponse instance encoded in resp body.
func (c *Client) DecodeLoginresponse(resp *http.Response) (*Loginresponse, error) {
	var decoded Loginresponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
