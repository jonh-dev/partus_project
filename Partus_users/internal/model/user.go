package model

import (
	"github.com/jonh-dev/partus_users/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	PersonalInfo PersonalInfo       `bson:"personalInfo,omitempty"`
	AccountInfo  AccountInfo        `bson:"accountInfo,omitempty"`
}

func (u *User) ToProto() *api.User {
	return &api.User{
		Id:           u.Id.Hex(),
		PersonalInfo: u.PersonalInfo.ToProto(),
		AccountInfo:  u.AccountInfo.ToProto(),
	}
}
