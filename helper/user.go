package helper

import (
	"errors"
	"time"

	"github.com/Adilfarooque/Footgo/config"
	"github.com/Adilfarooque/Footgo/utils/models"
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
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
	//It creates new jwt token using
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAccessToken(user models.UserDetailsResponse) (string, error) {
	// It calculates the expiration time for the access token
	expirationTime := time.Now().Add(15 * time.Minute)
	// This function generates a JWT token using the provided details
	tokenString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	// If the token is successfully generated, it returns the token string.
	return tokenString, nil
}

func GenerateRefreshToken(user models.UserDetailsResponse) (string, error) {
	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokenString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
