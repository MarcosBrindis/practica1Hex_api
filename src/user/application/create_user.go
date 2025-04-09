package application

import (
	"context"
	"practica1/src/user/domain/model"
	"practica1/src/user/domain/ports"
)

type CreateUserUsecase struct {
	Repository ports.UserRepository
	Encrypter  ports.Encrypter
}

func (uc *CreateUserUsecase) Execute(ctx context.Context, user *model.User) error {
	encryptedEmail, err := uc.Encrypter.Encrypt(user.Email)
	if err != nil {
		return err
	}
	user.Email = encryptedEmail
	return uc.Repository.Create(ctx, user)
}
