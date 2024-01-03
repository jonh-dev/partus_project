package model

import (
	"time"

	"github.com/jonh-dev/partus_users/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PersonalInfo struct {
	UserId       primitive.ObjectID `bson:"userId,omitempty"`
	FirstName    string             `bson:"firstName,omitempty"`
	LastName     string             `bson:"lastName,omitempty"`
	Email        string             `bson:"email,omitempty"`
	BirthDate    time.Time          `bson:"birthDate,omitempty"`
	Phone        string             `bson:"phone,omitempty"`
	ProfileImage string             `bson:"profileImage,omitempty"`
}

func (p *PersonalInfo) ToProto() *api.PersonalInfo {
	return &api.PersonalInfo{
		UserId:       p.UserId.Hex(),
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Email:        p.Email,
		BirthDate:    timestamppb.New(p.BirthDate),
		Phone:        p.Phone,
		ProfileImage: p.ProfileImage,
	}
}
