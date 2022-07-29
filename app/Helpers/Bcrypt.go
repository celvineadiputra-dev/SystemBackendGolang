package Helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	pass := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
