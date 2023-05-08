package domain

import "github.com/google/uuid"

// Ingredient struct to store ingredient data
type Ingredient struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
