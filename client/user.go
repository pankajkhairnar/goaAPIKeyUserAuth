package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// InfoUserPath computes a request path to the info action of user.
func InfoUserPath() string {
	return fmt.Sprintf("/user/info")
}

// This action is secure with the api_key scheme
func (c *Client) InfoUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewInfoUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewInfoUserRequest create the request corresponding to the info action endpoint of the user resource.
func (c *Client) NewInfoUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.APIKeySigner != nil {
		c.APIKeySigner.Sign(req)
	}
	return req, nil
}

// LoginUserPayload is the user login action payload.
type LoginUserPayload struct {
	Email string `form:"email" json:"email" xml:"email"`
}

// LoginUserPath computes a request path to the login action of user.
func LoginUserPath() string {
	return fmt.Sprintf("/user/login")
}

// This action does not require auth
func (c *Client) LoginUser(ctx context.Context, path string, payload *LoginUserPayload, contentType string) (*http.Response, error) {
	req, err := c.NewLoginUserRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginUserRequest create the request corresponding to the login action endpoint of the user resource.
func (c *Client) NewLoginUserRequest(ctx context.Context, path string, payload *LoginUserPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// LogoutUserPath computes a request path to the logout action of user.
func LogoutUserPath() string {
	return fmt.Sprintf("/user/logout")
}

// This action is secure with the api_key scheme
func (c *Client) LogoutUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewLogoutUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLogoutUserRequest create the request corresponding to the logout action endpoint of the user resource.
func (c *Client) NewLogoutUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.APIKeySigner != nil {
		c.APIKeySigner.Sign(req)
	}
	return req, nil
}

// RegisterUserPath computes a request path to the register action of user.
func RegisterUserPath() string {
	return fmt.Sprintf("/user/register")
}

// This endpoint will be used to register a user
func (c *Client) RegisterUser(ctx context.Context, path string, email *string, name *string, password *string, rePassword *string) (*http.Response, error) {
	req, err := c.NewRegisterUserRequest(ctx, path, email, name, password, rePassword)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegisterUserRequest create the request corresponding to the register action endpoint of the user resource.
func (c *Client) NewRegisterUserRequest(ctx context.Context, path string, email *string, name *string, password *string, rePassword *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if email != nil {
		values.Set("email", *email)
	}
	if name != nil {
		values.Set("name", *name)
	}
	if password != nil {
		values.Set("password", *password)
	}
	if rePassword != nil {
		values.Set("re_password", *rePassword)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
