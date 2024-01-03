package repositories

import (
	"context"
	"testing"

	mocks "github.com/jonh-dev/partus_users/internal/tests/mocks/repositories"
	"github.com/jonh-dev/partus_users/internal/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_CreateUser(t *testing.T) {
	userRepo := &mocks.MockUserRepository{}

	validUser := utils.CreateValidUser()
	invalidUser := utils.CreateInvalidUser()

	t.Run("success", func(t *testing.T) {
		userRepo.On("CreateUser", context.Background(), validUser).Return(validUser, nil)

		createdUser, err := userRepo.CreateUser(context.Background(), validUser)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)

		userRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userRepo.On("CreateUser", context.Background(), invalidUser).Return(nil, assert.AnError)

		_, err := userRepo.CreateUser(context.Background(), invalidUser)

		assert.Error(t, err)

		userRepo.AssertExpectations(t)
	})
}
