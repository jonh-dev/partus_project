package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/services"
	"github.com/jonh-dev/partus_users/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(tests.MockUserRepository)

	srv := services.NewUserService(mockRepo)

	tests := []struct {
		name          string
		request       *api.CreateUserRequest
		wantErr       bool
		expectedError string
	}{
		{
			name: "Valid User",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.valid@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       false,
			expectedError: "",
		},
		{
			name: "Invalid First Name - Lowercase",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "joão",
					LastName:  "Silva",
					Email:     "joao.lowercase@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o primeiro nome deve começar com uma letra maiúscula, conter apenas uma palavra e ter no máximo 20 caracteres",
		},
		{
			name: "Invalid First Name - Multiple Words",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João Carlos",
					LastName:  "Silva",
					Email:     "joao.carlos@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o primeiro nome deve começar com uma letra maiúscula, conter apenas uma palavra e ter no máximo 20 caracteres",
		},
		{
			name: "Invalid First Name - Contains Numbers",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João1",
					LastName:  "Silva",
					Email:     "joao.numbers@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o primeiro nome deve começar com uma letra maiúscula, conter apenas uma palavra e ter no máximo 20 caracteres",
		},
		{
			name: "Invalid First Name - Too Long",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "Joãooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo",
					LastName:  "Silva",
					Email:     "joao.long@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o primeiro nome deve começar com uma letra maiúscula, conter apenas uma palavra e ter no máximo 20 caracteres",
		},
		{
			name: "Invalid Last Name - Lowercase",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "silva",
					Email:     "joao.lowercase@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o sobrenome deve começar com uma letra maiúscula em cada palavra e ter no máximo 50 caracteres",
		},
		{
			name: "Invalid Last Name - Contains Numbers",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva1",
					Email:     "joao.numbers@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o sobrenome deve começar com uma letra maiúscula em cada palavra e ter no máximo 50 caracteres",
		},
		{
			name: "Invalid Last Name - Too Long",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silvaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
					Email:     "joao.long@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o sobrenome deve começar com uma letra maiúscula em cada palavra e ter no máximo 50 caracteres",
		},
		{
			name: "Invalid Email - Missing @",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.missingatexample.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o e-mail deve ser um endereço de e-mail válido",
		},
		{
			name: "Invalid Email - Missing Domain",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.missingdomain@",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o e-mail deve ser um endereço de e-mail válido",
		},
		{
			name: "Invalid Email - Missing Username",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o e-mail deve ser um endereço de e-mail válido",
		},
		{
			name: "Invalid Email - Missing Username and Domain",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "@",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o e-mail deve ser um endereço de e-mail válido",
		},
		{
			name: "Invalid Birth Date - Future Date",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.future@example.com",
					BirthDate: timestamppb.New(time.Now().AddDate(0, 0, 1)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: a data de nascimento deve estar no passado, o usuário deve ter pelo menos 13 anos e o ano deve ser entre 1900 e o ano atual",
		},
		{
			name: "Invalid Birth Date - User Too Young",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.young@example.com",
					BirthDate: timestamppb.New(time.Now().AddDate(-12, 0, 0)),
					Phone:     "(41) 99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: a data de nascimento deve estar no passado, o usuário deve ter pelo menos 13 anos e o ano deve ser entre 1900 e o ano atual",
		},
		{
			name: "Invalid Phone - Missing Area Code",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.invalid@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "99999-9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o telefone ou celular deve estar no formato correto, ou seja, começar com '+' seguido de 1 a 3 dígitos para números internacionais, ou começar diretamente com um dígito para números brasileiros, e ter entre 9 e 14 dígitos no total, sem conter nenhum caractere que não seja dígito ou '+'",
		},
		{
			name: "Invalid Phone - Wrong Format",
			request: &api.CreateUserRequest{
				PersonalInfo: &api.PersonalInfo{
					FirstName: "João",
					LastName:  "Silva",
					Email:     "joao.invalid@example.com",
					BirthDate: timestamppb.New(time.Date(2000, 8, 10, 0, 0, 0, 0, time.UTC)),
					Phone:     "(41) 99999 9999",
				},
			},
			wantErr:       true,
			expectedError: "Erro na validação do usuário: o telefone ou celular deve estar no formato correto, ou seja, começar com '+' seguido de 1 a 3 dígitos para números internacionais, ou começar diretamente com um dígito para números brasileiros, e ter entre 9 e 14 dígitos no total, sem conter nenhum caractere que não seja dígito ou '+'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = []*mock.Call{}

			mockRepo.On("GetUserByEmail", mock.Anything, tt.request.PersonalInfo.Email).Return(nil, nil)

			mockRepo.On("CreateUser", mock.Anything, tt.request.PersonalInfo).Return(tt.request.PersonalInfo, nil)

			user, err := srv.CreateUser(context.Background(), tt.request)
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, fmt.Sprintf("rpc error: code = InvalidArgument desc = %s", tt.expectedError))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.request.PersonalInfo, user.User.PersonalInfo)
			}
		})
	}

	t.Run("Email Already Exists", func(t *testing.T) {
		mockRepo.ExpectedCalls = []*mock.Call{}

		personalInfo := &api.PersonalInfo{
			FirstName: "João",
			LastName:  "Carlos",
			Email:     "joao.carlos@gmail.com",
			Phone:     "11987654321",
		}

		existingUser := &api.User{
			PersonalInfo: personalInfo,
		}

		mockRepo.On("GetUserByEmail", context.Background(), personalInfo.Email).Return(existingUser, nil)

		req := &api.CreateUserRequest{
			PersonalInfo: personalInfo,
		}

		res, err := srv.CreateUser(context.Background(), req)

		assert.Error(t, err)
		assert.Equal(t, codes.AlreadyExists, status.Code(err))

		assert.Nil(t, res)
	})

}
