package helpers

import (
	"fmt"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	strToHash := "super duper password"

	hashedString := HashAndSalt(strToHash)

	isEqual := CompareHashStr(hashedString, strToHash)
	fmt.Println(hashedString)
	if !isEqual {
		t.Errorf("Hashed string should not be equal")
	}
}

func TestHashAndSaltWrongPassword(t *testing.T) {
	strToHash := "super duper password"
	wrongPassword := "super duper wrong password"

	hashedString := HashAndSalt(strToHash)

	isEqual := CompareHashStr(hashedString, wrongPassword)

	if isEqual {
		t.Errorf("Hashed string should be equal")
	}
}
