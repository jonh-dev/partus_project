package mocks

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/stretchr/testify/mock"
)

type MockPersonalInfoService struct {
	mock.Mock
}

func (m *MockPersonalInfoService) CreatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	args := m.Called(ctx, personalInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.PersonalInfo), args.Error(1)
}

func (m *MockPersonalInfoService) GetPersonalInfo(ctx context.Context, req *api.GetPersonalInfoRequest) (*api.PersonalInfo, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.PersonalInfo), args.Error(1)
}

func (m *MockPersonalInfoService) UpdatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	args := m.Called(ctx, personalInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.PersonalInfo), args.Error(1)
}
