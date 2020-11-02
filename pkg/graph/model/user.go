package model

import "time"

// User ...
type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	PasswordHash string    `json:"passwordHash"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	LastLoginAt  time.Time `json:"lastLoginAt"`
}

// UserInput ...
type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
