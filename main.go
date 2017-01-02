//go:generate goagen bootstrap -d goaAPIKeyUserAuth/design

package main

import (
	"goaAPIKeyUserAuth/app"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/jinzhu/gorm"
)

var (
	// ErrUnauthorized is the error returned for unauthorized requests.
	ErrUnauthorized = goa.NewErrorClass("unauthorized", 401)
)

var db *gorm.DB

func main() {
	// Create service
	service := goa.New("")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// middleware
	app.UseAPIKeyMiddleware(service, NewAPIKeyMiddleware())

	// Mount "user" controller
	c := NewUserController(service)
	app.MountUserController(service, c)

	db, _ = gorm.Open("mysql", "root:password@/database")
	LoadSession()
	// Start service
	if err := service.ListenAndServe(":9001"); err != nil {
		service.LogError("startup", "err", err)
	}
}
func LoadSession() {
	session := []Session{}
	if db.Find(&session).RecordNotFound() != true {
		for _, row := range session {
			sessStore[row.Key] = Session{
				Key:          row.Key,
				UserID:       row.UserID,
				UserName:     row.UserName,
				LastAccessed: row.LastAccessed,
			}
		}
	}
}
