package user

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"intelligent-greenhouse-service/conf"
	"intelligent-greenhouse-service/domain/user"
	"intelligent-greenhouse-service/infra"
	jwt "intelligent-greenhouse-service/middleware"
	"intelligent-greenhouse-service/model"
)

type userDao struct {
	data *infra.Data
	log  *log.Helper
	conf *conf.Bootstrap
}

func (r *userDao) CreateUser(ctx context.Context, name, psw string, isAdmin bool) (int32, error) {
	var userInfo *model.User
	err := r.data.Db.Where(&model.User{Username: name}).First(&userInfo).Error
	if err != nil {
		return 0, err
	}

	if userInfo != nil {
		return 0, errors.New(409, "", "username has exist")
	}

	newUser := &model.User{Username: name, Password: psw, IsAdmin: isAdmin}
	if err = r.data.Db.Create(&newUser).Error; err != nil {
		return 0, err
	}

	return newUser.ID, nil
}

func (r *userDao) GetUserInfoList(ctx context.Context, page, size int32) ([]*model.User, error) {
	var userList []*model.User
	if err := r.data.Db.Limit(int(size)).Offset(int((page - 1) * size)).Find(&userList).Error; err != nil {
		return nil, err
	}

	return userList, nil
}

func (r *userDao) GetUserInfoByUserId(ctx context.Context, userId int32) (*model.User, error) {
	var u model.User
	tx := r.data.Db.First(&u, userId)
	err := tx.Error
	if err != nil {
		return nil, errors.New(404, "", "user not found")
	}

	return &u, nil
}

// Login 用户登录
func (r *userDao) Login(ctx context.Context, userName, passWord string) (id int32, err error) {
	// 获取数据库用户信息
	var userInfo model.User
	tx := r.data.Db.Where(&model.User{Username: userName, Password: passWord}).First(&userInfo)
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

	// 将token放入cookie中
	tr, _ := transport.FromServerContext(ctx)

	tr.ReplyHeader().Set("Authorization", token)

	return userInfo.ID, nil
}

var instance *userDao

// NewUserDao .
func NewUserDao(data *infra.Data, c *conf.Bootstrap, logger log.Logger) user.UserRepo {
	if instance == nil {
		instance = &userDao{
			data: data,
			log:  log.NewHelper(logger),
			conf: c,
		}
	}
	return instance
}

func GetUserDaoInstance() user.UserRepo {
	return instance
}
