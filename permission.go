package auth

import (
	"github.com/jinzhu/gorm"
)

// Permissions slice of all role permissions
type Permissions []*Permission

// Permission is individual permission a role can has
type Permission struct {
	gorm.Model
	Name string
	Description string
	CategoryID uint
	Category PermissionCategory
}

// PermissionCategory is used to categorize permissions
type PermissionCategory struct {
	gorm.Model
	Name string
}