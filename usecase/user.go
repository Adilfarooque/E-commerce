package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Adilfarooque/Footgo/config"
	"github.com/Adilfarooque/Footgo/helper"
	"github.com/Adilfarooque/Footgo/repository"
	"github.com/Adilfarooque/Footgo/utils/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func UsersSignUp(user models.UserSignUp) (*models.TokenUser, error) {
	//Check if user alerady exists by email
	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New("user with this email is already exists")
	}
	//Check if user already exists by phone
	phone, err := repository.CheckUserExistsByPhone(user.Phone)
	fmt.Println(phone, nil)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if phone != nil {
		return &models.TokenUser{}, errors.New("user with this phone is already exists")
	}
	//Hash the password
	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return &models.TokenUser{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword
	//Signup the  user
	userData, err := repository.UserSignUp(user)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not add the user")
	}
	//Create referral code for the user and send in details  of referred id of user if it exists
	id := uuid.New().ID()
	//converts it to a string
	str := strconv.Itoa(int(id))
	//takes the first 8 charachter
	userReferral := str[:8]
	//It creates a referral entry in the repository using the generated referral code and user data
	err = repository.CreateReferralEntry(userData, userReferral)
	if err != nil {
		return &models.TokenUser{}, err
	}
	if user.RefferralCode != "" {
		//First check whether if a user with that referralCode exist
		referredUserId, err := repository.GetUserIdFromReferrals(user.RefferralCode)
		if err != nil {
			return &models.TokenUser{}, err
		}
		if referredUserId != 0 {
			referralAmount := 150
			err := repository.UpdateReferralAmount(float64(referralAmount), referredUserId, userData.Id)
			if err != nil {
				return &models.TokenUser{}, err
			}
			referreason := "Amount crediterd for used referral code"
			err = repository.UpdateHistory(userData.Id, 0, float64(referralAmount), referreason)
			if err != nil {
				return &models.TokenUser{}, err
			}
			amount, err := repository.AmountInrefferals(userData.Id)
			if err != nil {
				return &models.TokenUser{}, err
			}
			walletExist, err := repository.ExitWallet(userData.Id)
			if err != nil {
				return &models.TokenUser{}, err
			}
			if !walletExist {
				err = repository.NewWallet(userData.Id, amount)
				if err != nil {
					return &models.TokenUser{}, err
				}
			}
			err = repository.UpdateReferUserWallet(amount, referredUserId)
			if err != nil {
				return &models.TokenUser{}, err
			}
			reason := "Amount credited for refer a new person"
			err = repository.UpdateHistory(referredUserId, 0, amount, reason)
			if err != nil {
				return &models.TokenUser{}, err
			}
		}
	}
	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create access token due to error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create refresh token due to error")
	}
	return &models.TokenUser{
		Users:        userData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func UserLogin(user models.LoginDetail) (*models.TokenUser, error) {
	email, err := repository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return &models.TokenUser{}, errors.New("email doesn't exists")
	}
	userdetails, err := repository.FindUserByEmail(user)
	if err != nil {
		return &models.TokenUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdetails.Password), []byte(user.Password))
	if err != nil {
		return &models.TokenUser{}, errors.New("password not matching")
	}
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userdetails)
	if err != nil {
		return &models.TokenUser{}, err
	}
	accessToken, err := helper.GenerateRefreshToken(userDetails)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userDetails)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create refreshtoken due to internal error")
	}
	return &models.TokenUser{
		Users:        userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func ForgotPasswordSend(phone string) error {
	cfg, _ := config.LoadConfig()
	ok := repository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exists")
	}
	helper.TwilioSetup(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	_, err := helper.TwilioSendOTP(phone, cfg.SERVICESSID)
	if err != nil {
		return errors.New("error occured while generating OTP")
	}
	return nil
}
