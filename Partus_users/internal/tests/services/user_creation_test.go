package services

import (
	"testing"
	"time"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/tests"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUserCreation(t *testing.T) {
	testCases := []tests.UserTest{
		{
			Name: "Valid User",
			Request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.valid@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
				AccountInfo: &api.AccountInfo{
					Username:              "joao.valid",
					Password:              "ValidPassword123!",
					AccountStatus:         api.AccountStatus_ACTIVE,
					StatusReason:          "",
					CreatedAt:             timestamppb.Now(),
					UpdatedAt:             timestamppb.Now(),
					LastLogin:             timestamppb.Now(),
					FailedLoginAttempts:   0,
					LastFailedLogin:       timestamppb.Now(),
					LastFailedLoginReason: "",
					AccountLockedUntil:    timestamppb.Now(),
					AccountLockedReason:   "",
				},
			},
			WantErr:       false,
			ExpectedError: "",
		},
		{
			Name: "Email Already Exists",
			Request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Carlos",
					Email:     "joao.carlos@gmail.com",
					Phone:     "11987654321",
				},
				AccountInfo: &api.AccountInfo{
					Username:              "joao.valid",
					Password:              "ValidPassword123!",
					AccountStatus:         api.AccountStatus_ACTIVE,
					StatusReason:          "",
					CreatedAt:             timestamppb.Now(),
					UpdatedAt:             timestamppb.Now(),
					LastLogin:             timestamppb.Now(),
					FailedLoginAttempts:   0,
					LastFailedLogin:       timestamppb.Now(),
					LastFailedLoginReason: "",
					AccountLockedUntil:    timestamppb.Now(),
					AccountLockedReason:   "",
				},
			},
			WantErr:       true,
			ExpectedError: "E-mail já está em uso",
		},
		// Outros casos de teste para criação do usuário aqui
	}

	for _, tt := range testCases {
		tt.RunTest(t)
	}
}
