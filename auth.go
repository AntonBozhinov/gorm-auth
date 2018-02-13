package auth

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

// StandardAuthenticator interface uses email and password authentication
type StandardAuthenticator interface {
	Login(email, password string) (userID int64, err error)
	SignUp(email, password string) (userID int64, err error)
	Logout(w http.ResponseWriter, r *http.Request)
}

// SessionManager is responsible for managing user sessions
type SessionManager interface {
	createSession(w http.ResponseWriter, c *http.Cookie) (sessionID int64, err error)
	deleteSession(w http.ResponseWriter) error
}

// Auth struct implements Authenticator interface
type Auth struct {
	loginPath    string
	signUpPath   string
	logoutPath   string
	redirectPath string

	db gorm.DB
}


