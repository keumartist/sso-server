package mocks

import (
	auth "art-sso/internal/service/auth"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	auth.AuthService
	mock.Mock
}

func (m *MockAuthService) SignUpWithEmail(input auth.SignUpInput) (auth.Tokens, error) {
	args := m.Called(input)
	return args.Get(0).(auth.Tokens), args.Error(1)
}

func (m *MockAuthService) SignInWithEmail(input auth.SignInInput) (auth.Tokens, error) {
	args := m.Called(input)
	return args.Get(0).(auth.Tokens), args.Error(1)
}
