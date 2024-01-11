package repositories

import (
	"context"
	"fmt"

	"github.com/jonh-dev/partus_users/internal/config"
	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
}

type UserRepository struct {
	dbService *config.DBService
}

func NewUserRepository(dbService *config.DBService) IUserRepository {
	return &UserRepository{
		dbService: dbService,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	collection := r.getCollection()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("falha ao inserir usuário no banco de dados: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
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

	return &user, nil
}

func (r *UserRepository) getCollection() *mongo.Collection {
	return r.dbService.Client.Database(r.dbService.DBName).Collection("users")
}
