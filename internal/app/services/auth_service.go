package app_services

import (
	"api/internal/app/use_cases"
	infra_services "api/internal/infra/services"
	env_cfg "api/internal/infra/services/config"
	"errors"
	"fmt"
	"time"
)

type IAuthService interface {
	Login(email string, password string) ([]string, error)
}

type AuthService struct {
	customerService use_cases.ICustomerService
	tokenService    infra_services.ITokenService
	tokenCfg        env_cfg.ITokenConfig
}

var _ IAuthService = (*AuthService)(nil)

func NewAuthService(customerService use_cases.ICustomerService, tokenService infra_services.ITokenService) *AuthService {
	return &AuthService{
		customerService: customerService,
		tokenService:    tokenService,
		tokenCfg:        env_cfg.LoadConfig(),
	}
}

func (as *AuthService) Login(email string, password string) ([]string, error) {
	customer, err := as.customerService.GetCustomerByEmail(email)

	fmt.Println(err)

	if err != nil {
		return nil, errors.New("invalid email")
	}

	if customer.Password != password {
		return nil, errors.New("invalid password")
	}

	secret := as.tokenCfg.GetTokenSecret()
	tokenExp := as.tokenCfg.GetTokenExpiration()
	refreshExp := as.tokenCfg.GetRefreshTokenExpiration()

	token, err := as.tokenService.Generate(customer.ID, secret, time.Duration(tokenExp)*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := as.tokenService.Generate(customer.ID, secret, time.Duration(refreshExp)*time.Minute)
	if err != nil {
		return nil, err
	}

	return []string{token, refreshToken}, nil

}
