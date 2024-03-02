package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	user "intelligent-greenhouse-service/api/web/user"
	domain_user "intelligent-greenhouse-service/domain/user"
	jwt "intelligent-greenhouse-service/middleware"
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
	err = s.uc.BindDeviceAndGreenhouse(ctx, req.DeviceCode, req.GreenHouseId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *UserService) GetAllUserList(ctx context.Context, req *user.Page) (rsp *user.UserList, err error) {
	if !s.uc.IsAdmin(ctx) {
		return nil, errors.New(403, "", "Forbidden")
	}
	list, err := s.uc.GetUserList(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	var userList []*user.UserInfo
	for _, u := range list {
		userList = append(userList, &user.UserInfo{UserId: u.ID, UserName: u.Username})
	}

	return &user.UserList{List: userList}, nil
}

func (s *UserService) GetUserGreenHorseList(ctx context.Context, req *user.UserId) (rsp *user.GreenHouseList, err error) {
	u, _ := jwt.FromLoginTokenContext(ctx)
	if u.UserID != req.UserId && !u.IsAdmin {
		return nil, errors.New(403, "", "Forbidden")
	}

	list, err := s.uc.GetUserGreenHouseList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var greenHouseList []*user.GreenHouse
	for _, g := range list {
		greenHouseList = append(greenHouseList, &user.GreenHouse{
			Id:   g.ID,
			Pos:  g.Pos,
			Size: g.Size,
			Des:  g.Des,
		})
	}

	return &user.GreenHouseList{List: greenHouseList}, nil
}

func (s *UserService) RegisterDevice(ctx context.Context, req *user.DeviceCode) (rsp *user.DeviceCode, err error) {
	if !s.uc.IsAdmin(ctx) {
		return nil, errors.New(403, "", "Forbidden")
	}

	device, err := s.uc.RegisterDevice(ctx, req.DeviceCode)
	if err != nil {
		return nil, err
	}

	return &user.DeviceCode{DeviceCode: device.DeviceId}, nil
}

func (s *UserService) AddUserByAdmin(ctx context.Context, request *user.AddUserRequest) (*user.UserId, error) {
	if !s.uc.IsAdmin(ctx) {
		return nil, errors.New(403, "", "Forbidden")
	}

	userId, err := s.uc.CreateNewUser(ctx, request.Username, request.Password, request.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &user.UserId{UserId: userId}, nil
}
