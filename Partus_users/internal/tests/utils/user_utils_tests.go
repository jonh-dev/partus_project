package utils

import (
	"time"

	"github.com/jonh-dev/partus_users/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateValidUser() *model.User {
	return &model.User{
		Id: primitive.NewObjectID(),
		PersonalInfo: model.PersonalInfo{
			UserId:       primitive.NewObjectID(),
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john.doe@example.com",
			BirthDate:    time.Now(),
			Phone:        "1234567890",
			ProfileImage: "profile.jpg",
		},
		AccountInfo: model.AccountInfo{
			UserId:                primitive.NewObjectID(),
			Username:              "johndoe",
			Password:              "password",
			AccountStatus:         model.AccountStatus_ACTIVE,
			StatusReason:          "",
			CreatedAt:             time.Now(),
			UpdatedAt:             time.Now(),
			LastLogin:             time.Now(),
			FailedLoginAttempts:   0,
			LastFailedLogin:       time.Time{},
			LastFailedLoginReason: "",
			AccountLockedUntil:    time.Time{},
			AccountLockedReason:   "",
		},
	}
}

func CreateInvalidUser() *model.User {
	return &model.User{
		Id: primitive.NewObjectID(),
		PersonalInfo: model.PersonalInfo{
			UserId:       primitive.NewObjectID(),
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "invalid email", // Endereço de e-mail inválido
			BirthDate:    time.Now(),
			Phone:        "1234567890",
			ProfileImage: "profile.jpg",
		},
		AccountInfo: model.AccountInfo{
			UserId:                primitive.NewObjectID(),
			Username:              "johndoe",
			Password:              "pw", // Senha muito curta
			AccountStatus:         model.AccountStatus_ACTIVE,
			StatusReason:          "",
			CreatedAt:             time.Now(),
			UpdatedAt:             time.Now(),
			LastLogin:             time.Now(),
			FailedLoginAttempts:   0,
			LastFailedLogin:       time.Time{},
			LastFailedLoginReason: "",
			AccountLockedUntil:    time.Time{},
			AccountLockedReason:   "",
		},
	}
}
