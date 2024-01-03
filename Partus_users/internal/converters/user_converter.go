package converters

import (
	"errors"
	"fmt"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToModelUser(user *api.User) (*model.User, error) {
	if user == nil {
		return nil, errors.New("user não pode ser nil")
	}

	objectId := primitive.NewObjectID()

	modelPersonalInfo, err := ToModelPersonalInfo(objectId, user.PersonalInfo)
	if err != nil {
		return nil, fmt.Errorf("falha ao converter PersonalInfo para o modelo: %w", err)
	}

	modelAccountInfo, err := ToModelAccountInfo(objectId, user.AccountInfo)
	if err != nil {
		return nil, fmt.Errorf("falha ao converter AccountInfo para o modelo: %w", err)
	}

	return &model.User{
		Id:           objectId,
		PersonalInfo: *modelPersonalInfo,
		AccountInfo:  *modelAccountInfo,
	}, nil
}
