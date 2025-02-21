package postgres

import (
	"database/sql"
	"practica1/src/user/domain/ports"
)

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository inicializa el repositorio.
func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &UserRepository{db: db}
}
