package auth

import (
	"fmt"
	"testing"
)

func TestAuthUser(t *testing.T) {
	u:= UserEntity {
		Email: "calidad@gmail.com",
		Password: "Calidad2020",
	}

	isValid, err := ValidateUserAuth(u)
	fmt.Println("Is valid: ", isValid)
	fmt.Println("Error: ", err)
}

