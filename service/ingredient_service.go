package service

import (
	"github.com/devdowns/ingredients-hexagonal/domain"
	"github.com/google/uuid"
)

type service struct {
	ingredientRepo domain.IngredientRepository
}

// NewIngredientService a service struct that attaches to a repository via the Repository interface
func NewIngredientService(ingredientRepo domain.IngredientRepository) *service {
	return &service{ingredientRepo: ingredientRepo}
}

func (s *service) Find(id uuid.UUID) (*domain.Ingredient, error) {
	return s.ingredientRepo.Find(id)
}
func (s *service) Store(ingredient *domain.Ingredient) error {
	return s.ingredientRepo.Store(ingredient)
}
func (s *service) Update(ingredient *domain.Ingredient) error {
	return s.ingredientRepo.Update(ingredient)
}

func (s *service) FindAll() ([]*domain.Ingredient, error) {
	return s.ingredientRepo.FindAll()
}

func (s *service) Delete(id uuid.UUID) error {
	return s.ingredientRepo.Delete(id)
}
