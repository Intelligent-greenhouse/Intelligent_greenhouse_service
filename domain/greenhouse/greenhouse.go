package greenhouse

import (
	"github.com/go-kratos/kratos/v2/log"
)

type GreenhouseRepo interface {
}

type GreenhouseDomain struct {
	greenhouseRepo GreenhouseRepo
	log            *log.Helper
}

func NewGreenhouseDomain(repo GreenhouseRepo, logger log.Logger) *GreenhouseDomain {
	return &GreenhouseDomain{
		greenhouseRepo: repo,
		log:            log.NewHelper(logger),
	}
}
