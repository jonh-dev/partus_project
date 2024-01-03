package mocks

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/stretchr/testify/mock"
)

type MockAccountInfoRepository struct {
	mock.Mock
}

func (m *MockAccountInfoRepository) CreateAccountInfo(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	args := m.Called(ctx, accountInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.AccountInfo), args.Error(1)
}

func (m *MockAccountInfoRepository) GetAccountInfo(ctx context.Context, username string) (*api.AccountInfo, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.AccountInfo), args.Error(1)
}

func (m *MockAccountInfoRepository) UpdateUserCredentials(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	args := m.Called(ctx, accountInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.AccountInfo), args.Error(1)
}

// Implemente os outros métodos conforme necessário...
