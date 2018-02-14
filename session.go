package auth

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Session persist user session
type Session struct {
	gorm.Model
	Language  string
	UserID 	  uint
	User   	  User
	RemoteIP  string
	UserAgent string
	ExpireAt  time.Time
	Channel   string
}

// SessionManager is responsible for managing user sessions
type SessionManager interface {
	GetSession(ID uint) (*Session, error)
	UpdateSession(s *Session) error
	CreateSession(name string) (*Session, error)
	DeleteSession(ID uint) error
}
