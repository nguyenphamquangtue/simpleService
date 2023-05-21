package service

import (
	"context"
	"errors"
	"simpleService/external/dto"
	"simpleService/external/utils/hash"
	"simpleService/internal/config/errorcode"
	requestmodel "simpleService/pkg/model/request"
	"simpleService/pkg/repository"
)

func (s userImpl) Register(ctx context.Context, request requestmodel.UserRegister) (id uint, err error) {
	var (
		repo = repository.User()
	)
	// hashPassword
	hashedPassword, err := hash.HashPassword(request.Password)
	if err != nil {
		return 0, errors.New(errorcode.FailedToHashPassword)
	}

	id, err = repo.Insert(dto.User{
		Username: request.Username,
		Password: hashedPassword,
	}, nil)

	return id, nil
}
