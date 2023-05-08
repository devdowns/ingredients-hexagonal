package repository

import (
	"context"
	"errors"
	"github.com/devdowns/ingredients-hexagonal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ingredientRepo struct {
	dbPool *pgxpool.Pool
}

// NewIngredientService creates a new instance of the ingredient service.
func NewIngredientRepo(dbPool *pgxpool.Pool) *ingredientRepo {
	return &ingredientRepo{
		dbPool: dbPool,
	}
}

func (s *ingredientRepo) Find(id uuid.UUID) (*domain.Ingredient, error) {
	query := "SELECT id, name, description FROM ingredients WHERE id = $1"
	row := s.dbPool.QueryRow(context.Background(), query, id)

	ingredient := &domain.Ingredient{}
	err := row.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return ingredient, nil
}

func (s *ingredientRepo) Store(ingredient *domain.Ingredient) error {
	query := "INSERT INTO ingredients (name, description) VALUES ($1, $2) RETURNING id"
	row := s.dbPool.QueryRow(context.Background(), query, ingredient.Name, ingredient.Description)
	err := row.Scan(&ingredient.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ingredientRepo) Update(ingredient *domain.Ingredient) error {
	query := "UPDATE ingredients SET name = $1, description = $2 WHERE id = $3"
	_, err := s.dbPool.Exec(context.Background(), query, ingredient.Name, ingredient.Description, ingredient.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ingredientRepo) FindAll() ([]*domain.Ingredient, error) {
	query := "SELECT id, name, description FROM ingredients"
	rows, err := s.dbPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []*domain.Ingredient

	for rows.Next() {
		ingredient := &domain.Ingredient{}
		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Description)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

func (s *ingredientRepo) Delete(id uuid.UUID) error {
	query := "DELETE FROM ingredients WHERE id = $1"
	result, err := s.dbPool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}

	return nil
}
