package service

import (
	"context"
	requestmodel "simpleService/pkg/model/request"
)

type UserInterface interface {
	Login(ctx context.Context, request requestmodel.UserLogin) (string, error)
	Logout(ctx context.Context, accessToken string) error
	Register(ctx context.Context, request requestmodel.UserRegister) (uint, error)
}

type userImpl struct{}

func User() UserInterface {
	return userImpl{}
}
