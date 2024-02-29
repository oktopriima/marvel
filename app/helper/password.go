package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) string {
	bytePass := []byte(password)
	pass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	return string(pass)
}
