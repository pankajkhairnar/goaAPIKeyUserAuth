package main

import (
	"time"

	"github.com/twinj/uuid"
)

type Session struct {
	Key          string
	UserID       uint
	UserName     string
	CreatedAt    time.Time
	LastAccessed time.Time
}

var sessStore = make(map[string]Session)

func LoadSession() {
	session := []Session{}
	if db.Find(&session).RecordNotFound() != true {
		currentTime := time.Now()
		time := 3 * time.Hour
		for _, row := range session {
			diff := currentTime.Sub(row.CreatedAt)
			if diff >= time {
				db.Where("id = ?", row.UserID).Delete(&Session{})
			} else {
				sessStore[row.Key] = Session{
					Key:          row.Key,
					UserID:       row.UserID,
					UserName:     row.UserName,
					CreatedAt:    row.CreatedAt,
					LastAccessed: row.LastAccessed,
				}
			}

		}
	}
}

// Register : method will create new session id and store session values in map, it will return session tokey
func (sess *Session) Register() string {
	uuidStr := sess.Save()
	sessStore[uuidStr] = Session{
		Key:          uuidStr,
		UserID:       sess.UserID,
		UserName:     sess.UserName,
		CreatedAt:    time.Now(),
		LastAccessed: time.Now(),
	}
	return uuidStr
}

func (sess *Session) Get() bool {
	var session Session
	db.Model(&session).Where("id = ?", sess.UserID).Update("LastAccessed", time.Now())
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
		CreatedAt:    time.Now(),
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
	id := sessStore[sess.Key].UserID
	db.Where("user_id = ?", id).Delete(&Session{})
	delete(sessStore, sess.Key)
	return true
}
