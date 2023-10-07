package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(str string) string {
	strByteArray := []byte(str)

	hash, err := bcrypt.GenerateFromPassword(strByteArray, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func CompareHashStr(hashedStr string, plainStr string) bool {
	byteHash := []byte(hashedStr)
	plainStrByteArray := []byte(plainStr)
	err := bcrypt.CompareHashAndPassword(byteHash, plainStrByteArray)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
