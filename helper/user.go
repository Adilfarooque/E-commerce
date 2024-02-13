package helper

import (
	"errors"
	"time"

	"github.com/Adilfarooque/Footgo_Ecommerce/config"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUserClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func PasswordHash(password string) (string, error) {
	//It generates a bcrypt hash from the password using
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	//If there is an error during the hashing process, it returns an empty string along with a custom error message ("Internal server error")
	if err != nil {
		return "", errors.New("internal server error")
	}
	//Otherwise, it coverts the hashed password to string and returns it along with a nil error
	hash := string(hashPassword)
	return hash, nil
}

func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {
	cfg, _ := config.LoadConfig()
	claims := &AuthUserClaims{
		Id:    userID,
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(cfg.KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAccessToken(user models.UserDetailsResponse) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenrateRefreshToken(user models.UserDetailsResponse)(string,error){
	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokenString ,err := GenerateTokenUsers(user.Id,user.Email,expirationTime)
	if err != nil{
		return "",err
	}
	return tokenString,nil
}

