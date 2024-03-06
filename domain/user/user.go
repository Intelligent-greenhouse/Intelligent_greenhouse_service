package user

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/domain/device"
	"intelligent-greenhouse-service/domain/greenhouse"
	jwt "intelligent-greenhouse-service/middleware"
	"intelligent-greenhouse-service/model"
)

// UserRepo .
type UserRepo interface {
	Login(ctx context.Context, userName, passWord string) (id int32, err error)
	GetUserInfoByUserId(ctx context.Context, userId int32) (*model.User, error)
	GetUserInfoList(ctx context.Context, page, size int32) ([]*model.User, error)
	CreateUser(ctx context.Context, name, psw string, isAdmin bool) (int32, error)
}

// UserDomain .
type UserDomain struct {
	userRepo       UserRepo
	deviceRepo     device.DeviceRepo
	greenhouseRepo greenhouse.GreenhouseRepo
	log            *log.Helper
}

// NewUserDomain .
func NewUserDomain(userRepo UserRepo, deviceRepo device.DeviceRepo, greenhouseRepo greenhouse.GreenhouseRepo, logger log.Logger) *UserDomain {
	return &UserDomain{
		userRepo:       userRepo,
		deviceRepo:     deviceRepo,
		greenhouseRepo: greenhouseRepo,
		log:            log.NewHelper(logger),
	}
}

func (uc *UserDomain) UserLogin(ctx context.Context, userName, passWord string) (int32, error) {
	id, err := uc.userRepo.Login(ctx, userName, passWord)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc UserDomain) UserAuthTest(ctx context.Context) (int32, error) {
	token, _ := jwt.FromLoginTokenContext(ctx)
	return token.UserID, nil
}

func (uc UserDomain) IsAdmin(ctx context.Context) bool {
	token, _ := jwt.FromLoginTokenContext(ctx)
	return token.IsAdmin
}

func (uc UserDomain) GetUserInfo(ctx context.Context, userId int32) (*model.User, error) {
	return uc.userRepo.GetUserInfoByUserId(ctx, userId)
}

func (uc UserDomain) RegisterDevice(ctx context.Context, deviceCode string) (*model.Device, error) {
	return uc.deviceRepo.CreateDeviceInfo(ctx, deviceCode)
}

func (uc UserDomain) BindDeviceAndGreenhouse(ctx context.Context, deviceId, greenhouseId int32) error {

	// 查找设备是否存在
	deviceInfo, err := uc.deviceRepo.GetDeviceById(ctx, deviceId)
	if err != nil {
		return err
	}
	// 绑定设备必须在设备未激活的情况下
	if deviceInfo.IsActivation {
		return errors.New("device has active")
	}

	// 查找大棚是否存在
	if _, err = uc.greenhouseRepo.GetGreenhouseInfoById(ctx, greenhouseId); err != nil {
		return err
	}

	userInfo, _ := jwt.FromLoginTokenContext(ctx)
	// 绑定数据
	return uc.greenhouseRepo.BandGreenhouseAndDevice(ctx, deviceId, greenhouseId, userInfo.UserID)
}

func (uc UserDomain) GetUserList(ctx context.Context, page, size int32) ([]*model.User, error) {
	return uc.userRepo.GetUserInfoList(ctx, page, size)
}

func (uc UserDomain) GetUserGreenHouseList(ctx context.Context, userId int32) ([]*model.Greenhouse, error) {
	return uc.greenhouseRepo.GetGreenhouseListByUserId(ctx, userId)
}

func (uc UserDomain) CreateNewUser(ctx context.Context, name, psw string, isAdmin bool) (int32, error) {
	return uc.userRepo.CreateUser(ctx, name, psw, isAdmin)
}
