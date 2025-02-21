package application

import (
	"context"
	"practica1/src/user/domain/ports"
)

type DeleteUserUsecase struct {
	Repository ports.UserRepository
}

func (uc *DeleteUserUsecase) Execute(ctx context.Context, id int) error {
	return uc.Repository.Delete(ctx, id)
}
