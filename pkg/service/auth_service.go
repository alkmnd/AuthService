package service

import (
	"authService/pkg/repository "
	"authService/pkg/repository /models"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	key             = "higyuftydr"
	accessTokenTTL  = 24 * time.Hour
	refreshTokenTTL = 24 * time.Hour * 90
)

type accessTokenClaims struct {
	jwt.StandardClaims
	UserId    uuid.UUID `json:"user_id"`
	IpAddress string    `json:"ip_address"`
}

type AuthService struct {
	repo *repository.AuthPostgres
}

func NewAuthService(repo *repository.AuthPostgres) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(email string) (uuid.UUID, error) {
	id, err := s.repo.CreateUser(email)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *AuthService) GenerateAccessToken(userId uuid.UUID, ipAddress string) (string, string, error) {
	tokenId := uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &accessTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        tokenId,
		},
		userId,
		ipAddress,
	})

	signedString, err := token.SignedString([]byte(key))

	if err != nil {
		return "", "", err
	}

	return signedString, tokenId, nil
}

func (s *AuthService) GenerateRefreshToken(userId uuid.UUID, ipAddress string, jti string) (string, error) {
	refreshToken := uuid.New().String()
	encodedRefreshToken := base64.StdEncoding.EncodeToString([]byte(refreshToken))
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	createdAt := time.Now()
	expiresAt := createdAt.Add(refreshTokenTTL)
	err = s.repo.CreateToken(models.Token{
		UserId:    userId,
		TokenHash: string(hashedToken),
		IpAddress: ipAddress,
		Jti:       jti,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
		IsRevoked: false,
	})

	return encodedRefreshToken, nil
}
