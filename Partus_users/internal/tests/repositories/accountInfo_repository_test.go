package repositories

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/jonh-dev/partus_users/internal/tests/mocks/repositories"
	"github.com/jonh-dev/partus_users/internal/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestAccountInfoRepository_CreateAccountInfo(t *testing.T) {
	mockAccountInfoRepo := new(mocks.MockAccountInfoRepository)
	validAccountInfo := utils.CreateValidAccountInfo()
	invalidUserIdAccountInfo := utils.CreateInvalidUserIdAccountInfo()

	t.Run("success", func(t *testing.T) {
		mockAccountInfoRepo.On("CreateAccountInfo", context.Background(), validAccountInfo).Return(validAccountInfo, nil)

		createdAccountInfo, err := mockAccountInfoRepo.CreateAccountInfo(context.Background(), validAccountInfo)

		assert.NoError(t, err)
		assert.NotNil(t, createdAccountInfo)

		mockAccountInfoRepo.AssertExpectations(t)
	})

	t.Run("error on invalid user id", func(t *testing.T) {
		mockAccountInfoRepo.On("CreateAccountInfo", context.Background(), invalidUserIdAccountInfo).Return(nil, errors.New("falha ao converter UserId para ObjectId"))

		_, err := mockAccountInfoRepo.CreateAccountInfo(context.Background(), invalidUserIdAccountInfo)

		assert.Error(t, err)

		mockAccountInfoRepo.AssertExpectations(t)
	})
}
