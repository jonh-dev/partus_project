package services

import (
	"context"
	"fmt"
	"log"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/converters"
	"github.com/jonh-dev/partus_users/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error)
	GetUser(ctx context.Context, req *api.GetUserRequest) (*api.UserResponse, error)
	UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UserResponse, error)
	DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.UserResponse, error)
	HandleFailedLogin(ctx context.Context, req *api.HandleFailedLoginRequest) (*api.UserResponse, error)
}

type userService struct {
	userRepo            repositories.IUserRepository
	personalInfoService IPersonalInfoService
	accountInfoService  IAccountInfoService
}

func NewUserService(userRepo repositories.IUserRepository, personalInfoService IPersonalInfoService, accountInfoService IAccountInfoService) *userService {
	return &userService{
		userRepo:            userRepo,
		personalInfoService: personalInfoService,
		accountInfoService:  accountInfoService,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error) {
	modelUser, err := converters.ToModelUser(req.User)
	if err != nil {
		log.Printf("Erro ao converter o usuário para o modelo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao converter o usuário para o modelo: %v", err)
	}

	user, err := s.userRepo.CreateUser(ctx, modelUser)
	if err != nil {
		log.Printf("Erro ao criar o usuário: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao criar o usuário: %v", err)
	}

	apiPersonalInfo := modelUser.PersonalInfo.ToProto()
	_, err = s.personalInfoService.CreatePersonalInfo(ctx, apiPersonalInfo)
	if err != nil {
		log.Printf("Erro ao criar PersonalInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao criar PersonalInfo: %v", err)
	}

	apiAccountInfo := modelUser.AccountInfo.ToProto()
	_, err = s.accountInfoService.CreateAccountInfo(ctx, apiAccountInfo)
	if err != nil {
		log.Printf("Erro ao criar AccountInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao criar AccountInfo: %v", err)
	}

	apiUser := user.ToProto()

	return &api.UserResponse{
		User:    apiUser,
		Message: "Usuário criado com sucesso",
	}, nil
}

func (s *userService) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.UserResponse, error) {
	modelUser, err := s.userRepo.GetUser(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("falha ao obter User: %w", err)
	}

	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("falha ao converter userId para ObjectID: %w", err)
	}

	apiPersonalInfo, err := s.personalInfoService.GetPersonalInfo(ctx, &api.GetPersonalInfoRequest{UserId: req.Id})
	if err != nil {
		return nil, fmt.Errorf("falha ao obter PersonalInfo: %w", err)
	}

	modelPersonalInfo, err := converters.ToModelPersonalInfo(objectId, apiPersonalInfo)
	if err != nil {
		return nil, fmt.Errorf("falha ao converter PersonalInfo para o modelo: %w", err)
	}

	modelUser.PersonalInfo = *modelPersonalInfo

	apiAccountInfo, err := s.accountInfoService.GetAccountInfo(ctx, &api.GetAccountInfoRequest{UserId: req.Id})
	if err != nil {
		return nil, fmt.Errorf("falha ao obter AccountInfo: %w", err)
	}

	modelAccountInfo, err := converters.ToModelAccountInfo(objectId, apiAccountInfo)
	if err != nil {
		return nil, fmt.Errorf("falha ao converter AccountInfo para o modelo: %w", err)
	}
	modelUser.AccountInfo = *modelAccountInfo

	apiUser := modelUser.ToProto()

	return &api.UserResponse{
		User:    apiUser,
		Message: "Usuário obtido com sucesso",
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UserResponse, error) {
	modelUser, err := converters.ToModelUser(req.User)
	if err != nil {
		log.Printf("Erro ao converter o usuário para o modelo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao converter o usuário para o modelo: %v", err)
	}

	user, err := s.userRepo.UpdateUser(ctx, modelUser)
	if err != nil {
		log.Printf("Erro ao atualizar o usuário: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao atualizar o usuário: %v", err)
	}

	apiPersonalInfo := modelUser.PersonalInfo.ToProto()
	_, err = s.personalInfoService.UpdatePersonalInfo(ctx, apiPersonalInfo)
	if err != nil {
		log.Printf("Erro ao atualizar PersonalInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao atualizar PersonalInfo: %v", err)
	}

	apiAccountInfo := modelUser.AccountInfo.ToProto()
	_, err = s.accountInfoService.UpdateUserCredentials(ctx, apiAccountInfo)
	if err != nil {
		log.Printf("Erro ao atualizar AccountInfo: %v", err)
		return nil, status.Errorf(codes.Internal, "Erro ao atualizar AccountInfo: %v", err)
	}

	apiUser := user.ToProto()

	return &api.UserResponse{
		User:    apiUser,
		Message: "Usuário atualizado com sucesso",
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, req *api.DeleteUserRequest) (*api.UserResponse, error) {
	// Implemente a lógica de exclusão do usuário aqui
	// ...
	return nil, nil
}

func (s *userService) HandleFailedLogin(ctx context.Context, req *api.HandleFailedLoginRequest) (*api.UserResponse, error) {
	// Implemente a lógica de manipulação de falha de login aqui
	// ...
	return nil, nil
}
