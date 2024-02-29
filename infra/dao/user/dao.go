package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain/user"
	"intelligent-greenhouse-service/infra"
	jwt "intelligent-greenhouse-service/middleware"
)

type userDao struct {
	data *infra.Data
	log  *log.Helper
	conf *conf.Bootstrap
}

// Login 用户登录
func (r *userDao) Login(ctx context.Context, userName, passWord string) (id int32, err error) {
	// 获取数据库用户信息
	var userInfo User
	tx := r.data.Db.Where(&User{Username: userName, Password: passWord}).First(&userInfo)
	err = tx.Error
	if err != nil {
		return 0, errors.New(404, "", "user not found")
	}

	// 签署token验证
	token, err := jwt.CreateLoginToken(jwt.LoginToken{
		UserID:  userInfo.ID,
		IsAdmin: userInfo.IsAdmin,
	}, []byte(r.conf.Jwt.VerifyKey), r.conf.Jwt.LoginExpireTime)
	if err != nil {
		return 0, err
	}
	jwt.TokenInject(ctx, token)

	return userInfo.ID, nil
}

// NewUserDao .
func NewUserDao(data *infra.Data, c *conf.Bootstrap, logger log.Logger) user.UserRepo {
	return &userDao{
		data: data,
		log:  log.NewHelper(logger),
		conf: c,
	}
}
