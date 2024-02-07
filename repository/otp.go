package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo/db"
	"github.com/Adilfarooque/Footgo/domain"
	"github.com/Adilfarooque/Footgo/utils/models"
	"gorm.io/gorm"
)

func FindUserByPhoneNumber(phone string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where(&domain.User{Phone: phone}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error) {
	var userDetails models.UserDetailsResponse
	if err := db.DB.Raw("SELECT * FROM users WHERE phone = ?", phone).Scan(&userDetails).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}
	return userDetails, nil
}
