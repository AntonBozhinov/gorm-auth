package auth

import (
	"time"
)

// Session persist user session
type Session struct {
	ID        uint `gorm:"index"`
	Language  string
	UserID    int
	RemoteIP  string
	UserAgent string
	ExpireAt  time.Time
	CreatedAt time.Time
	Channel   string
}

// SessionManager is responsible for managing user sessions
type SessionManager interface {
	GetSession(ID uint) (*Session, error)
	UpdateSession(s *Session) error
	CreateSession(name string) (*Session, error)
	DeleteSession(ID uint) error
}
