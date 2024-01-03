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
)

func TestAccountInfoService_CreateAccountInfo(t *testing.T) {
	mockAccountInfoRepo := new(mocks.MockAccountInfoRepository)

	testCases := []struct {
		name          string
		accountInfo   *api.AccountInfo
		expectedError error
	}{
		{
			name:          "username válido",
			accountInfo:   utils.CreateValidAccountInfo(),
			expectedError: nil,
		},
		{
			name:          "username curto",
			accountInfo:   utils.CreateShortUsernameAccountInfo(),
			expectedError: validation.ErrInvalidUsername,
		},
		{
			name:          "username longo",
			accountInfo:   utils.CreateLongUsernameAccountInfo(),
			expectedError: validation.ErrInvalidUsername,
		},
		{
			name:          "username com caracteres especiais consecutivos",
			accountInfo:   utils.CreateSpecialCharConsecutiveUsernameAccountInfo(),
			expectedError: validation.ErrInvalidUsername,
		},
		{
			name:          "username começa com caractere especial",
			accountInfo:   utils.CreateSpecialCharStartUsernameAccountInfo(),
			expectedError: validation.ErrInvalidUsername,
		},
		{
			name:          "username termina com caractere especial",
			accountInfo:   utils.CreateSpecialCharEndUsernameAccountInfo(),
			expectedError: validation.ErrInvalidUsername,
		},
		{
			name:          "username com caracteres não permitidos",
			accountInfo:   utils.CreateInvalidCharUsernameAccountInfo(),
			expectedError: validation.ErrInvalidUsername,
		},
		{
			name:          "senha sem letra minúscula",
			accountInfo:   utils.CreateNoLowercasePasswordAccountInfo(),
			expectedError: validation.ErrInvalidPassword,
		},
		{
			name:          "senha sem letra maiúscula",
			accountInfo:   utils.CreateNoUppercasePasswordAccountInfo(),
			expectedError: validation.ErrInvalidPassword,
		},
		{
			name:          "senha sem número",
			accountInfo:   utils.CreateNoNumberPasswordAccountInfo(),
			expectedError: validation.ErrInvalidPassword,
		},
		{
			name:          "senha sem caractere especial",
			accountInfo:   utils.CreateNoSpecialCharPasswordAccountInfo(),
			expectedError: validation.ErrInvalidPassword,
		},
		{
			name:          "senha curta",
			accountInfo:   utils.CreateShortPasswordAccountInfo(),
			expectedError: validation.ErrInvalidPassword,
		},
		{
			name:          "senha longa",
			accountInfo:   utils.CreateLongPasswordAccountInfo(),
			expectedError: validation.ErrInvalidPassword,
		},
		{
			name:          "status da conta inválido",
			accountInfo:   utils.CreateInvalidAccountStatusInfo(),
			expectedError: validation.ErrInvalidAccountStatus,
		},
		{
			name:          "motivo do status inválido",
			accountInfo:   utils.CreateInvalidStatusReasonInfo(),
			expectedError: validation.ErrInvalidStatusReason,
		},
	}

	s := services.NewAccountInfoService(mockAccountInfoRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockAccountInfoRepo.On("CreateAccountInfo", mock.Anything, mock.AnythingOfType("*api.AccountInfo")).Return(tc.accountInfo, tc.expectedError)

			accountInfo, err := s.CreateAccountInfo(context.Background(), tc.accountInfo)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, accountInfo)
			}

			mockAccountInfoRepo.AssertExpectations(t)
		})
	}
}
