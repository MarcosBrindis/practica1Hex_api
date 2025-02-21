package postgres

import (
	"context"
	"practica1/src/user/domain/model"
)

func (r *UserRepository) Update(ctx context.Context, id int, user *model.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, id)
	return err
}
