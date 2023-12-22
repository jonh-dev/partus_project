package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
UserTest é uma struct que representa um teste de criação de usuário. Ela possui os seguintes atributos:

  - name: nome do teste
  - request: requisição de criação de usuário
  - wantErr: booleano que indica se o teste espera um erro ou não
  - expectedError: mensagem de erro esperada
*/
type UserTest struct {
	Name          string
	Request       *api.CreateUserRequest
	WantErr       bool
	ExpectedError string
}

/*
RunTest é um método que executa um teste de criação de usuário. Ele recebe um ponteiro para um objeto UserTest e um ponteiro para um objeto testing.T. O método é responsável por configurar o mock do repositório de usuários e executar o teste de criação de usuário. O método é chamado pelo método RunTests.

  - mockRepo: mock do repositório de usuários
  - srv: serviço de usuários
*/
func (ut *UserTest) RunTest(t *testing.T) {
	mockRepo := new(MockUserRepository)
	srv := services.NewUserService(mockRepo)

	t.Run(ut.Name, func(t *testing.T) {
		setupMockRepository(mockRepo, ut)
		runCreateUserTest(srv, ut, t)
	})
}

/*
setupMockRepository é um método que configura o mock do repositório de usuários. Ele recebe um ponteiro para um objeto MockUserRepository e um ponteiro para um objeto UserTest. O método é responsável por configurar o mock do repositório de usuários de acordo com o teste de criação de usuário. O método é chamado pelo método RunTest.

  - mockRepo: mock do repositório de usuários
  - ut: teste de criação de usuário
*/
func setupMockRepository(mockRepo *MockUserRepository, ut *UserTest) {
	setupGetUserByEmailMock(mockRepo, ut)
	setupCreateUserMock(mockRepo, ut)
}

/*
setupGetUserByEmailMock é um método que configura o mock do método GetUserByEmail do repositório de usuários. Ele recebe um ponteiro para um objeto MockUserRepository e um ponteiro para um objeto UserTest. O método é responsável por configurar o mock do método GetUserByEmail do repositório de usuários de acordo com o teste de criação de usuário. O método é chamado pelo método setupMockRepository.

  - mockRepo: mock do repositório de usuários
  - ut: teste de criação de usuário
*/
func setupGetUserByEmailMock(mockRepo *MockUserRepository, ut *UserTest) {
	if ut.Name != "Email Already Exists" {
		mockRepo.On("GetUserByEmail", mock.Anything, ut.Request.PersonalInfo.Email).Return(nil, nil)
		return
	}
	mockRepo.On("GetUserByEmail", mock.Anything, ut.Request.PersonalInfo.Email).Return(&api.User{}, nil)
}

/*
setupCreateUserMock é um método que configura o mock do método CreateUser do repositório de usuários. Ele recebe um ponteiro para um objeto MockUserRepository e um ponteiro para um objeto UserTest. O método é responsável por configurar o mock do método CreateUser do repositório de usuários de acordo com o teste de criação de usuário. O método é chamado pelo método setupMockRepository.

  - mockRepo: mock do repositório de usuários
  - ut: teste de criação de usuário
*/
func setupCreateUserMock(mockRepo *MockUserRepository, ut *UserTest) {
	mockRepo.On("CreateUser", mock.Anything, ut.Request.PersonalInfo, ut.Request.AccountInfo).Return(&api.User{
		PersonalInfo: ut.Request.PersonalInfo,
		AccountInfo:  ut.Request.AccountInfo,
	}, nil)
}

/*
runCreateUserTest é um método que executa o teste de criação de usuário. Ele recebe um ponteiro para um objeto UserService, um ponteiro para um objeto UserTest e um ponteiro para um objeto testing.T. O método é responsável por executar o teste de criação de usuário de acordo com o teste de criação de usuário. O método é chamado pelo método RunTest.

  - srv: serviço de usuários
  - ut: teste de criação de usuário
  - t: objeto de teste
*/
func runCreateUserTest(srv services.UserService, ut *UserTest, t *testing.T) {
	resp, err := srv.CreateUser(context.Background(), ut.Request)
	if !ut.WantErr {
		assert.NoError(t, err)
		assert.Equal(t, ut.Request.PersonalInfo, resp.User.PersonalInfo)
		assert.Equal(t, ut.Request.AccountInfo, resp.User.AccountInfo)
		return
	}
	assert.Error(t, err)
	assertCorrectErrorMessage(t, err, ut)
}

/*
assertCorrectErrorMessage é um método que verifica se a mensagem de erro retornada pelo teste de criação de usuário está correta. Ele recebe um ponteiro para um objeto testing.T, um objeto de erro e um ponteiro para um objeto UserTest. O método é responsável por verificar se a mensagem de erro retornada pelo teste de criação de usuário está correta de acordo com o teste de criação de usuário. O método é chamado pelo método runCreateUserTest.

  - t: objeto de teste
  - err: objeto de erro
  - ut: teste de criação de usuário
*/
func assertCorrectErrorMessage(t *testing.T, err error, ut *UserTest) {
	if ut.Name != "Email Already Exists" {
		assert.EqualError(t, err, fmt.Sprintf("rpc error: code = InvalidArgument desc = %s", ut.ExpectedError))
		return
	}
	assert.EqualError(t, err, fmt.Sprintf("rpc error: code = AlreadyExists desc = %s", ut.ExpectedError))
}
