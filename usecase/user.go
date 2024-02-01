package usecase

import (
	"errors"
	"fmt"

	"github.com/Adilfarooque/Footgo/repository"
	"github.com/Adilfarooque/Footgo/utils/models"
)

func UsersSigUp(user models.UserSignUp) (*models.TokenUser, error) {
	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New("user with this email is already exists")
	}

	phone, err := repository.CheckUserExistsByPhone(user.Phone)
	fmt.Println(phone, nil)
	if err != nil{
		return &models.TokenUser{},errors.New("error with server")
	}
	if phone != nil{
		return &models.TokenUser{},errors.New("user with this phone is already exists")
	}
	
}
