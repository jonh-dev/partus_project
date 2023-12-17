package validation

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jonh-dev/partus_users/api"
)

var (
	ErrInvalidFirstName = errors.New("o primeiro nome deve começar com uma letra maiúscula, conter apenas uma palavra e ter no máximo 20 caracteres")
	ErrInvalidLastName  = errors.New("o sobrenome deve começar com uma letra maiúscula em cada palavra e ter no máximo 50 caracteres")
	ErrInvalidUserEmail = errors.New("o e-mail deve ser um endereço de e-mail válido")
	ErrInvalidBirthDate = errors.New("a data de nascimento deve estar no passado, o usuário deve ter pelo menos 13 anos e o ano deve ser entre 1900 e o ano atual")
	ErrInvalidPhone     = errors.New("o telefone ou celular deve estar no formato correto, ou seja, começar com '+' seguido de 1 a 3 dígitos para números internacionais, ou começar diretamente com um dígito para números brasileiros, e ter entre 9 e 14 dígitos no total, sem conter nenhum caractere que não seja dígito ou '+'")
)

func ValidateUser(user *api.User) error {
	if !isValidFirstName(user.PersonalInfo.FirstName) {
		return ErrInvalidFirstName
	}

	if !isValidLastName(user.PersonalInfo.LastName) {
		return ErrInvalidLastName
	}

	if !isValidEmail(user.PersonalInfo.Email) {
		return ErrInvalidUserEmail
	}

	if !isValidBirthDate(user.PersonalInfo.BirthDate) {
		return ErrInvalidBirthDate
	}

	if err := isValidPhone(user.PersonalInfo.Phone); err != nil {
		return err
	}

	return nil
}

func ValidateCreateUserRequest(req *api.CreateUserRequest) error {
	user := &api.User{
		PersonalInfo: req.PersonalInfo,
	}
	return ValidateUser(user)
}

func isValidFirstName(name string) bool {
	re := regexp.MustCompile(`^[A-ZÁÉÍÓÚÂÊÎÔÛÃÕ][a-záéíóúâêîôûãõA-ZÁÉÍÓÚÂÊÎÔÛÃÕ]{0,19}$`)
	return re.MatchString(name)
}

func isValidLastName(name string) bool {
	if len(name) > 50 {
		return false
	}
	re := regexp.MustCompile(`^([A-ZÁÉÍÓÚÂÊÎÔÛÃÕ][a-záéíóúâêîôûãõA-ZÁÉÍÓÚÂÊÎÔÛÃÕ]*\s*)+$`)
	return re.MatchString(name)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidBirthDate(date *timestamp.Timestamp) bool {
	t := date.AsTime()

	if time.Now().Before(t) {
		return false
	}

	year := t.Year()
	if year < 1900 || year > time.Now().Year() {
		return false
	}

	years := time.Now().Year() - t.Year()
	if t.After(time.Now().AddDate(-years, 0, 0)) {
		years--
	}

	if years < 13 {
		return false
	}

	return true
}

func isValidPhone(phone string) error {
	re := regexp.MustCompile(`^(\+\d{2})?(\d{2}|\(\d{2}\))\s?\d{4,5}-?\d{4}$`)

	if !re.MatchString(phone) {
		return ErrInvalidPhone
	}

	// Verifica se o número tem o código de área
	areaCode := re.FindStringSubmatch(phone)[2]
	if areaCode == "" {
		return ErrInvalidPhone
	}

	return nil
}
