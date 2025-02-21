package ports

import (
	"context"
	"practica1/src/FilmHub/domain/model"
)

type FilmHubRepository interface {
	Create(ctx context.Context, filmhub *model.FilmHub) error
	Update(ctx context.Context, id int, filmhub *model.FilmHub) error
	GetByID(ctx context.Context, id int) (*model.FilmHub, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.FilmHub, error)
}
