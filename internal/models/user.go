package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
	UserStatusBanned   = "banned"

	MinPasswordLength = 8
	MaxPasswordLength = 72
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"password,omitempty" gorm:"not null"`
	Role      Role           `json:"role" gorm:"type:varchar(20);not null;default:'user'"`
	Status    string         `json:"status" gorm:"type:varchar(20);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		// Name validation
		validation.Field(&u.Name,
			validation.Required.Error("name is required"),
			validation.Length(2, 50).Error("name must be between 2 and 50 characters"),
		),

		// Email validation
		validation.Field(&u.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),

		// Password validation
		validation.Field(&u.Password,
			validation.Required.Error("password is required"),
			validation.Length(MinPasswordLength, MaxPasswordLength).
				Error("password must be between 8 and 72 characters"),
		),

		// Role validation
		validation.Field(&u.Role,
			validation.Required.Error("role is required"),
			validation.In(RoleAdmin, RoleUser).Error("invalid role"),
		),

		// Status validation
		validation.Field(&u.Status,
			validation.Required.Error("status is required"),
			validation.In(UserStatusActive, UserStatusInactive, UserStatusBanned).
				Error("invalid status"),
		),
	)
}

func (u *User) ValidateUpdate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name,
			validation.Required.Error("name is required"),
			validation.Length(2, 50).Error("name must be between 2 and 50 characters"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),
		// Password is optional during update
		validation.Field(&u.Password,
			validation.When(len(u.Password) > 0, validation.Length(MinPasswordLength, MaxPasswordLength).
				Error("password must be between 8 and 72 characters")),
		),
		validation.Field(&u.Role,
			validation.Required.Error("role is required"),
			validation.In(RoleAdmin, RoleUser).Error("invalid role"),
		),
		validation.Field(&u.Status,
			validation.Required.Error("status is required"),
			validation.In(UserStatusActive, UserStatusInactive, UserStatusBanned).
				Error("invalid status"),
		),
	)
}

func (u *User) HashPassword() error {
	if len(u.Password) == 0 {
		return validation.NewError("validation_error", "password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) beforeCreate(tx *gorm.DB) error {
	// Set timestamps
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now

	// Set defaults if empty
	if u.Role == "" {
		u.Role = RoleUser
	}
	if u.Status == "" {
		u.Status = UserStatusActive
	}

	return nil
}

func (u *User) beforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
