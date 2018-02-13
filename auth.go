package auth

import (
	"github.com/jinzhu/gorm"
)

// Auth struct implements Authenticator interface
type Auth struct {
	DB *gorm.DB
}

// New is used to initialize the auth struct giving a db to write to
func New(db *gorm.DB) *Auth {
	a := &Auth{DB: db}
	a.DB.AutoMigrate(&Session{})
	return a
}
