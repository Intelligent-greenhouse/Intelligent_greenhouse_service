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
	id, err := s.uc.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &user.LoginReply{
		UserId: id,
	}, nil
}
