package services

import (
	"context"
	"log"

	"github.com/jonh-dev/go-error/errors"
	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/repositories"
	"github.com/jonh-dev/partus_users/internal/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IPersonalInfoService interface {
	CreatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error)
	GetPersonalInfo(ctx context.Context, req *api.GetPersonalInfoRequest) (*api.PersonalInfo, error)
	UpdatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error)
}

type PersonalInfoService struct {
	personalInfoRepo repositories.IPersonalInfoRepository
}

func NewPersonalInfoService(personalInfoRepo repositories.IPersonalInfoRepository) *PersonalInfoService {
	return &PersonalInfoService{personalInfoRepo: personalInfoRepo}
}

func (s *PersonalInfoService) CreatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	err := validation.ValidatePersonalInfo(personalInfo, validation.Create)
	if err != nil {
		return nil, errors.New(codes.InvalidArgument, "Erro na validação das informações pessoais: "+err.Error())
	}

	emailExists, err := s.personalInfoRepo.DoesEmailExist(ctx, personalInfo.Email)
	if err != nil {
		return nil, errors.New(codes.Internal, "Erro ao verificar a existência do e-mail: "+err.Error())
	}

	if emailExists {
		return nil, errors.New(codes.AlreadyExists, "O e-mail já existe")
	}

	createdPersonalInfo, err := s.personalInfoRepo.CreatePersonalInfo(ctx, personalInfo)
	if err != nil {
		return nil, errors.New(codes.Internal, "Erro ao criar PersonalInfo: "+err.Error())
	}

	return createdPersonalInfo, nil
}

func (s *PersonalInfoService) GetPersonalInfo(ctx context.Context, req *api.GetPersonalInfoRequest) (*api.PersonalInfo, error) {
	personalInfo, err := s.personalInfoRepo.GetPersonalInfo(ctx, req.UserId)
	if err != nil {
		log.Printf("Erro ao obter PersonalInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao obter PersonalInfo: %v", err)
	}

	return personalInfo, nil
}

func (s *PersonalInfoService) UpdatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	err := validation.ValidatePersonalInfo(personalInfo, validation.Update)
	if err != nil {
		log.Printf("Erro ao validar PersonalInfo: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro ao validar PersonalInfo: %v", err)
	}

	updatedPersonalInfo, err := s.personalInfoRepo.UpdatePersonalInfo(ctx, personalInfo)
	if err != nil {
		log.Printf("Erro ao atualizar PersonalInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao atualizar PersonalInfo: %v", err)
	}

	return updatedPersonalInfo, nil
}
