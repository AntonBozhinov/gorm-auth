package auth

import (
	"github.com/jinzhu/gorm"
)

// Permission is individual permission a role can have
type Permission struct {
	gorm.Model
	RoleID      uint
	Action      string
	Resource    string
	CategoryID  uint
	Category    PermissionCategory
	Protections []Protection
}

// PermissionCategory is used to categorize permissions
type PermissionCategory struct {
	gorm.Model
	Name string
}

// Protection defines specific request values to be required when requesting for permission
// a combination of permission action and permission resource can be additional protected with some extra
// parameter that is required to be present in the request
type Protection struct {
	gorm.Model
	PermissionID uint
	Name         string
	Description  string
	HeaderKey    string
	BodyKey      string
	Credential   string
}
