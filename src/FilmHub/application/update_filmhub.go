package application

import (
	"context"
	"practica1/src/FilmHub/domain/model"
	"practica1/src/FilmHub/domain/ports"
)

type UpdateFilmHubUsecase struct {
	Repository ports.FilmHubRepository
}

func (uc *UpdateFilmHubUsecase) Execute(ctx context.Context, id int, filmhub *model.FilmHub) error {
	return uc.Repository.Update(ctx, id, filmhub)
}
