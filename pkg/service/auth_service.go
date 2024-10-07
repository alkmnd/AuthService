package service

import (
	"authService/models"
	"authService/pkg/repository"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/smtp"
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
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &accessTokenClaims{
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
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(encodedRefreshToken), bcrypt.DefaultCost)
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
	})
	if err != nil {
		return "", err
	}

	return encodedRefreshToken, nil
}

func (s *AuthService) ParseAccessToken(accessToken string) (uuid.UUID, string, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &accessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(key), nil
	})
	if err != nil {
		return uuid.Nil, "", "", err
	}

	claims, ok := token.Claims.(*accessTokenClaims)
	if !ok {
		return uuid.Nil, "", "", errors.New("incorrect claims")
	}

	return claims.UserId, claims.IpAddress, claims.Id, nil
}

func (s *AuthService) IsRefreshValid(refreshToken string, userId uuid.UUID, accessId string, ipAddress string) bool {

	repoToken, err := s.repo.GetTokenInfo(userId)
	if err != nil {
		return false
	}
	println(repoToken.TokenHash)
	err = bcrypt.CompareHashAndPassword([]byte(repoToken.TokenHash), []byte(refreshToken))
	if err != nil {
		return false
	}
	if repoToken.IpAddress != ipAddress {
		fmt.Println(ipAddress)
		_ = s.SendWarning(userId)
		return true
	}
	if accessId != repoToken.Jti {
		return false
	}

	return true
}

func (s *AuthService) SendWarning(userId uuid.UUID) error {
	user, err := s.repo.GetUser(userId)
	from := "email@gmail.com"
	password := "qwerty"
	to := user.Email

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: warning: ip address was change"
	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}
	return nil

}

func (s *AuthService) UpdateAccessId(userId uuid.UUID, jti string) error {
	return s.repo.UpdateAccessId(userId, jti)
}
