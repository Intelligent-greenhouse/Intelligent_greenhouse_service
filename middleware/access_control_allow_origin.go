package jwt

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func AddHeader() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tr.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
				tr.ReplyHeader().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
				tr.ReplyHeader().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				tr.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
			}
			return handler(ctx, req)
		}
	}
}
