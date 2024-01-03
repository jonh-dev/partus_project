package mocks

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/stretchr/testify/mock"
)

type MockAccountInfoService struct {
	mock.Mock
}

func (m *MockAccountInfoService) CreateAccountInfo(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	args := m.Called(ctx, accountInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.AccountInfo), args.Error(1)
}

func (m *MockAccountInfoService) GetAccountInfo(ctx context.Context, req *api.GetAccountInfoRequest) (*api.AccountInfo, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.AccountInfo), args.Error(1)
}

func (m *MockAccountInfoService) UpdateUserCredentials(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	args := m.Called(ctx, accountInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.AccountInfo), args.Error(1)
}
