package modelsdb

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type PasswordReset struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserID    uuid.UUID `gorm:"type:uuid;index:,unique"`
	Token     string    `gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time `gorm:"type:timestamptz(6);default:now()"`
	ExpiresAt time.Time `gorm:"type:timestamptz(6);not null"`
}

// Role represents the roles table structure.
type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name string    `gorm:"type:varchar(50);not null;unique"`
}

// Session represents the sessions table structure.
type Session struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserID    uuid.UUID  `gorm:"type:uuid;index:,unique"`
	Token     string     `gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time  `gorm:"type:timestamptz(6);default:now()"`
	ExpiresAt time.Time  `gorm:"type:timestamptz(6);not null"`
	DeleteAt  *time.Time `gorm:"type:timestamptz(6)"`
}

// UserInfo represents the user_info table structure.
type UserInfo struct {
	UserID    uuid.UUID      `gorm:"type:uuid;primary_key"`
	Details   datatypes.JSON `gorm:"type:jsonb;not null"`
	CreatedAt time.Time      `gorm:"type:timestamptz(6);default:now()"`
	UpdatedAt time.Time      `gorm:"type:timestamptz(6);default:now()"`
}

// User represents the users table structure.
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Username     string    `gorm:"type:varchar(50);not null;unique"`
	Email        string    `gorm:"type:varchar(100);not null;unique"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"type:timestamptz(6);default:now()"`
	UpdatedAt    time.Time `gorm:"type:timestamptz(6);default:now()"`
	IsActive     bool      `gorm:"type:bool;default:true"`
	IsVerified   bool      `gorm:"type:bool;default:false"`
}
