package trigger

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"intelligent-greenhouse-service/conf"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(config *conf.Trigger, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if config.Http.Network != "" {
		opts = append(opts, http.Network(config.Http.Network))
	}
	if config.Http.Addr != "" {
		opts = append(opts, http.Address(config.Http.Addr))
	}
	if config.Http.Timeout != nil {
		opts = append(opts, http.Timeout(config.Http.Timeout.AsDuration()))
	}
	server := http.NewServer(opts...)
	return server
}
