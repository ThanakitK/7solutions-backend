package authorization

import "github.com/stretchr/testify/mock"

type MockAuthorization struct {
	mock.Mock
}

func NewAuthorizationMock() *MockAuthorization {
	return &MockAuthorization{}
}

func (m *MockAuthorization) GenerateToken(payload AppAuthorizationClaim) (token string, err error) {
	args := m.Called(payload)
	return args.String(0), args.Error(1)
}

func (m *MockAuthorization) ValidateToken(tokenString string, paserTo interface{}) (err error) {
	args := m.Called(tokenString, paserTo)
	return args.Error(0)
}
