package service

import (
	"golang.org/x/crypto/bcrypt"

	"practica1/src/user/domain/ports"
)

type BcryptEncrypter struct {
	cost int
}

func NewBcryptEncrypter(cost int) ports.Encrypter {
	return &BcryptEncrypter{cost: cost}
}

func (b *BcryptEncrypter) Encrypt(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), b.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
