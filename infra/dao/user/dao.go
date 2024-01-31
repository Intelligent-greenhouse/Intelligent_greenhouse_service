package user

import (
	"context"

	"intelligent-greenhouse-service/domain/user"
	"intelligent-greenhouse-service/infra"

	"github.com/go-kratos/kratos/v2/log"
)

type userDao struct {
	data *infra.Data
	log  *log.Helper
}

// NewUserDao .
func NewUserDao(data *infra.Data, logger log.Logger) user.UserRepo {
	return &userDao{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userDao) Save(ctx context.Context, g *user.Greeter) (*user.Greeter, error) {
	return g, nil
}
