package repositories

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/config"
	"github.com/jonh-dev/partus_users/internal/converters"
	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, personalInfo *api.PersonalInfo, accountInfo *api.AccountInfo) (*api.User, error)
	UpdateUser(ctx context.Context, user *api.User) (*api.User, error)
	GetUser(ctx context.Context, id string) (*api.User, error)
	GetUserByEmail(ctx context.Context, email string) (*api.User, error)
	GetUserByUsername(ctx context.Context, username string) (*api.User, error)
}

type UserRepository struct {
	dbService *config.DBService
}

func NewUserRepository(dbService *config.DBService) IUserRepository {
	return &UserRepository{
		dbService: dbService,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, personalInfo *api.PersonalInfo, accountInfo *api.AccountInfo) (*api.User, error) {
	collection := r.getCollection()

	modelPersonalInfo := converters.ToModelPersonalInfo(personalInfo)
	modelAccountInfo := converters.ToModelAccountInfo(accountInfo)

	user := &model.User{
		PersonalInfo: modelPersonalInfo,
		AccountInfo:  modelAccountInfo,
	}

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *api.User) (*api.User, error) {
	collection := r.getCollection()

	objectId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return nil, err
	}

	modelUser := converters.ToModelUser(user)

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{
		"personalinfo": modelUser.PersonalInfo,
		"accountinfo":  modelUser.AccountInfo,
	}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return modelUser.ToProto(), nil
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*api.User, error) {
	collection := r.getCollection()

	var user model.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.ToProto(), nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*api.User, error) {
	collection := r.getCollection()

	var user model.User
	filter := bson.M{"personalinfo.email": email}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return user.ToProto(), nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*api.User, error) {
	collection := r.getCollection()

	var user model.User
	filter := bson.M{"accountinfo.username": username}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return user.ToProto(), nil
}

func (r *UserRepository) getCollection() *mongo.Collection {
	return r.dbService.Client.Database(r.dbService.DBName).Collection("users")
}
