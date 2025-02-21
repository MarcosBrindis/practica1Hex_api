package postgres

import (
	"context"
	"practica1/src/FilmHub/domain/model"
)

func (r *FilmHubRepository) Update(ctx context.Context, id int, film *model.FilmHub) error {
	query := `UPDATE films SET title = $1, type = $2, genre = $3, duration = $4, release_year = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query, film.Title, film.Type, film.Genre, film.Duration, film.ReleaseYear, id)
	return err
}
