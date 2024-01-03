package services

import (
	"context"
	"log"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/repositories"
	"github.com/jonh-dev/partus_users/internal/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IAccountInfoService interface {
	CreateAccountInfo(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error)
	GetAccountInfo(ctx context.Context, req *api.GetAccountInfoRequest) (*api.AccountInfo, error)
	UpdateUserCredentials(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error)
}

type AccountInfoService struct {
	accountInfoRepo repositories.IAccountInfoRepository
}

func NewAccountInfoService(accountInfoRepo repositories.IAccountInfoRepository) *AccountInfoService {
	return &AccountInfoService{accountInfoRepo: accountInfoRepo}
}

func (s *AccountInfoService) CreateAccountInfo(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	err := validation.ValidateAccountInfo(accountInfo, validation.Create, nil)
	if err != nil {
		log.Printf("Erro ao validar AccountInfo: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro ao validar AccountInfo: %v", err)
	}

	createdAccountInfo, err := s.accountInfoRepo.CreateAccountInfo(ctx, accountInfo)
	if err != nil {
		log.Printf("Erro ao criar AccountInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao criar AccountInfo: %v", err)
	}

	return createdAccountInfo, nil
}

func (s *AccountInfoService) GetAccountInfo(ctx context.Context, req *api.GetAccountInfoRequest) (*api.AccountInfo, error) {
	accountInfo, err := s.accountInfoRepo.GetAccountInfo(ctx, req.UserId)
	if err != nil {
		log.Printf("Erro ao obter AccountInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao obter AccountInfo: %v", err)
	}

	return accountInfo, nil
}

func (s *AccountInfoService) UpdateUserCredentials(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	err := validation.ValidateAccountInfo(accountInfo, validation.Update, nil)
	if err != nil {
		log.Printf("Erro ao validar AccountInfo: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro ao validar AccountInfo: %v", err)
	}

	updatedAccountInfo, err := s.accountInfoRepo.UpdateUserCredentials(ctx, accountInfo)
	if err != nil {
		log.Printf("Erro ao atualizar AccountInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao atualizar AccountInfo: %v", err)
	}

	return updatedAccountInfo, nil
}
