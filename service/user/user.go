package user

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	user "intelligent-greenhouse-service/api/web/user"
	domain_user "intelligent-greenhouse-service/domain/user"
)

type UserService struct {
	uc *domain_user.UserDomain
}

func NewUserService(uc *domain_user.UserDomain) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) Login(ctx context.Context, req *user.LoginRequest) (rsp *user.LoginReply, err error) {
	id, err := s.uc.UserLogin(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &user.LoginReply{
		UserId: id,
	}, nil
}

func (s *UserService) UserAuth(ctx context.Context, req *emptypb.Empty) (rsp *user.LoginReply, err error) {
	//token, ok := jwt.FromLoginTokenContext(ctx)
	//if !ok {
	//    return nil, errors.New(403, "", "TokenNotFound")
	//}

	return &user.LoginReply{UserId: 1}, nil
}
