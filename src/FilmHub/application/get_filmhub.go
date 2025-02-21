package application

import (
	"context"
	"practica1/src/FilmHub/domain/model"
	"practica1/src/FilmHub/domain/ports"
)

type GetFilmHubUsecase struct {
	Repository ports.FilmHubRepository
}

func (uc *GetFilmHubUsecase) Execute(ctx context.Context, id int) (*model.FilmHub, error) {
	return uc.Repository.GetByID(ctx, id)
}
