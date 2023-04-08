package helper

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct{
	Email string
	FirstName string
	LastName string
	UserType string
	UserId string
	jwt.RegisteredClaims
}

var SecretKey = "DFB6C2407462BF22D2C659Ni23C1sh9E2C8FCatF11867A42316CBA203D54"

func GenerateAllTokens(email string, firstName string, lastName string, userType string, userId string)(signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		UserType: userType,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(360 * time.Minute)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SecretKey))
	if err != nil {
		log.Fatalln("Error in Tokenization...... %v", err)
		return
	}

	return token, refreshToken, err
}