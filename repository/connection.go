package repository

import (
	"context"
	"fmt"
	"github.com/devdowns/ingredients-hexagonal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnection(dbConfig *config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?connect_timeout=%d", dbConfig.User, dbConfig.Password, dbConfig.URL, dbConfig.DB, dbConfig.Timeout)
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return pool, nil
}
