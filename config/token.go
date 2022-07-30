package config

import (
	"github.com/golang-jwt/jwt/v4"
	"mvcGolang/app/Helpers"
	"strconv"
	"time"
)

type Token interface {
	GenerateToken(email string, userId int) (string, error)
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
