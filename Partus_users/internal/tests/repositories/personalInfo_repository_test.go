package repositories

import (
	"context"
	"errors"
	"testing"

	mocks "github.com/jonh-dev/partus_users/internal/tests/mocks/repositories"
	"github.com/jonh-dev/partus_users/internal/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestPersonalInfoRepository_CreatePersonalInfo(t *testing.T) {
	mockPersonalInfoRepo := new(mocks.MockPersonalInfoRepository)
	validPersonalInfo := utils.CreateValidPersonalInfo()
	invalidUserIdPersonalInfo := utils.CreateInvalidUserIdPersonalInfo()

	t.Run("success", func(t *testing.T) {
		mockPersonalInfoRepo.On("CreatePersonalInfo", context.Background(), validPersonalInfo).Return(validPersonalInfo, nil)

		createdPersonalInfo, err := mockPersonalInfoRepo.CreatePersonalInfo(context.Background(), validPersonalInfo)

		assert.NoError(t, err)
		assert.NotNil(t, createdPersonalInfo)

		mockPersonalInfoRepo.AssertExpectations(t)
	})

	t.Run("error on invalid user id", func(t *testing.T) {
		mockPersonalInfoRepo.On("CreatePersonalInfo", context.Background(), invalidUserIdPersonalInfo).Return(nil, errors.New("falha ao converter UserId para ObjectId"))

		_, err := mockPersonalInfoRepo.CreatePersonalInfo(context.Background(), invalidUserIdPersonalInfo)

		assert.Error(t, err)

		mockPersonalInfoRepo.AssertExpectations(t)
	})

}
