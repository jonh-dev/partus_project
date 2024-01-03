package utils

import (
	"time"

	"github.com/jonh-dev/partus_users/api"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateValidPersonalInfo() *api.PersonalInfo {
	return &api.PersonalInfo{
		UserId:       "valid_user_id",
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john.doe@example.com",
		BirthDate:    timestamppb.New(time.Now().AddDate(-20, 0, 0)), // Data de nascimento 20 anos atrás
		Phone:        "+5511987654321",                               // Número de telefone brasileiro válido
		ProfileImage: "https://example.com/profile.jpg",              // URL de imagem de perfil válido
	}
}

func CreateFirstNameWithLowerCasePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.FirstName = "john" // Primeiro nome começa com letra minúscula
	return personalInfo
}

func CreateFirstNameWithMoreThanOneWordPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.FirstName = "John Doe" // Primeiro nome contém mais de uma palavra
	return personalInfo
}

func CreateFirstNameWithMoreThan20CharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.FirstName = "Johndoedoejohndoedoejohndoe" // Primeiro nome contém mais de 20 caracteres
	return personalInfo
}

func CreateLastNameWithMoreThan50CharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.LastName = "DoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoe" // Sobrenome com mais de 50 caracteres
	return personalInfo
}

func CreateLastNameWithLowerCaseStartPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.LastName = "doe" // Sobrenome começa com letra minúscula
	return personalInfo
}

func CreateInvalidEmailPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid email" // Endereço de e-mail inválido
	return personalInfo
}

func CreateExistingEmailPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "existing.email@example.com" // E-mail que já existe na base de dados
	return personalInfo
}

func CreateFutureBirthDatePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Now().AddDate(1, 0, 0)) // Data de nascimento no futuro
	return personalInfo
}

func CreateBirthDateBefore1900PersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Date(1899, time.January, 1, 0, 0, 0, 0, time.UTC)) // Data de nascimento antes de 1900
	return personalInfo
}

func CreateBirthDateAfterCurrentYearPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)) // Data de nascimento após o ano atual
	return personalInfo
}

func CreateUnderagePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Now().AddDate(-12, 0, 0)) // Usuário com menos de 13 anos
	return personalInfo
}

func CreatePhoneWithInvalidAreaCodePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "021234567890" // Número de telefone com código de área inválido
	return personalInfo
}

func CreatePhoneWithLessThanNineDigitsPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "55119876" // Número de telefone com menos de nove dígitos
	return personalInfo
}

func CreatePhoneWithMoreThanFourteenDigitsPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "55119876543211234" // Número de telefone com mais de quatorze dígitos
	return personalInfo
}

func CreatePhoneWithInvalidCharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "55119876abc" // Número de telefone com caracteres inválidos
	return personalInfo
}

func CreateInvalidUserIdPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.UserId = "invalid_user_id" // UserId inválido
	return personalInfo
}
