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

type IPersonalInfoRepository interface {
	CreatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error)
	GetPersonalInfo(ctx context.Context, id string) (*api.PersonalInfo, error)
	UpdatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error)
	DoesEmailExist(ctx context.Context, email string) (bool, error)
}

type PersonalInfoRepository struct {
	dbService *config.DBService
}

func NewPersonalInfoRepository(dbService *config.DBService) IPersonalInfoRepository {
	return &PersonalInfoRepository{
		dbService: dbService,
	}
}

func (r *PersonalInfoRepository) CreatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	collection := r.getCollection()

	userId, err := utils.ConvertToObjectId(personalInfo.UserId)
	if err != nil {
		return nil, err
	}

	dbPersonalInfo := &model.PersonalInfo{
		UserId:       userId,
		FirstName:    personalInfo.FirstName,
		LastName:     personalInfo.LastName,
		Email:        personalInfo.Email,
		BirthDate:    personalInfo.BirthDate.AsTime(),
		Phone:        personalInfo.Phone,
		ProfileImage: personalInfo.ProfileImage,
	}

	_, err = collection.InsertOne(ctx, dbPersonalInfo)
	if err != nil {
		return nil, fmt.Errorf("falha ao inserir PersonalInfo no banco de dados: %w", err)
	}

	return personalInfo, nil
}

func (r *PersonalInfoRepository) GetPersonalInfo(ctx context.Context, id string) (*api.PersonalInfo, error) {
	collection := r.getCollection()

	userId, err := utils.ConvertToObjectId(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userId": userId}
	dbPersonalInfo := &model.PersonalInfo{}
	err = collection.FindOne(ctx, filter).Decode(dbPersonalInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "PersonalInfo não encontrado")
		}
		return nil, fmt.Errorf("falha ao buscar PersonalInfo do banco de dados: %w", err)
	}

	personalInfo := &api.PersonalInfo{
		UserId:       id,
		FirstName:    dbPersonalInfo.FirstName,
		LastName:     dbPersonalInfo.LastName,
		Email:        dbPersonalInfo.Email,
		BirthDate:    timestamppb.New(dbPersonalInfo.BirthDate),
		Phone:        dbPersonalInfo.Phone,
		ProfileImage: dbPersonalInfo.ProfileImage,
	}

	return personalInfo, nil
}

func (r *PersonalInfoRepository) UpdatePersonalInfo(ctx context.Context, personalInfo *api.PersonalInfo) (*api.PersonalInfo, error) {
	collection := r.getCollection()

	userId, err := utils.ConvertToObjectId(personalInfo.UserId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userId": userId}
	update := bson.M{
		"$set": bson.M{
			"firstName":    personalInfo.FirstName,
			"lastName":     personalInfo.LastName,
			"email":        personalInfo.Email,
			"birthDate":    personalInfo.BirthDate.AsTime(),
			"phone":        personalInfo.Phone,
			"profileImage": personalInfo.ProfileImage,
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("falha ao atualizar PersonalInfo no banco de dados: %w", err)
	}

	return personalInfo, nil
}

func (r *PersonalInfoRepository) DoesEmailExist(ctx context.Context, email string) (bool, error) {
	collection := r.getCollection()

	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, fmt.Errorf("falha ao buscar PersonalInfo do banco de dados: %w", err)
	}

	return true, nil
}

func (r *PersonalInfoRepository) getCollection() *mongo.Collection {
	return r.dbService.Client.Database(r.dbService.DBName).Collection("personal_info")
}
