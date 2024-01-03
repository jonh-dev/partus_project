package repositories

import (
	"context"
	"fmt"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/config"
	"github.com/jonh-dev/partus_users/internal/model"
	"github.com/jonh-dev/partus_users/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IAccountInfoRepository interface {
	CreateAccountInfo(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error)
	GetAccountInfo(ctx context.Context, id string) (*api.AccountInfo, error)
	UpdateUserCredentials(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error)
}

type AccountInfoRepository struct {
	dbService *config.DBService
}

func NewAccountInfoRepository(dbService *config.DBService) IAccountInfoRepository {
	return &AccountInfoRepository{
		dbService: dbService,
	}
}

func (r *AccountInfoRepository) CreateAccountInfo(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	collection := r.getCollection()

	userId, err := utils.ConvertToObjectId(accountInfo.UserId)
	if err != nil {
		return nil, err
	}

	dbAccountInfo := &model.AccountInfo{
		UserId:                userId,
		Username:              accountInfo.Username,
		Password:              accountInfo.Password,
		AccountStatus:         model.AccountStatus(accountInfo.AccountStatus),
		StatusReason:          accountInfo.StatusReason,
		CreatedAt:             accountInfo.CreatedAt.AsTime(),
		UpdatedAt:             accountInfo.UpdatedAt.AsTime(),
		LastLogin:             accountInfo.LastLogin.AsTime(),
		FailedLoginAttempts:   accountInfo.FailedLoginAttempts,
		LastFailedLogin:       accountInfo.LastFailedLogin.AsTime(),
		LastFailedLoginReason: accountInfo.LastFailedLoginReason,
		AccountLockedUntil:    accountInfo.AccountLockedUntil.AsTime(),
		AccountLockedReason:   accountInfo.AccountLockedReason,
	}

	_, err = collection.InsertOne(ctx, dbAccountInfo)
	if err != nil {
		return nil, fmt.Errorf("falha ao inserir AccountInfo no banco de dados: %w", err)
	}

	return accountInfo, nil
}

func (r *AccountInfoRepository) GetAccountInfo(ctx context.Context, id string) (*api.AccountInfo, error) {
	collection := r.getCollection()

	userId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userId": userId}
	dbAccountInfo := &model.AccountInfo{}
	err = collection.FindOne(ctx, filter).Decode(dbAccountInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "AccountInfo não encontrado")
		}
		return nil, fmt.Errorf("falha ao buscar AccountInfo do banco de dados: %w", err)
	}

	accountInfo := &api.AccountInfo{
		UserId:                id,
		Username:              dbAccountInfo.Username,
		Password:              dbAccountInfo.Password,
		AccountStatus:         api.AccountStatus(dbAccountInfo.AccountStatus),
		StatusReason:          dbAccountInfo.StatusReason,
		CreatedAt:             timestamppb.New(dbAccountInfo.CreatedAt),
		UpdatedAt:             timestamppb.New(dbAccountInfo.UpdatedAt),
		LastLogin:             timestamppb.New(dbAccountInfo.LastLogin),
		FailedLoginAttempts:   dbAccountInfo.FailedLoginAttempts,
		LastFailedLogin:       timestamppb.New(dbAccountInfo.LastFailedLogin),
		LastFailedLoginReason: dbAccountInfo.LastFailedLoginReason,
		AccountLockedUntil:    timestamppb.New(dbAccountInfo.AccountLockedUntil),
		AccountLockedReason:   dbAccountInfo.AccountLockedReason,
	}

	return accountInfo, nil
}

func (r *AccountInfoRepository) UpdateUserCredentials(ctx context.Context, accountInfo *api.AccountInfo) (*api.AccountInfo, error) {
	collection := r.getCollection()

	userId, err := utils.ConvertToObjectId(accountInfo.UserId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userId": userId}
	update := bson.M{
		"$set": bson.M{
			"username": accountInfo.Username,
			"password": accountInfo.Password,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("falha ao atualizar UserCredentials no banco de dados: %w", err)
	}

	return accountInfo, nil
}

func (r *AccountInfoRepository) getCollection() *mongo.Collection {
	return r.dbService.Client.Database(r.dbService.DBName).Collection("account_info")
}
