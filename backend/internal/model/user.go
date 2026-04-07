package model

import "time"

type User struct {
	ID             int64     `json:"id" db:"id"`
	Email          string    `json:"email" db:"email"`
	PasswordHash   string    `json:"-" db:"password_hash"`
	Name           string    `json:"name" db:"name"`
	Role           string    `json:"role" db:"role"`
	MaxRestaurants int       `json:"max_restaurants" db:"max_restaurants"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// Admin creates/updates users
type AdminCreateUserRequest struct {
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	Name           string `json:"name" binding:"required"`
	Role           string `json:"role"`
	MaxRestaurants *int   `json:"max_restaurants"`
}

type AdminUpdateUserRequest struct {
	Email          *string `json:"email"`
	Password       *string `json:"password"`
	Name           *string `json:"name"`
	Role           *string `json:"role"`
	MaxRestaurants *int    `json:"max_restaurants"`
}
