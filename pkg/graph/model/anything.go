package model

import "time"

// Anything ...
type Anything struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// AnythingInput ...
type AnythingInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
