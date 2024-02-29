package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/protobuf/ptypes/duration"
	"time"
)

type (
	// loginTokenKey context中登录token的键
	loginTokenKey struct{}
)

// LoginToken 登录token
type LoginToken struct {
	UserID  int32 `json:"user_id"`  // 用户id
	IsAdmin bool  `json:"is_admin"` // 是否是管理员账号
	jwt.RegisteredClaims
}

// FromLoginTokenContext 返回存储在context中的登录token，如果有
func FromLoginTokenContext(ctx context.Context) (loginToken *LoginToken, ok bool) {
	loginToken, ok = ctx.Value(loginTokenKey{}).(*LoginToken)
	fmt.Println(ctx)
	return
}

// LoginAuthMiddleware 登录鉴权中间件
func LoginAuthMiddleware(secret []byte) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("Authorization")
				if token == "" {
					return nil, errors.New("TokenNotFound")
				}
				loginToken, success := ValidateLoginToken(token, secret)
				if !success {
					return nil, errors.New("TokenValidateError")
				}
				ctx = context.WithValue(ctx, loginTokenKey{}, loginToken)
			}
			return handler(ctx, req)
		}
	}
}

// CreateLoginToken 创建登录token
func CreateLoginToken(claims LoginToken, secret []byte, expirationTime *duration.Duration) (signedToken string, err error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expirationTime.AsDuration()))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secret)
	return
}

// ValidateLoginToken 验证登录token
func ValidateLoginToken(signedToken string, secret []byte) (claims *LoginToken, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &LoginToken{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return secret, nil
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := token.Claims.(*LoginToken)
	if ok && token.Valid {
		success = true
		return
	}
	return
}
