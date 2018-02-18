package auth

import (
	"github.com/jinzhu/gorm"
)

// RoleCategory is used to categorize roles to given entity. For example users organizations groups and so on...
type RoleCategory struct {
	gorm.Model
	Name string
	Description string
}

// Role is the user role holdining a set of permissions used for user authorization
type Role struct {
	gorm.Model
	UserID uint
	Name string
	Description string
	Category RoleCategory
	CategoryID uint
	Permissions []Permission
}