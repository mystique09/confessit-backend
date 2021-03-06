package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Username  string    `gorm:"unique"`
	Password  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}

type JwtUserPayload struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

func (u *User) HashPassword() error {
	hash_password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hash_password)
	return nil
}

func (u *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

type UserCreatePayload struct {
	Username string `json:"username" validate:"required,min=8,max=18"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginPayload struct {
	Username string `json:"username" validate:"required,min=8,max=18"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdatePayload struct {
	Username string `json:"username" validate:"required,min=8,max=18"`
	Password string `json:"password" validate:"required,min=12"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
}
