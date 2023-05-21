package repository

import (
	"simpleService/external/dto"
	"simpleService/internal/pg"

	"gorm.io/gorm"
)

func (userImpl) Find(user dto.User, db *gorm.DB) (*dto.User, error) {
	if db == nil {
		db = pg.GetDB()
	}

	result := db.Where("username = ?", user.Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
