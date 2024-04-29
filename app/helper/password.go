package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) string {
	bytePass := []byte(password)
	pass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	return string(pass)
}

func CheckPassword(password string, hash string) bool {
	bytePass := []byte(password)
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		return false
	}
	return true
}
