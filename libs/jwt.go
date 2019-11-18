package lib

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenExpiryDuration74H = time.Hour * 72
	tokenExpiryDuration24H = time.Hour * 24
)
const ErrorMessageExpire   = "Timing is everything"
const ErrorMessageBadToken = "That's not even a token"
type (
	transferClaims struct {
		Data []byte `json:"data"`
		jwt.StandardClaims
	}
)

func JwtUnpackStruct (tokenString *string, secret *[]byte) (data []byte, err error) {
	claims := transferClaims{}
	token, err := jwt.ParseWithClaims(*tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return *secret, nil
	})

	if token.Valid {
		return claims.Data, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return data, fmt.Errorf(ErrorMessageBadToken)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return data, fmt.Errorf(ErrorMessageExpire)
		} else {
			return data, fmt.Errorf("Couldn't handle this token: %v", err)
		}
	} else {
		return data, fmt.Errorf("Couldn't handle this token: %v", err)
	}
}

func JwtPackStruct (data interface{}, secret *[]byte, duration time.Duration) (tokenString string, err error) {
	arrByte, err := json.Marshal(data)

	if err != nil {
		return tokenString, fmt.Errorf("error convet data to arr byte : %v", err)
	}

	claims := transferClaims{
		arrByte,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(*secret)
}
