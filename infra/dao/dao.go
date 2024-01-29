package dao

import (
	"context"

	"intelligent-greenhouse-service/domain"
	"intelligent-greenhouse-service/infra"

	"github.com/go-kratos/kratos/v2/log"
)

type authDao struct {
	data *infra.Data
	log  *log.Helper
}

// NewAuthDao .
func NewAuthDao(data *infra.Data, logger log.Logger) domain.AuthRepo {
	return &authDao{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authDao) Save(ctx context.Context, g *domain.Greeter) (*domain.Greeter, error) {
	return g, nil
}
