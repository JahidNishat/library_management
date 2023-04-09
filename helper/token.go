package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UserType  string
	UserId    string
	jwt.RegisteredClaims
}

var SecretKey = "DFB6C2407462BF22D2C659Ni23C1sh9E2C8FCatF11867A42316CBA203D54"

func GenerateAllTokens(email string, firstName string, lastName string, userType string, userId string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UserType:  userType,
		UserId:    userId,
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

func Validation(clientToken string)(claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		clientToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok{
		msg = fmt.Sprintf("Invalid token")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		msg = fmt.Sprintf("Token is Expired")
		msg = err.Error()
		return
	}

	return claims, msg
}