package encryption

import "golang.org/x/crypto/bcrypt"

type PasswordEncryptor interface {
	EncryptPassword(password string) (string, error)
}

type BcryptPasswordEncryptor struct{}

func (b *BcryptPasswordEncryptor) EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
