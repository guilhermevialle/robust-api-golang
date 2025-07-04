package infra_services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ITokenService interface {
	Generate(payload any, secretKey string, expTime time.Duration) (string, error)
	Validate(token string, secretKey string) (bool, error)
}

type TokenService struct{}

var _ ITokenService = (*TokenService)(nil)

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (ts *TokenService) Generate(payload any, secretKey string, expTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(expTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func (ts *TokenService) Validate(tokenString string, secretKey string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
