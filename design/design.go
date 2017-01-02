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

var _ = API("goaAPIKeyUserAuth", func() {
	Title("goaAPIKeyUserAuth")
})

var _ = Resource("user", func() {
	Description("This resource uses an API key to secure its endpoints")
	DefaultMedia(RegisterUser)
	Security(APIKey)

	Action("register", func() {
		Description("This endpoint will be used to register a user")
		Routing(POST("/user/register"))
		/*	Params(func() {
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
		})*/
		NoSecurity()
		Response(OK, UserCommonResponseMedia)
	})

	Action("login", func() {
		Description("This action does not require auth")
		Routing(POST("/user/login"))
		NoSecurity()
		Response(OK, UserLoginResponseMedia)
		Response(Unauthorized)
	})

	Action("info", func() {
		Description("This action is secure with the api_key scheme")
		Routing(GET("/user/info"))
		Response(OK, UserDataResponse)
		Response(Unauthorized)
	})

	Action("logout", func() {
		Description("This action is secure with the api_key scheme")
		Routing(GET("/user/logout"))
		Response(OK, UserCommonResponseMedia)
		Response(Unauthorized)
	})
})

// UserType : contain information of user
var UserType = Type("UserPayload", func() {
	Description("User data payload")
	Attribute("UserId", Integer, "User Id")
	Attribute("UserName", String, "User Name")
	Attribute("UserEmail", String, "User Email")
	Required("UserId", "UserName", "UserEmail")
})

//UserDataResponse : this will be output for user info api call
var UserDataResponse = MediaType("application/vnd.userDataResponse", func() {
	Description("This will be response when user/info endpoint called")
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

// UserCommonResponseMedia : This will be a common response for most the user end points
var UserCommonResponseMedia = MediaType("application/vnd.userCommonResponse", func() {
	Description("This will be a common response for most the user end points")
	Attributes(func() {
		Attribute("code", String, "Response code")
		Attribute("message", String, "Response Message")
		Attribute("status", Integer, "Error Code")
		Required("code", "message", "status")
	})

	View("default", func() {
		Attribute("code")
		Attribute("message")
		Attribute("status")
	})
})

var RegisterUser = MediaType("application/vnd.registerUser", func() {
	Description("This will be a common response for most the user end points")
	Attributes(func() {
		Attribute("name", String, "Input name")
		Attribute("email", String, "Input email")
		Attribute("password", String, "Input password")
		Attribute("re_password", String, "Re-enter password")
		Required("name", "email", "password", "re_password")
	})

	View("default", func() {
		Attribute("name")
		Attribute("email")
		Attribute("password")
		Attribute("re_password")
	})
})
