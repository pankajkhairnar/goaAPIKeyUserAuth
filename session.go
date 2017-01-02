package main

import (
	"time"

	"github.com/twinj/uuid"
)

type Session struct {
	Key          string
	UserID       uint
	UserName     string
	LastAccessed time.Time
	//	DeletedAt    time.Time
}

var sessStore = make(map[string]Session)

// Register : method will create new session id and store session values in map, it will return session tokey
func (sess *Session) Register() string {
	//uuidStr := uuid.NewV4().String()
	// uuidStr1 := uuid.UnmarshalText(uuidStr)
	uuidStr := sess.Save()
	sessStore[uuidStr] = Session{
		Key:          uuidStr,
		UserID:       sess.UserID,
		UserName:     sess.UserName,
		LastAccessed: time.Now(),
	}
	return uuidStr
}

func (sess *Session) Get() bool {
	if val, ok := sessStore[sess.Key]; ok {
		sess.UserID = val.UserID
		sess.UserName = val.UserName

		sessStore[sess.Key] = Session{
			Key:          sess.Key,
			UserID:       sess.UserID,
			UserName:     sess.UserName,
			LastAccessed: time.Now(),
		}
		return true
	}

	return false
}

func (sess *Session) Save() string {
	sess.Key = uuid.NewV4().String()
	session := Session{
		Key:          sess.Key,
		UserID:       sess.UserID,
		UserName:     sess.UserName,
		LastAccessed: time.Now(),
	}
	db.Create(&session)
	return sess.Key
}

//Destroy : method will destroy session
func (sess *Session) Destroy() bool {
	if sess.Key == "" {
		return false
	}
	delete(sessStore, sess.Key)
	return true
}
