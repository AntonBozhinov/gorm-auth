package auth

import (
	"github.com/jinzhu/gorm"
)

// Auth struct implements Authenticator interface
type Auth struct {
	version string
	DB *gorm.DB
}

// New is used to initialize the auth struct with a given db to write to
func New(db *gorm.DB) *Auth {
	a := &Auth{DB: db}
	a.DB.AutoMigrate(&Session{})
	a.DB.AutoMigrate(&User{})
	a.DB.AutoMigrate(&AccessPolicy{})
	a.DB.AutoMigrate(&PolicyCredential{})
	a.DB.AutoMigrate(&Permission{})
	a.DB.AutoMigrate(&PermissionCategory{})
	a.DB.AutoMigrate(&Protection{})
	a.DB.AutoMigrate(&Role{})
	a.DB.AutoMigrate(&RoleCategory{})
	return a
}