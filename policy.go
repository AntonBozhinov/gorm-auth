package auth

import (
	"time"

	"github.com/jinzhu/gorm"
)

// AccessPolicy type
type AccessPolicy struct {
	gorm.Model
	Name            string
	Description     string
	expiresOn       time.Time
	SessionDuration time.Time
	Credentials     []PolicyCredential
	isEnabled       bool
}

// AccessTimeFrame type
type AccessTimeFrame struct {
	gorm.Model
	ActivePeriod   string
	ActivateAt     time.Time
	ActiveWeekDays int
}

// PolicyCredential type
type PolicyCredential struct {
	gorm.Model
	Name                string
	Description         string
	Regex               string
	MinChar             int
	MaxChar             int
	ExpiresOn           time.Time
	NotificationChannel string
	AllowChange         bool
	NotifyOnChange      bool
	NotifyOnCreate      bool
	PolicyID            uint
	Policy              AccessPolicy
	AllowedAttempts     uint
}


