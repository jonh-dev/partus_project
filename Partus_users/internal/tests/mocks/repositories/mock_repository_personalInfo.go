package mocks

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/stretchr/testify/mock"
)

type MockPersonalInfoRepository struct {
	mock.Mock
}

func (m *MockPersonalInfoRepository) CreatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	args := m.Called(ctx, personalInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.PersonalInfo), args.Error(1)
}

func (m *MockPersonalInfoRepository) GetPersonalInfo(ctx context.Context, id string) (*api.PersonalInfo, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.PersonalInfo), args.Error(1)
}

func (m *MockPersonalInfoRepository) UpdatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	args := m.Called(ctx, personalInfo)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.PersonalInfo), args.Error(1)
}

func (m *MockPersonalInfoRepository) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return false, args.Error(1)
	}
	return args.Get(0).(bool), args.Error(1)
}
