package service

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

const (
	signingKey = "higyuftydr"
	tokenTTL   = 24 * time.Hour
)

type accessTokenClaims struct {
	jwt.StandardClaims
	UserId    uuid.UUID `json:"user_id"`
	IpAddress string    `json:"ip_address"`
}

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(email string) (uuid.UUID, error) {
	return uuid.Nil, errors.New("")
}

func (s *AuthService) GenerateAccessToken(userId uuid.UUID, ipAddress string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
		ipAddress,
	})

	signedString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (s *AuthService) GenerateRefreshToken(userId uuid.UUID, ipAddress string, jti string) (string, error) {
	refreshToken := uuid.New().String()
	encodedRefreshToken := base64.StdEncoding.EncodeToString([]byte(refreshToken))

	return encodedRefreshToken, nil
}
