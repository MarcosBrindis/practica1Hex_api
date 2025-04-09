package ports

type Encrypter interface {
	Encrypt(plainText string) (string, error)
}
