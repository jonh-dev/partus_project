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
		BirthDate:    timestamppb.New(time.Now().AddDate(-20, 0, 0)),
		Phone:        "+5511987654321",
		ProfileImage: "https://example.com/profile.jpg",
	}
}

func CreateFirstNameWithLowerCasePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.FirstName = "john"
	return personalInfo
}

func CreateFirstNameWithMoreThanOneWordPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.FirstName = "John Doe"
	return personalInfo
}

func CreateFirstNameWithMoreThan20CharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.FirstName = "Johndoedoejohndoedoejohndoe"
	return personalInfo
}

func CreateLastNameWithMoreThan50CharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.LastName = "DoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoeDoe"
	return personalInfo
}

func CreateLastNameWithLowerCaseStartPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.LastName = "doe"
	return personalInfo
}

func CreateInvalidEmailPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid email"
	return personalInfo
}

func CreateEmailWithoutAtSymbolPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid.email.com"
	return personalInfo
}

func CreateEmailWithoutDotPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid@com"
	return personalInfo
}

func CreateEmailWithMultipleAtSymbolsPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid@@example.com"
	return personalInfo
}

func CreateEmailWithSpecialCharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid!@example.com"
	return personalInfo
}

func CreateEmailWithSpacesPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid @example.com"
	return personalInfo
}

func CreateEmailWithInvalidTLDPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalid@example.c"
	return personalInfo
}

func CreateEmailWithMoreThan254CharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "invalidddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd@example.com"
	return personalInfo
}

func CreateExistingEmailPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Email = "existing.email@example.com"
	return personalInfo
}

func CreateFutureBirthDatePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Now().AddDate(1, 0, 0))
	return personalInfo
}

func CreateBirthDateBefore1900PersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Date(1899, time.January, 1, 0, 0, 0, 0, time.UTC))
	return personalInfo
}

func CreateBirthDateAfterCurrentYearPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC))
	return personalInfo
}

func CreateUnderagePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.BirthDate = timestamppb.New(time.Now().AddDate(-12, 0, 0))
	return personalInfo
}

func CreatePhoneWithInvalidAreaCodePersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "021234567890"
	return personalInfo
}

func CreatePhoneWithLessThanNineDigitsPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "55119876"
	return personalInfo
}

func CreatePhoneWithMoreThanFourteenDigitsPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "55119876543211234"
	return personalInfo
}

func CreatePhoneWithInvalidCharactersPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.Phone = "55119876abc"
	return personalInfo
}

func CreateInvalidUserIdPersonalInfo() *api.PersonalInfo {
	personalInfo := CreateValidPersonalInfo()
	personalInfo.UserId = "invalid_user_id"
	return personalInfo
}
