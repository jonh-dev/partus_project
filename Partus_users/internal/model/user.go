package model

import (
	"time"

	"github.com/jonh-dev/partus_users/api"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountStatus int32

const (
	AccountStatus_ACTIVE    AccountStatus = 0
	AccountStatus_INACTIVE  AccountStatus = 1
	AccountStatus_PENDING   AccountStatus = 2
	AccountStatus_SUSPENDED AccountStatus = 3
)

type PersonalInfo struct {
	FirstName string    `bson:"firstName,omitempty"`
	LastName  string    `bson:"lastName,omitempty"`
	Email     string    `bson:"email,omitempty"`
	BirthDate time.Time `bson:"birthDate,omitempty"`
	Phone     string    `bson:"phone,omitempty"`
}

type AccountInfo struct {
	Username              string        `bson:"username,omitempty"`
	Password              string        `bson:"password,omitempty"`
	AccountStatus         AccountStatus `bson:"accountStatus,omitempty"`
	StatusReason          string        `bson:"statusReason,omitempty"`
	CreatedAt             time.Time     `bson:"createdAt,omitempty"`
	UpdatedAt             time.Time     `bson:"updatedAt,omitempty"`
	LastLogin             time.Time     `bson:"lastLogin,omitempty"`
	FailedLoginAttempts   int32         `bson:"failedLoginAttempts,omitempty"`
	LastFailedLogin       time.Time     `bson:"lastFailedLogin,omitempty"`
	LastFailedLoginReason string        `bson:"lastFailedLoginReason,omitempty"`
	AccountLockedUntil    time.Time     `bson:"accountLockedUntil,omitempty"`
	AccountLockedReason   string        `bson:"accountLockedReason,omitempty"`
}

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	PersonalInfo *PersonalInfo      `bson:"personalInfo,omitempty"`
	AccountInfo  *AccountInfo       `bson:"accountInfo,omitempty"`
}

func (u *User) ToProto() *api.User {
	return &api.User{
		Id: u.Id.Hex(),
		PersonalInfo: &api.PersonalInfo{
			FirstName: u.PersonalInfo.FirstName,
			LastName:  u.PersonalInfo.LastName,
			Email:     u.PersonalInfo.Email,
			BirthDate: timestamppb.New(u.PersonalInfo.BirthDate),
			Phone:     u.PersonalInfo.Phone,
		},
		AccountInfo: &api.AccountInfo{
			Username:              u.AccountInfo.Username,
			Password:              u.AccountInfo.Password,
			AccountStatus:         api.AccountStatus(u.AccountInfo.AccountStatus),
			StatusReason:          u.AccountInfo.StatusReason,
			CreatedAt:             timestamppb.New(u.AccountInfo.CreatedAt),
			UpdatedAt:             timestamppb.New(u.AccountInfo.UpdatedAt),
			LastLogin:             timestamppb.New(u.AccountInfo.LastLogin),
			FailedLoginAttempts:   u.AccountInfo.FailedLoginAttempts,
			LastFailedLogin:       timestamppb.New(u.AccountInfo.LastFailedLogin),
			LastFailedLoginReason: u.AccountInfo.LastFailedLoginReason,
			AccountLockedUntil:    timestamppb.New(u.AccountInfo.AccountLockedUntil),
			AccountLockedReason:   u.AccountInfo.AccountLockedReason,
		},
	}
}
