package config

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"mvcGolang/app/Helpers"
	"strconv"
	"time"
)

type Token interface {
	GenerateToken(email string, userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

func GenerateToken(email string, userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = Helpers.Encrypt(strconv.Itoa(userId))
	claim["email"] = email
	claim["access_limit"] = time.Now().Add(time.Hour * 3).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, error := token.SignedString([]byte("T77H1nG25_0o0L07_000+"))
	if error != nil {
		return signedToken, error
	}

	return signedToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte("T77H1nG25_0o0L07_000+"), nil
	})

	if err != nil {
		return token, nil
	}

	return token, nil
}
