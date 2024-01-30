package service

import (
	"context"

	user "intelligent-greenhouse-service/api/web/user"
)

func (s *AuthService) Login(context.Context, *user.LoginRequest) (*user.LoginReply, error) {
	return &user.LoginReply{User: &user.UserInfo{Username: "aaa"}}, nil
}
