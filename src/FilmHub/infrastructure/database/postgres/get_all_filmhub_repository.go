package postgres

import (
	"context"
	"practica1/src/FilmHub/domain/model"
)

func (r *FilmHubRepository) GetAll(ctx context.Context) ([]*model.FilmHub, error) {
	query := `SELECT id, title, type, genre, duration, release_year FROM films`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var films []*model.FilmHub
	for rows.Next() {
		filmhub := &model.FilmHub{}
		if err := rows.Scan(&filmhub.ID, &filmhub.Title, &filmhub.Type, &filmhub.Genre, &filmhub.Duration, &filmhub.ReleaseYear); err != nil {
			return nil, err
		}
		films = append(films, filmhub)
	}
	return films, nil
}
