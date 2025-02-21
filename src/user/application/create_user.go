package application

import (
	"context"
	"practica1/src/user/domain/model"
	"practica1/src/user/domain/ports"
)

type CreateUserUsecase struct {
	Repository ports.UserRepository
}

func (uc *CreateUserUsecase) Execute(ctx context.Context, user *model.User) error {
	return uc.Repository.Create(ctx, user)
}
