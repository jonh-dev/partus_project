package converters

import (
	"errors"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToModelPersonalInfo(objectId primitive.ObjectID, personalInfo *api.PersonalInfo) (*model.PersonalInfo, error) {
	if personalInfo == nil {
		return nil, errors.New("personalInfo não pode ser nil")
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
