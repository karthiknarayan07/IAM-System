package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`     // UUID primary key
	Username  string    `gorm:"size:100;not null;unique"` // Unique username
	Email     string    `gorm:"size:100;not null;unique"` // Unique email
	CreatedAt time.Time `gorm:"autoCreateTime"`           // Auto timestamp
	UpdatedAt time.Time `gorm:"autoUpdateTime"`           // Auto timestamp
}
