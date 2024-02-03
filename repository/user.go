package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo/db"
	"github.com/Adilfarooque/Footgo/domain"
	"github.com/Adilfarooque/Footgo/utils/models"
	"gorm.io/gorm"
)

// checks if a user with the given email exists in the database.
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

// checks if a user with the given phone number exists in the database.
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

// handles user registration. It performs a raw SQL insert query and returns user details upon success.
func UserSignUp(user models.UserSignUp) (models.UserDetailsResponse, error) {
	var SignupDetail models.UserDetailsResponse
	err := db.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUES(?,?,?,?,?)RETURNING id,firstname,lastname,email,password,phone", user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&SignupDetail).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	return SignupDetail, nil
}

func FindUserByEmail(user models.LoginDetail) (models.UserLoginResponse, error) {
	var userDetails models.UserLoginResponse
	err := db.DB.Raw("SELECT * FROM users WHERE email=? and blocked=false", user.Email).Scan(&userDetails).Error
	if err != nil {
		return models.UserLoginResponse{}, errors.New("error checking user details")
	}
	return userDetails, nil
}

func AddAddress(userID int, address models.AddressInfo) error {
	err := db.DB.Exec("INSERT INTO addresses(user_id,name,house_name,street,city,state,pin)VALUES(?,?,?,?,?,?,?)", userID, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Error
	if err != nil {
		return errors.New("could not add address")
	}
	return nil
}

func GetAllAddress(userId int) ([]models.AddressInfoResponse, error) {
	var addressInfoResponse []models.AddressInfoResponse
	if err := db.DB.Raw("SELECT * FORM 	addresses WHERE user_id = ?", userId).Scan(&addressInfoResponse).Error; err != nil {
		return []models.AddressInfoResponse{}, err
	}
	return addressInfoResponse, nil
}

func UserDetails(userID int) (models.UsersProfileDetails, error) {
	var userDetails models.UsersProfileDetails
	err := db.DB.Raw("SELECT u.firstname,u.lastname,u.email,u.phone, FROM users u.id = ?", userID).Row().Scan(&userDetails.Firstname, &userDetails.Lastname, &userDetails.Email, &userDetails.Phone)
	if err != nil {
		return models.UsersProfileDetails{}, err
	}
	err = db.DB.Raw("SELECT referral_code FROM referrals WHERE user_id = ?", userID).Scan(&userDetails.RefferralCode).Error
	if err != nil {
		return models.UsersProfileDetails{}, err
	}
	return userDetails, nil
}

func CheckUserAvailabilityWithUserID(userID int) bool {
	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM users WHERE id = ?", userID).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func UpadateUserEmail(email string, userID int) error {
	err := db.DB.Exec("UPDATE users SET email = ? WHERE id = ?", email, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPhone(phone string, userID int) error {
	if err := db.DB.Exec("UPDATE users SET phone = ? WHERE id = ?", phone, userID).Error; err != nil {
		return err
	}
	return nil
}

func UpadatFirstName(name string, userID int) error {
	if err := db.DB.Exec("UPDATE users SET firstname = ? WHERE id = ?", name, userID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateLastName(name string, userID int) error {
	if err := db.DB.Exec("UPDATE users SET lastname = ? WHERE id = ?", name, userID).Error; err != nil {
		return err
	}
	return nil
}

func CheckAddressAvailabilityWithAddressID(addressID, userID int) bool {
	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM addresses WHERE id = ? AND user_id = ?", addressID, userID).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func UpdateName(name string, addressID int) error {

	err := db.DB.Exec("UPDATE addresses SET name = ? WHERE id = ?", name, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateHouseName(HouseName string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET house_name = ? WHERE id = ?", HouseName, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateStreet(street string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET state = ? WHERE id = ?", street, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCity(city string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET city = ? WHERE id = ?", city, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateState(state string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET state = ? WHERE id = ?", state, addressID).Error
	if err != nil {
		return err
	}
	return err
}

func UpdatePin(pin string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET pin = ? WHERE id = ?", pin, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateReferralEntry(userDetails models.UserDetailsResponse, userReferal string) error {
	err := db.DB.Exec("INSERT INTO referrals (user_id,referral_code,referral_amount)VALUES(?,?,?)", userDetails.Id, userReferal, 0).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserIdFromReferrals(ReferralCode string) (int, error) {
	var referredUserId int
	err := db.DB.Raw("SELECT user_id FROM referrals WHERE referral_code = ?", ReferralCode).Scan(&referredUserId).Error
	if err != nil {
		return 0, err
	}
	return referredUserId, nil
}

func UpdateReferralAmount(referralAmount float64, referredUserId int, currentUserID int) error {
	err := db.DB.Exec("UPDATE referrals SET referral_amount = ? , referred_user_id = ?", referralAmount, referredUserId, currentUserID).Error
	if err != nil {
		return err
	}
	//find the current amount in referred users referral table and add 100 with that
	err = db.DB.Exec("UPDATE referrals SET referral_amount = referral_amount + ? WHERE user_id = ? ", referralAmount, referredUserId).Error
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

func ExitWallet(userID int) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM wallets WHERE user_id = ? ", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func NewWallet(userID int, amount float64) error {
	err := db.DB.Exec("INSERT INTO wallets (user_id,amount)VALUES(?,?)", userID, amount).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateReferUserWallet(amount float64, userID int) error {
	err := db.DB.Exec("INSERT INTO wallets (user_id,amount)VALUES(?,?)", userID, amount).Error
	if err != nil {
		return err
	}
	return nil
}
