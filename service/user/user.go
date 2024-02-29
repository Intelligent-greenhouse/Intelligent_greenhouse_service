package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
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
	test, err := s.uc.UserAuthTest(ctx)

	return &user.LoginReply{UserId: test}, nil
}

func (s *UserService) BindDevice(ctx context.Context, req *user.BindDeviceInfo) (rsp *emptypb.Empty, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) GetAllUserList(ctx context.Context, req *emptypb.Empty) (rsp *user.UserList, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) GetUserGreenHorseList(ctx context.Context, req *user.UserId) (rsp *user.GreenHouseList, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *UserService) RegisterDevice(ctx context.Context, req *user.DeviceCode) (rsp *emptypb.Empty, err error) {
	if !s.uc.IsAdmin(ctx) {
		return nil, errors.New(403, "", "Forbidden")
	}
	panic("implement me")
}
