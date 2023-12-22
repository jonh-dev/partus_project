package services

import (
	"context"
	"log"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/repositories"
	"github.com/jonh-dev/partus_users/internal/validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService interface {
	CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error)
	GetUser(ctx context.Context, req *api.GetUserRequest) (*api.UserResponse, error)
	UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UserResponse, error)
	DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.UserResponse, error)
	HandleFailedLogin(ctx context.Context, req *api.HandleFailedLoginRequest) (*api.UserResponse, error)
}

type userService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error) {
	err := validation.ValidatePersonalInfo(req.PersonalInfo)
	if err != nil {
		log.Printf("Erro na validação do usuário: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro na validação do usuário: %v", err)
	}

	err = validation.ValidateAccountInfo(req.AccountInfo, validation.Create, nil)
	if err != nil {
		log.Printf("Erro na validação da conta: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro na validação do usuário: %v", err)
	}

	existingUser, err := s.repo.GetUserByEmail(ctx, req.PersonalInfo.Email)
	if err != nil {
		log.Printf("Erro ao buscar o usuário: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao buscar o usuário: %v", err)
	}

	if existingUser != nil {
		log.Printf("E-mail já está em uso: %v", req.PersonalInfo.Email)
		return nil, status.Errorf(codes.AlreadyExists, "E-mail já está em uso")
	}

	now := timestamppb.Now()
	req.AccountInfo.CreatedAt = now
	req.AccountInfo.UpdatedAt = now

	user, err := s.repo.CreateUser(ctx, req.PersonalInfo, req.AccountInfo)
	if err != nil {
		log.Printf("Erro ao criar o usuário: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao criar o usuário: %v", err)
	}

	log.Printf("Usuário criado com sucesso: %v", user.PersonalInfo.Email)
	return &api.UserResponse{User: user, Message: "Usuário criado com sucesso"}, nil
}

func (s *userService) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Erro ao obter o usuário: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "Usuário não encontrado")
	}

	return &api.UserResponse{User: user, Message: "Usuário obtido com sucesso"}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UserResponse, error) {
	return &api.UserResponse{}, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.UserResponse, error) {
	return &api.UserResponse{}, nil
}

func (s *userService) HandleFailedLogin(ctx context.Context, req *api.HandleFailedLoginRequest) (*api.UserResponse, error) {
	user, err := s.repo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		log.Printf("Erro ao buscar o usuário: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao buscar o usuário: %v", err)
	}

	user.AccountInfo.FailedLoginAttempts++
	user.AccountInfo.LastFailedLogin = timestamppb.Now()
	user.AccountInfo.LastFailedLoginReason = req.Reason

	err = validation.ValidateLoginFields(user.AccountInfo)
	if err != nil {
		log.Printf("Erro na validação da conta: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Erro na validação do usuário: %v", err)
	}

	updatedUser, err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		log.Printf("Erro ao atualizar o usuário: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao atualizar o usuário: %v", err)
	}

	return &api.UserResponse{
		User:    updatedUser,
		Message: "Login falhou, usuário atualizado com sucesso",
	}, nil
}
