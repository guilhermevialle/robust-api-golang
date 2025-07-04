package env_cfg

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ITokenConfig interface {
	GetTokenSecret() string
	GetTokenExpiration() int
	GetRefreshTokenExpiration() int
}

type TokenConfig struct {
	tokenSecret            string
	tokenExpiration        int
	refreshTokenExpiration int
}

func LoadConfig() *TokenConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tokenSecret := os.Getenv("TOKEN_SECRET")

	tokenExpirationStr := os.Getenv("TOKEN_EXPIRATION")
	tokenExpiration, err := strconv.Atoi(tokenExpirationStr)
	if err != nil {
		log.Fatalf("Invalid TOKEN_EXPIRATION value: %v", err)
	}

	refreshExpirationStr := os.Getenv("REFRESH_TOKEN_EXPIRATION")
	refreshExpiration, err := strconv.Atoi(refreshExpirationStr)
	if err != nil {
		log.Fatalf("Invalid REFRESH_TOKEN_EXPIRATION value: %v", err)
	}

	return &TokenConfig{
		tokenSecret:            tokenSecret,
		tokenExpiration:        tokenExpiration,
		refreshTokenExpiration: refreshExpiration,
	}
}

func (t *TokenConfig) GetTokenSecret() string {
	return t.tokenSecret
}

func (t *TokenConfig) GetTokenExpiration() int {
	return t.tokenExpiration
}

func (t *TokenConfig) GetRefreshTokenExpiration() int {
	return t.refreshTokenExpiration
}
