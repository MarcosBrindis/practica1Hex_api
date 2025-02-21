package application

import (
	"context"
	"practica1/src/user/domain/model"
	"practica1/src/user/domain/ports"
)

type GetAllUsersUsecase struct {
	Repository ports.UserRepository
}

func (uc *GetAllUsersUsecase) Execute(ctx context.Context) ([]*model.User, error) {
	return uc.Repository.GetAll(ctx)
}
