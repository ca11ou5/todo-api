package repository

import (
	"database/sql"
	"todo-list/configs"
	"todo-list/internal/repository/postgres"
)

type Repository struct {
	*sql.DB
}

func NewRepository(cfg *configs.Config) *Repository {
	return &Repository{postgres.ConnectToPostgres(cfg.PostgresURL)}
}
