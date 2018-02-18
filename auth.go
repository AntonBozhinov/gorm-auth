package auth

import (
	"github.com/jinzhu/gorm"
)

// Auth struct implements Authenticator interface
type Auth struct {
	DB *gorm.DB
}

// New is used to initialize the auth struct with a given db to write to
func New(db *gorm.DB) *Auth {
	a := &Auth{DB: db}
	a.DB.AutoMigrate(&Session{})
	a.DB.AutoMigrate(&User{})
	return a
}