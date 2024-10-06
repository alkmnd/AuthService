package service

import (
	"errors"
	"github.com/google/uuid"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(email string) (uuid.UUID, error) {
	return uuid.Nil, errors.New("")
}

func (s *AuthService) GenerateAccessToken(userId uuid.UUID) (string, error) {
	return "", errors.New("")
}

func (s *AuthService) GenerateRefreshToken(userId uuid.UUID) (string, error) {
	return "", errors.New("")
}
