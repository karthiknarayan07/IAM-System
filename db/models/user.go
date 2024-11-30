package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserStatus defines the allowed statuses for the User model.
type UserStatus string

const (
	StatusActive      UserStatus = "active"
	StatusSuspended   UserStatus = "suspended"
	StatusDeactivated UserStatus = "deactivated"
)

// User represents a user in the IAM system.
type User struct {
	ID                uuid.UUID      `gorm:"type:uuid;primaryKey"`                       // UUID primary key
	Email             *string        `gorm:"size:255"`                                   // Email address for recovery and communication
	PasswordHash      *string        `gorm:"size:255"`                                   // Hashed password
	FullName          *string        `gorm:"size:255"`                                   // Full name of the user
	PhoneNumber       *string        `gorm:"size:20"`                                    // Optional phone number for contact/2FA
	TwoFactorEnabled  bool           `gorm:"default:false"`                              // Indicates if 2FA is enabled
	LastLogin         *time.Time     `gorm:"index"`                                      // Tracks the last login time
	OrganizationID    *uuid.UUID     `gorm:"index"`                                      // Links the user to an organization, if applicable
	Status            UserStatus     `gorm:"type:varchar(50);not null;default:'active'"` // Enum-like user status
	LastPasswordReset *time.Time     `gorm:"index"`                                      // Tracks the last password reset time
	CreatedAt         time.Time      `gorm:"autoCreateTime"`                             // Timestamp when the user was created
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`                             // Timestamp when the user was last updated
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

// Implement the Scanner and Valuer interfaces for UserStatus

// Value converts the UserStatus to a driver.Value (for database storage).
func (s UserStatus) Value() (driver.Value, error) {
	switch s {
	case StatusActive, StatusSuspended, StatusDeactivated:
		return string(s), nil
	default:
		return nil, errors.New("invalid UserStatus value")
	}
}

// Scan converts a value from the database into a UserStatus.
func (s *UserStatus) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid type for UserStatus")
	}

	switch UserStatus(str) {
	case StatusActive, StatusSuspended, StatusDeactivated:
		*s = UserStatus(str)
		return nil
	default:
		return errors.New("invalid UserStatus value")
	}
}
