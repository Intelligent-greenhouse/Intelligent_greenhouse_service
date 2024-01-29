package service

import (
	//"context"

	//v1 "cypunsource-auth/api/auth/web/v1"
	"intelligent-greenhouse-service/domain"
	//"google.golang.org/protobuf/types/known/emptypb"
)

// AuthService .
type AuthService struct {
	//v1.UnimplementedAuthServiceServer

	uc *domain.AuthUsecase
}

// NewAuthService .
func NewAuthService(uc *domain.AuthUsecase) *AuthService {

	return &AuthService{uc: uc}
}

//func (s *AuthService) GetPublicKey(context.Context, *emptypb.Empty) (*v1.GetPublicKeyReply, error) {
//	return nil, nil
//}
//func (s *AuthService) UserAuth(context.Context, *v1.UserAuthRequest) (*v1.UserAuthReply, error) {
//	return nil, nil
//}
//func (s *AuthService) RefreshAuth(context.Context, *emptypb.Empty) (*v1.SessionReply, error) {
//	return nil, nil
//}
//func (s *AuthService) CreateAccount(context.Context, *v1.CreateAccountRequest) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *AuthService) VerifyEmail(context.Context, *v1.VerifyEmailRequest) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *AuthService) VerifyNationalID(context.Context, *v1.VerifyNationalIDRequest) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *AuthService) CloseUserAuth(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
//	return nil, nil
//}
//func (s *AuthService) ChangePassword(context.Context, *v1.ChangePasswordRequest) (*v1.UserAuthReply, error) {
//	return nil, nil
//}
