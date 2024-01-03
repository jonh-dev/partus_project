package services

import (
	"context"
	"testing"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/services"
	mocks "github.com/jonh-dev/partus_users/internal/tests/mocks/repositories"
	"github.com/jonh-dev/partus_users/internal/tests/utils"
	"github.com/jonh-dev/partus_users/internal/validation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestPersonalInfoService_CreatePersonalInfo(t *testing.T) {
	mockPersonalInfoRepo := new(mocks.MockPersonalInfoRepository)

	testCases := []struct {
		name          string
		personalInfo  *api.PersonalInfo
		expectedError error
	}{
		{
			name:          "valid",
			personalInfo:  utils.CreateValidPersonalInfo(),
			expectedError: nil,
		},
		{
			name:          "first name with lower case",
			personalInfo:  utils.CreateFirstNameWithLowerCasePersonalInfo(),
			expectedError: validation.ErrInvalidFirstName,
		},
		{
			name:          "first name with more than one word",
			personalInfo:  utils.CreateFirstNameWithMoreThanOneWordPersonalInfo(),
			expectedError: validation.ErrInvalidFirstName,
		},
		{
			name:          "first name with more than 20 characters",
			personalInfo:  utils.CreateFirstNameWithMoreThan20CharactersPersonalInfo(),
			expectedError: validation.ErrInvalidFirstName,
		},
		{
			name:          "last name with more than 50 characters",
			personalInfo:  utils.CreateLastNameWithMoreThan50CharactersPersonalInfo(),
			expectedError: validation.ErrInvalidLastName,
		},
		{
			name:          "last name with lower case start",
			personalInfo:  utils.CreateLastNameWithLowerCaseStartPersonalInfo(),
			expectedError: validation.ErrInvalidLastName,
		},
		{
			name:          "invalid email",
			personalInfo:  utils.CreateInvalidEmailPersonalInfo(),
			expectedError: validation.ErrInvalidUserEmail,
		},
		{
			name:          "existing email",
			personalInfo:  utils.CreateExistingEmailPersonalInfo(),
			expectedError: status.Errorf(codes.AlreadyExists, "E-mail já existe"),
		},
		{
			name:          "future birth date",
			personalInfo:  utils.CreateFutureBirthDatePersonalInfo(),
			expectedError: validation.ErrInvalidBirthDate,
		},
		{
			name:          "birth date before 1900",
			personalInfo:  utils.CreateBirthDateBefore1900PersonalInfo(),
			expectedError: validation.ErrInvalidBirthDate,
		},
		{
			name:          "birth date after current year",
			personalInfo:  utils.CreateBirthDateAfterCurrentYearPersonalInfo(),
			expectedError: validation.ErrInvalidBirthDate,
		},
		{
			name:          "underage",
			personalInfo:  utils.CreateUnderagePersonalInfo(),
			expectedError: validation.ErrInvalidBirthDate,
		},
		{
			name:          "phone with invalid area code",
			personalInfo:  utils.CreatePhoneWithInvalidAreaCodePersonalInfo(),
			expectedError: validation.ErrInvalidPhone,
		},
		{
			name:          "phone with less than nine digits",
			personalInfo:  utils.CreatePhoneWithLessThanNineDigitsPersonalInfo(),
			expectedError: validation.ErrInvalidPhone,
		},
		{
			name:          "phone with more than fourteen digits",
			personalInfo:  utils.CreatePhoneWithMoreThanFourteenDigitsPersonalInfo(),
			expectedError: validation.ErrInvalidPhone,
		},
		{
			name:          "phone with invalid characters",
			personalInfo:  utils.CreatePhoneWithInvalidCharactersPersonalInfo(),
			expectedError: validation.ErrInvalidPhone,
		},

		// Adicione mais cenários de teste conforme necessário...
	}

	s := services.NewPersonalInfoService(mockPersonalInfoRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name != "invalid email" {
				if tc.name == "existing email" {
					mockPersonalInfoRepo.On("DoesEmailExist", mock.Anything, tc.personalInfo.Email).Return(true, nil)
				} else {
					mockPersonalInfoRepo.On("DoesEmailExist", mock.Anything, tc.personalInfo.Email).Return(false, nil)
				}
			}

			mockPersonalInfoRepo.On("CreatePersonalInfo", mock.Anything, mock.AnythingOfType("*api.PersonalInfo")).Return(tc.personalInfo, tc.expectedError)

			personalInfo, err := s.CreatePersonalInfo(context.Background(), tc.personalInfo)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, personalInfo)
			}

			mockPersonalInfoRepo.AssertExpectations(t)
		})
	}

}
