package tests

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, personalInfo *api.PersonalInfo, accountInfo *api.AccountInfo) (*api.User, error) {
	args := m.Called(ctx, personalInfo, accountInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	user := args.Get(0).(*api.User)
	return user, args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *api.User) (*api.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.User), args.Error(1)
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

func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (*api.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.User), args.Error(1)
}
