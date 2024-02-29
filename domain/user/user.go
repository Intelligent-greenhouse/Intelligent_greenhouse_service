package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	jwt "intelligent-greenhouse-service/middleware"
)

// UserRepo .
type UserRepo interface {
	Login(ctx context.Context, userName, passWord string) (id int32, err error)
	// ...
}

// UserDomain .
type UserDomain struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserDomain .
func NewUserDomain(repo UserRepo, logger log.Logger) *UserDomain {
	return &UserDomain{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserDomain) UserLogin(ctx context.Context, userName, passWord string) (int32, error) {
	id, err := uc.repo.Login(ctx, userName, passWord)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc UserDomain) UserAuthTest(ctx context.Context) (int32, error) {
	token, _ := jwt.FromLoginTokenContext(ctx)
	return token.UserID, nil
}
