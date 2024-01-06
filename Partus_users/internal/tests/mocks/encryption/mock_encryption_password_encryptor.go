package encryption

import (
	"github.com/stretchr/testify/mock"
)

type MockPasswordEncryptor struct {
	mock.Mock
}

func (m *MockPasswordEncryptor) EncryptPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}
