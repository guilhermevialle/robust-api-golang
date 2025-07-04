package env_config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	TokenSecret            string
	TokenExpiration        string
	RefreshTokenExpiration string
}

func LoadEnv() *EnvVariables {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Aviso: .env não carregado ou não encontrado (pode ser esperado em produção)")
	}

	config := &EnvVariables{
		TokenSecret:            os.Getenv("TOKEN_SECRET"),
		TokenExpiration:        os.Getenv("TOKEN_EXPIRATION"),
		RefreshTokenExpiration: os.Getenv("REFRESH_TOKEN_EXPIRATION"),
	}

	return config
}
