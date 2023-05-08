package domain

import "github.com/google/uuid"

// IngredientService , interface defining all CRUD operations
type IngredientService interface {
	Find(id uuid.UUID) (*Ingredient, error)
	Store(ingredient *Ingredient) error
	Update(ingredient *Ingredient) error
	FindAll() ([]*Ingredient, error)
	Delete(id uuid.UUID) error
}

// IngredientRepository , interface acting like a port for the database implementation
type IngredientRepository interface {
	Find(id uuid.UUID) (*Ingredient, error)
	Store(ingredient *Ingredient) error
	Update(ingredient *Ingredient) error
	FindAll() ([]*Ingredient, error)
	Delete(id uuid.UUID) error
}
