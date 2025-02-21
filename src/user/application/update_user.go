package application

import (
	"context"
	"practica1/src/user/domain/model"
	"practica1/src/user/domain/ports"
)

type UpdateUserUsecase struct {
	Repository ports.UserRepository
}

func (uc *UpdateUserUsecase) Execute(ctx context.Context, id int, user *model.User) error {
	return uc.Repository.Update(ctx, id, user)
}
