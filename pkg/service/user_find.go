package service

import (
	"context"
	"errors"
	"fmt"
	"simpleService/external/dto"
	"simpleService/external/middleware"
	"simpleService/external/utils/hash"
	"simpleService/internal/config/errorcode"
	requestmodel "simpleService/pkg/model/request"
	"simpleService/pkg/repository"
)

func (s userImpl) Login(ctx context.Context, request requestmodel.UserLogin) (string, error) {
	var (
		repo = repository.User()
	)

	// find user
	loginData, err := repo.Find(dto.User{
		Username: request.Username,
	}, nil)
	if err != nil {
		return "", errors.New(errorcode.UserDoesNotExist)
	}

	// comparepassword
	if err := hash.ComparePassword(loginData.Password, request.Password); err != nil {
		fmt.Println(err)
		return "", errors.New("Invalid username or password")
	}

	accessToken, err := middleware.GenerateJWT(loginData)
	if err != nil {
		return "", errors.New(errorcode.FailedToGenerateAccessToken)
	}

	return accessToken, nil
}

func (s userImpl) Logout(ctx context.Context, accessToken string) error {
	if err := middleware.Revoke(accessToken); err != nil {
		return err
	}
	return nil
}
