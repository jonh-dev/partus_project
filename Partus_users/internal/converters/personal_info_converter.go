package converters

import (
	"errors"
	"fmt"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToModelPersonalInfo(userId string, personalInfo *api.PersonalInfo) (*model.PersonalInfo, error) {
	if personalInfo == nil {
		return nil, errors.New("personalInfo não pode ser nil")
	}

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, fmt.Errorf("falha ao converter userId para ObjectID: %w", err)
	}

	return &model.PersonalInfo{
		UserId:       objectId,
		FirstName:    personalInfo.FirstName,
		LastName:     personalInfo.LastName,
		Email:        personalInfo.Email,
		BirthDate:    personalInfo.BirthDate.AsTime(),
		Phone:        personalInfo.Phone,
		ProfileImage: personalInfo.ProfileImage,
	}, nil
}
