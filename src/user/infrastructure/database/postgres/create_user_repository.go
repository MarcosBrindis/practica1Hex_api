package postgres

import (
	"context"
	"practica1/src/user/domain/model"
)

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID)
}
