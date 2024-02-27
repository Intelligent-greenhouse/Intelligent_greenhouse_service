package user

import (
	"context"
	domain_user "intelligent-greenhouse-service/domain/user"

	user "intelligent-greenhouse-service/api/web/user"
)

type UserService struct {
	uc *domain_user.UserDomain
}

func NewUserService(uc *domain_user.UserDomain) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Login(ctx context.Context, req *user.LoginRequest) (rsp *user.LoginReply, err error) {
	//TODO implement me
	return &user.LoginReply{User: &user.UserInfo{Username: req.Password}}, nil
}
