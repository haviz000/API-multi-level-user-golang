package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID   uuid.UUID `gorm:"primaryKey"`
	Title       string    `gorm:"not null; type:varchar(50)"`
	Description string    `gorm:"not null; type:varchar(255)"`
	UserID      uuid.UUID
	User        *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductCreateResponse struct {
	ProductID   uuid.UUID `json:"product_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id"`
}

type ProductResponse struct {
	ProductID   string `json:"product_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
