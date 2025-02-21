package postgres

import (
	"context"
	"practica1/src/FilmHub/domain/model"
)

func (r *FilmHubRepository) Create(ctx context.Context, filmhub *model.FilmHub) error {
	query := `INSERT INTO films (title, type, genre, duration, release_year) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return r.db.QueryRowContext(ctx, query, filmhub.Title, filmhub.Type, filmhub.Genre, filmhub.Duration, filmhub.ReleaseYear).Scan(&filmhub.ID)
}
