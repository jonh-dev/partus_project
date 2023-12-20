package validation

import (
	"errors"
	"regexp"

	"github.com/jonh-dev/partus_users/api"
)

var (
	ErrInvalidUsername = errors.New("o nome de usuário deve começar e terminar com um caractere alfanumérico, pode conter letras, números, pontos, hifens e sublinhados, não pode conter caracteres especiais consecutivos, e deve ter entre 3 e 20 caracteres")
)

func ValidateAccountInfo(accountInfo *api.AccountInfo) error {
	if !isValidUsername(accountInfo.Username) {
		return ErrInvalidUsername
	}

	// Adicione aqui as outras validações do AccountInfo

	return nil
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^(?i)[a-z0-9]+([._-]?[a-z0-9]+)*$`)
	return len(username) >= 3 && len(username) <= 20 && re.MatchString(username)
}
