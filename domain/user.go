package domain

import (
	"errors"
	"time"
)

// User represents a user in the IAM system.
type User struct {
	ID                string     // UUID in string format
	Email             *string    // Email address for recovery and communication
	PasswordHash      *string    // Hashed password
	FullName          *string    // Full name of the user (optional)
	PhoneNumber       *string    // Optional phone number for contact/2FA
	TwoFactorEnabled  bool       // Indicates if 2FA is enabled
	LastLogin         *time.Time // Tracks the last login time (optional)
	OrganizationID    *string    // Links the user to an organization, if applicable (optional)
	Status            string     // User status (active, suspended, deactivated)
	LastPasswordReset *time.Time // Tracks the last password reset time (optional)
	CreatedAt         time.Time  // Timestamp when the user was created
	UpdatedAt         time.Time  // Timestamp when the user was last updated
	DeletedAt         *time.Time // Optional soft delete time (null if not deleted)
}

// Validate checks the integrity of the user data.
func (u *User) Validate() error {
	if u.Status == "" {
		return errors.New("status cannot be empty")
	}
	if u.Status != "active" && u.Status != "suspended" && u.Status != "deactivated" {
		return errors.New("invalid status")
	}
	// No validation for optional fields like FullName, PhoneNumber, etc.
	return nil
}
