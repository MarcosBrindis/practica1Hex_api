package application

import (
	"context"
	"practica1/src/FilmHub/domain/model"
	"practica1/src/FilmHub/domain/ports"
)

type CreateFilmHubUsecase struct {
	Repository ports.FilmHubRepository
}

func (uc *CreateFilmHubUsecase) Execute(ctx context.Context, filmhub *model.FilmHub) error {
	return uc.Repository.Create(ctx, filmhub)
}
