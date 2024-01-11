package services

import (
	"context"
	"fmt"

	"github.com/jonh-dev/go-error/errors"
	"github.com/jonh-dev/go-logger/logger"
	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/converters"
	"github.com/jonh-dev/partus_users/internal/repositories"
	"github.com/jonh-dev/partus_users/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
)

type UserService interface {
	CreateUser(ctx context.Context, req *api.CreateUserRequest) (*api.UserResponse, error)
	GetUser(ctx context.Context, req *api.GetUserRequest) (*api.UserResponse, error)
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
		logger.Error("Erro ao converter o usuário para o modelo: " + err.Error())
		return nil, errors.New(codes.Internal, "Erro ao converter o usuário para o modelo: "+err.Error())
	}

	apiPersonalInfo := modelUser.PersonalInfo.ToProto()
	_, err = s.personalInfoService.CreatePersonalInfo(ctx, apiPersonalInfo)
	if err != nil {
		if e, ok := err.(*errors.Error); ok {
			logger.Error("Erro ao criar usuário: " + e.Error())
			return nil, errors.New(e.GRPCStatus().Code(), "Erro ao criar usuário: "+e.Error())
		}
		logger.Error(err.Error())
		return nil, err
	}

	apiAccountInfo := modelUser.AccountInfo.ToProto()
	_, err = s.accountInfoService.CreateAccountInfo(ctx, apiAccountInfo)
	if err != nil {
		if e, ok := err.(*errors.Error); ok {
			return nil, e.GRPCStatus().Err()
		}
		return nil, errors.New(codes.Internal, "Erro desconhecido ao criar AccountInfo")
	}

	modelUser.AccountInfo.Password = apiAccountInfo.Password
	modelUser.AccountInfo.CreatedAt = apiAccountInfo.CreatedAt.AsTime()

	user, err := s.userRepo.CreateUser(ctx, modelUser)
	if err != nil {
		return nil, errors.New(codes.Internal, "Erro ao criar o usuário: "+err.Error())
	}

	apiUser := user.ToProto()

	logger.Success(fmt.Sprintf("Usuário criado com sucesso: ID: %s, Nome: %s %s, Email: %s", apiUser.Id, apiUser.PersonalInfo.FirstName, apiUser.PersonalInfo.LastName, apiUser.PersonalInfo.Email))
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

	modelUser.AccountInfo.CreatedAt = utils.ReadjustToSaoPaulo(modelUser.AccountInfo.CreatedAt)

	apiUser := modelUser.ToProto()

	return &api.UserResponse{
		User:    apiUser,
		Message: "Usuário obtido com sucesso",
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
