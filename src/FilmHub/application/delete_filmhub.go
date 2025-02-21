package application

import (
	"context"
	"practica1/src/FilmHub/domain/ports"
)

type DeleteFilmHubUsecase struct {
	Repository ports.FilmHubRepository
}

func (uc *DeleteFilmHubUsecase) Execute(ctx context.Context, id int) error {
	return uc.Repository.Delete(ctx, id)
}
