package model

import "github.com/google/uuid"

type User struct {
	UserID   uuid.UUID `gorm:"primaryKey"`
	Name     string    `gorm:"not null; type:varchar(100)"`
	Email    string    `gorm:"not null; type:varchar(100); unique"`
	Password string    `gorm:"not null; type:varchar(100)"`
	Role     bool      `gorm:"not null"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserRegisterResponse struct {
	UserID uuid.UUID `json:"userID"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
