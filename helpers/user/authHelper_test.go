package helper

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("xyz@abc.com", "John", "Cena", "fake")
	if err != nil {
		t.Errorf("failed with error -  %v", err)
	}

	if token == "" {
		t.Errorf("empty/invalid token is generated.")
	}
}

func TestVerifyToken(t *testing.T) {
	expectedEmail := "xyz@abc.com"
	token, _ := CreateToken(expectedEmail, "John", "Cena", "fake")
	claims, verifyErr := VerifyToken(token)
	if verifyErr != nil {
		t.Errorf("failed with error -  %v", verifyErr)
	}

	if claims.Email != expectedEmail {
		t.Errorf(
			"actualEmail - %s is not equals to expectedEmail - %s",
			claims.Email,
			expectedEmail,
		)
	}

}
