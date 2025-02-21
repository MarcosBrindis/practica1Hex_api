package postgres

import (
	"context"
	"practica1/src/FilmHub/domain/model"
)

func (r *FilmHubRepository) GetByID(ctx context.Context, id int) (*model.FilmHub, error) {
	query := `SELECT id, title, type, genre, duration, release_year FROM films WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	filmhub := &model.FilmHub{}
	err := row.Scan(&filmhub.ID, &filmhub.Title, &filmhub.Type, &filmhub.Genre, &filmhub.Duration, &filmhub.ReleaseYear)
	return filmhub, err
}
