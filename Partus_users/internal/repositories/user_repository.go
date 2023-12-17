package repositories

import (
	"context"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, personalInfo *api.PersonalInfo) (*api.User, error)
	GetUser(ctx context.Context, id string) (*api.User, error)
	GetUserByEmail(ctx context.Context, email string) (*api.User, error)
}

type UserRepository struct {
	dbService *config.DBService
}

func NewUserRepository(dbService *config.DBService) IUserRepository {
	return &UserRepository{
		dbService: dbService,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, personalInfo *api.PersonalInfo) (*api.User, error) {
	collection := r.getCollection()

	user := &api.User{
		PersonalInfo: personalInfo,
	}

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*api.User, error) {
	collection := r.getCollection()

	var user *api.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID}
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*api.User, error) {
	collection := r.getCollection()

	var user *api.User
	filter := bson.M{"personalInfo.email": email}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) getCollection() *mongo.Collection {
	return r.dbService.Client.Database(r.dbService.DBName).Collection("users")
}
