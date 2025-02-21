package postgres

import (
	"context"
	"practica1/src/user/domain/model"
)

func (r *UserRepository) GetByID(ctx context.Context, id int) (*model.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}
