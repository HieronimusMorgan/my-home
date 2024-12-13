package master

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID         uint      `gorm:"primaryKey" json:"user_id,omitempty"`
	UUIDKey        string    `gorm:"unique" json:"uuid_key,omitempty"`
	ClientID       string    `gorm:"unique" json:"client_id,omitempty"`
	Username       string    `gorm:"unique" json:"username,omitempty"`
	Password       string    `json:"-"`
	FirstName      string    `json:"first_name,omitempty"`
	LastName       string    `json:"last_name,omitempty"`
	FullName       string    `json:"full_name,omitempty"`
	PhoneNumber    string    `gorm:"unique" json:"phone_number,omitempty"`
	ProfilePicture string    `json:"profile_picture,omitempty"`
	RoleID         uint      `gorm:"not null" json:"role_id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	CreatedBy      string    `json:"created_by,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	UpdatedBy      string    `json:"updated_by,omitempty"`
	DeletedAt      bool      `json:"deleted_at,omitempty"`
	DeletedBy      string    `json:"deleted_by,omitempty"`
}

var UserDB *gorm.DB
