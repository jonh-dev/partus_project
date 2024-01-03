package services_test

import (
	"context"
	"testing"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/services"
	repository "github.com/jonh-dev/partus_users/internal/tests/mocks/repositories"
	mocks "github.com/jonh-dev/partus_users/internal/tests/mocks/services"
	"github.com/jonh-dev/partus_users/internal/tests/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateUser(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	mockPersonalInfoService := new(mocks.MockPersonalInfoService)
	mockAccountInfoService := new(mocks.MockAccountInfoService)

	validUser := utils.CreateValidUser()

	validCreateUserRequest := &api.CreateUserRequest{
		User: validUser.ToProto(),
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*model.User")).Return(validUser, nil)
		mockPersonalInfoService.On("CreatePersonalInfo", mock.Anything, mock.AnythingOfType("*api.PersonalInfo")).Return(validUser.PersonalInfo.ToProto(), nil)
		mockAccountInfoService.On("CreateAccountInfo", mock.Anything, mock.AnythingOfType("*api.AccountInfo")).Return(validUser.AccountInfo.ToProto(), nil)

		u := services.NewUserService(mockUserRepo, mockPersonalInfoService, mockAccountInfoService)
		user, err := u.CreateUser(context.Background(), validCreateUserRequest)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepo.AssertExpectations(t)
		mockPersonalInfoService.AssertExpectations(t)
		mockAccountInfoService.AssertExpectations(t)
	})
}
