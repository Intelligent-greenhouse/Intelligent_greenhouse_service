package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	api "intelligent-greenhouse-service/api/web/user"
	domain "intelligent-greenhouse-service/domain/user"
	service "intelligent-greenhouse-service/service/user"

	dao "intelligent-greenhouse-service/infra/dao/user"
)

func UserRegister() RegisterDomainHandler {
	return func(ctx context.Context, s *http.Server) {
		conf := Config(ctx)
		data := Data(ctx)
		logger := Logger(ctx)
		repo := dao.NewUserDao(data, conf, logger)
		usecase := domain.NewUserDomain(repo, logger)
		srv := service.NewUserService(usecase)
		api.RegisterUserHTTPServer(s, srv)
	}
}
