package models

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type User struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserName  string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName string     `gorm:"type:varchar(100);not null"`
	LastName  string     `gorm:"type:varchar(100);not null"`
	Email     string     `gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone     string     `gorm:"type:varchar(20);uniqueIndex;not null"`
	Password  string     `gorm:"type:varchar(100);not null"`
	Verified  *bool      `gorm:"not null;default:false"`
	Role      *string    `gorm:"type:varchar(50);default:'user';not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	// One-to-one relationship with Account
	Account *Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type SignUpInput struct {
	UserName        string `json:"user_name" validate:"required"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Phone           string `json:"phone" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type SignInByPhone struct {
	Phone    string `json:"phone"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	AccountId uuid.UUID `json:"account_id,omitempty"`
	ID        uuid.UUID `json:"id,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdate struct {
	UserName  string `json:"user_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
	Verified  bool   `json:"verified,omitempty"`
}

type UserPasswordUpdate struct {
	OldPassword        string `json:"old_password,omitempty"`
	NewPassword        string `json:"new_password,omitempty"`
	ConfirmNewPassword string `json:"confirm_new_password,omitempty"`
}

type SignInResponse struct {
	Token       string       `json:"token"`
	UserDetails UserResponse `json:"user"`
}

func FilterUserRecord(user *User, accountId *uuid.UUID) UserResponse {
	var accountID uuid.UUID

	// Check if accountId is nil, handle accordingly
	if accountId != nil {
		accountID = *accountId
	}

	return UserResponse{
		AccountId: accountID, // AccountId is nil if no account
		ID:        *user.ID,
		UserName:  user.UserName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      *user.Role,
		CreatedAt: *user.CreatedAt, // Ensure non-nil value
		UpdatedAt: *user.UpdatedAt, // Ensure non-nil value
	}
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
