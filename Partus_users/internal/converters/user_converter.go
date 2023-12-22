package converters

import (
	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToModelUser(user *api.User) *model.User {
	objectId, _ := primitive.ObjectIDFromHex(user.Id)
	return &model.User{
		Id:           objectId,
		PersonalInfo: ToModelPersonalInfo(user.PersonalInfo),
		AccountInfo:  ToModelAccountInfo(user.AccountInfo),
	}
}
