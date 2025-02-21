package postgres

import (
	"database/sql"
	"practica1/src/FilmHub/domain/ports"
)

type FilmHubRepository struct {
	db *sql.DB
}

// NewFilmHubRepository inicializa el repositorio.
func NewFilmHubRepository(db *sql.DB) ports.FilmHubRepository {
	return &FilmHubRepository{db: db}
}
