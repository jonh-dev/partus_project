package mocks

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.UserResponse), args.Error(1)
}

func (m *MockUserService) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.UserResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.UserResponse), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UserResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.UserResponse), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.UserResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.UserResponse), args.Error(1)
}

func (m *MockUserService) HandleFailedLogin(ctx context.Context, req *api.HandleFailedLoginRequest) (*api.UserResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.UserResponse), args.Error(1)
}
