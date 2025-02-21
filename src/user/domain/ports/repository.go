package ports

import (
	"context"
	"practica1/src/user/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, id int, user *model.User) error
	GetByID(ctx context.Context, id int) (*model.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.User, error)
}
