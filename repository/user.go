package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"gorm.io/gorm"
)

func CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}

func CheckUserExistsByPhone(phone string) (*domain.User, error) {
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

func UserSignUp(user models.UserSignUp) (models.UserDetailsResponse, error) {
	var SginupDetails models.UserDetailsResponse
	err := db.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUES(?,?,?,?,?)RETURNING id,firstname,lastname,email,password,phone", user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&SginupDetails).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	return SginupDetails, nil
}

func GetUserIdFromReferrals(ReferralCode string) (int, error) {

	var referredUserId int
	err := db.DB.Raw("SELECT user_id FROM referrals WHERE referral_code = ?", ReferralCode).Scan(&referredUserId).Error
	if err != nil {
		return 0, nil
	}

	return referredUserId, nil
}

func CreateReferralEntry(userDetails models.UserDetailsResponse, userReferral string) error {
	err := db.DB.Exec("INSERT INTO referrals (user_id,referral_code,referral_amount) VALUES (?,?,?)", userDetails.Id, userReferral, 0).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateReferralAmount(referralAmont float64, referredUserId int, currentUserID int) error {
	err := db.DB.Exec("UPDATE referrals SET referral_amount = ? , referred_user_id = ? ", referralAmont, referredUserId, currentUserID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateHistory(userID, orderID int, amount float64, reason string) error {
	err := db.DB.Exec("INSERT INTO wallet_histories (user_id,order_id,description,amount) VALUES (?,?,?,?)", userID, orderID, reason, amount).Error
	if err != nil {
		return err
	}
	return nil
}

func AmountInrefferals(userID int) (float64, error) {
	var a float64
	err := db.DB.Raw("SELECT referral_amount FROM referrals WHERE user_id = ?", userID).Scan(&a).Error
	if err != nil {
		return 0.0, err
	}
	return a, nil
}

func ExistWallet(userID int) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM wallets WHERE user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewWallet(userID int, amount float64) error {
	err := db.DB.Exec("INSERT INTO wallets (user_id,amount) VALUES(?,?) ", userID, amount).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateReferUserWallet(amount float64, userID int) error {
	err := db.DB.Exec("UPDATE wallets SET amount = amount + ? WHERE user_id = ?", amount, userID).Error
	if err != nil {
		return err
	}
	return nil
}

