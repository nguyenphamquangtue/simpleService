package repository

import (
	"simpleService/external/dto"
	"simpleService/internal/pg"

	"gorm.io/gorm"
)

func (userImpl) Insert(user dto.User, db *gorm.DB) (uint, error) {
	if db == nil {
		db = pg.GetDB()
	}

	result := db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}
