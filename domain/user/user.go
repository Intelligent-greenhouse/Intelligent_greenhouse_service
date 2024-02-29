package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/domain/device"
	jwt "intelligent-greenhouse-service/middleware"
)

// UserRepo .
type UserRepo interface {
	Login(ctx context.Context, userName, passWord string) (id int32, err error)
	// ...
}

// UserDomain .
type UserDomain struct {
	userRepo   UserRepo
	deviceRepo device.DeviceRepo
	log        *log.Helper
}

// NewUserDomain .
func NewUserDomain(userRepo UserRepo, deviceRepo device.DeviceRepo, logger log.Logger) *UserDomain {
	return &UserDomain{
		userRepo:   userRepo,
		deviceRepo: deviceRepo,
		log:        log.NewHelper(logger),
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
