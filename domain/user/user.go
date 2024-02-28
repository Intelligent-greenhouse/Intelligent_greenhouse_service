package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo .
type UserRepo interface {
	Login(ctx context.Context, userName, passWord string) (int32, error)
	// ...
}

// UserDomain .
type UserDomain struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserDomain .
func NewUserDomain(repo UserRepo, logger log.Logger) *UserDomain {
	return &UserDomain{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserDomain) Login(ctx context.Context, userName, passWord string) (int32, error) {
	id, err := uc.repo.Login(ctx, userName, passWord)
	return id, err
}
