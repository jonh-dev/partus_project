package model

import (
	"time"

	"github.com/jonh-dev/partus_users/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountInfo struct {
	UserId                primitive.ObjectID `bson:"userId,omitempty"`
	Username              string             `bson:"username,omitempty"`
	Password              string             `bson:"password,omitempty"`
	AccountStatus         AccountStatus      `bson:"accountStatus,omitempty"`
	StatusReason          string             `bson:"statusReason,omitempty"`
	CreatedAt             time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt             time.Time          `bson:"updatedAt,omitempty"`
	LastLogin             time.Time          `bson:"lastLogin,omitempty"`
	FailedLoginAttempts   int32              `bson:"failedLoginAttempts,omitempty"`
	LastFailedLogin       time.Time          `bson:"lastFailedLogin,omitempty"`
	LastFailedLoginReason string             `bson:"lastFailedLoginReason,omitempty"`
	AccountLockedUntil    time.Time          `bson:"accountLockedUntil,omitempty"`
	AccountLockedReason   string             `bson:"accountLockedReason,omitempty"`
}

func (a *AccountInfo) ToProto() *api.AccountInfo {
	return &api.AccountInfo{
		UserId:                a.UserId.Hex(),
		Username:              a.Username,
		Password:              a.Password,
		AccountStatus:         api.AccountStatus(a.AccountStatus),
		StatusReason:          a.StatusReason,
		CreatedAt:             timestamppb.New(a.CreatedAt),
		UpdatedAt:             timestamppb.New(a.UpdatedAt),
		LastLogin:             timestamppb.New(a.LastLogin),
		FailedLoginAttempts:   a.FailedLoginAttempts,
		LastFailedLogin:       timestamppb.New(a.LastFailedLogin),
		LastFailedLoginReason: a.LastFailedLoginReason,
		AccountLockedUntil:    timestamppb.New(a.AccountLockedUntil),
		AccountLockedReason:   a.AccountLockedReason,
	}
}
