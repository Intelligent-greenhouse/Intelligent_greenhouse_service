package service

import (
	//"context"
	user "intelligent-greenhouse-service/api/web/user"
	"intelligent-greenhouse-service/domain"
)

// AuthService .
type AuthService struct {
	user.UserHTTPServer

	uc *domain.AuthUsecase
}

// NewAuthService .
func NewAuthService(uc *domain.AuthUsecase) *AuthService {

	return &AuthService{uc: uc}
}
