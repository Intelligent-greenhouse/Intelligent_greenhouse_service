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

func (r *userDao) Login(ctx context.Context, userName, passWord string) (id int32, err error) {
	var userInfo User
	tx := r.data.Db.Where(&User{Username: userName, Password: passWord}).First(&userInfo)
	err = tx.Error
	if err != nil {
		return 0, err
	}
	if tx.RowsAffected == 0 {
		return 0, nil
	}

	return userInfo.ID, nil
}

// NewUserDao .
func NewUserDao(data *infra.Data, logger log.Logger) user.UserRepo {
	return &userDao{
		data: data,
		log:  log.NewHelper(logger),
	}
}
