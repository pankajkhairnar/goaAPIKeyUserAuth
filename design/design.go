package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// APIKey defines a security scheme using an API key (shared secret).  The scheme uses the
// "secret-key" header to lookup the key.
var APIKey = APIKeySecurity("api_key", func() {
	Header("secret-key")
})

var _ = Resource("user", func() {
	Description("This resource uses an API key to secure its endpoints")
	DefaultMedia(UserResponseMedia)
	Security(APIKey)

	Action("register", func() {
		Description("This endpoint will be used to register a user")
		Routing(POST("/user/register"))
		Params(func() {
			Param("name", String, "Name of user", func() {
				MinLength(2)
				MaxLength(50)
				Pattern("^[ a-zA-Z]+$")
			})

			Param("email", String, "Email of user", func() {
				MinLength(5)
				MaxLength(50)
			})

			Param("password", String, "Password", func() {
				MinLength(5)
				MaxLength(20)
			})

			Param("re_password", String, "Password", func() {
				MinLength(5)
				MaxLength(20)
			})
		})
		NoSecurity()
		Response(OK)
	})

	Action("login", func() {
		Description("This action does not require auth")
		Routing(POST("/user/login"))
		Params(func() {
			Param("email", String, "Email ID of user")
			Param("password", String, "Password")
		})
		NoSecurity()
		Response(OK)
	})

	Action("info", func() {
		Description("This action is secure with the api_key scheme")
		Routing(GET("/user/info/:id"))
		Params(func() {
			Param("id", Integer, "UserId")
		})
		Response(OK)
		Response(Unauthorized)
	})
})

// UserType : contain information of user
var UserType = Type("UserPayload", func() {
	Description("")
	Attribute("UserId", Integer, "User Id")
	Attribute("UserName", String, "User Name")
	Required("userId", "UserName")
})

//UserResponseMedia : this will be response skeleton for each user api endpoints
var UserResponseMedia = MediaType("application/vnd.goaAPIKeyUserAuth", func() {
	Description("The common media type to all request responses for user service")
	Attributes(func() {
		Attribute("code", String, "Response code")
		Attribute("message", String, "Response Message")
		Attribute("status", Integer, "Error Code")
		Attribute("data", UserType, "User Info")
		Required("code", "message", "status", "data")
	})

	View("default", func() {
		Attribute("code")
		Attribute("message")
		Attribute("status")
		Attribute("data")
	})
})

// UserLoginResponseMedia : if user logged in successfully secret_key will be sent else error
var UserLoginResponseMedia = MediaType("application/vnd.loginResponse", func() {
	Description("After login response")
	Attributes(func() {
		Attribute("code", String, "Response code")
		Attribute("message", String, "Response Message")
		Attribute("status", Integer, "Error Code")
		Attribute("secret_key", String, "Secret key")
		Required("code", "message", "status", "secret_key")
	})

	View("default", func() {
		Attribute("code")
		Attribute("message")
		Attribute("status")
		Attribute("secret_key")
	})
})
