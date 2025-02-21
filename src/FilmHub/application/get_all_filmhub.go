package application

import (
	"context"
	"practica1/src/FilmHub/domain/model"
	"practica1/src/FilmHub/domain/ports"
)

type GetAllFilmHubUsecase struct {
	Repository ports.FilmHubRepository
}

func (uc *GetAllFilmHubUsecase) Execute(ctx context.Context) ([]*model.FilmHub, error) {
	return uc.Repository.GetAll(ctx)
}
