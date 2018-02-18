package auth

import (
	"github.com/jinzhu/gorm"
)

// User struct represents system user
type User struct {
	gorm.Model
	Email string
	Username string
	Password string
	Roles []Role
}