package tests

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, personalInfo *api.PersonalInfo) (*api.User, error) {
	args := m.Called(ctx, personalInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	user := &api.User{
		PersonalInfo: args.Get(0).(*api.PersonalInfo),
	}
	return user, args.Error(1)
}

func (m *MockUserRepository) GetUser(ctx context.Context, id string) (*api.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*api.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.User), args.Error(1)
}
