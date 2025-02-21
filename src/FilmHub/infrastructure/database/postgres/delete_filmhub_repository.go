package postgres

import "context"

func (r *FilmHubRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM films WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
