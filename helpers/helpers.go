package helpers

import (
	"context"
	"fmt"
	"unicode"

	emailVerifier "github.com/AfterShip/email-verifier"
	"github.com/ank809/JWT-Authentication-Go/database"
	"gopkg.in/mgo.v2/bson"
)

var (
	verifer = emailVerifier.NewVerifier()
)

func CheckPassword(password string) (bool, string) {
	if len(password) < 6 {
		return false, "Password length should be greater than 6"

	}
	containsUpper := false
	containsLower := false
	containsDigits := false
	containsSpecial := false

	for _, ch := range password {
		if unicode.IsUpper(ch) {
			containsUpper = true
		} else if unicode.IsLower(ch) {
			containsLower = true
		} else if unicode.IsDigit(ch) {
			containsDigits = true
		} else {
			containsSpecial = true
		}

	}
	if containsDigits && containsLower && containsSpecial && containsUpper {
		return true, "Password mets requuirements"
	} else {
		return false, "Password should have uppercase, lowercase digits and special characters"
	}

}

func CheckUsername(username string) (bool, string) {
	if len(username) < 3 {
		return false, "Username length should be minimum 4"
	}
	collection_name := "users"
	collection := database.OpenCollection(database.Client, collection_name)
	filter := bson.M{"username": username}
	docs, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
	}
	if docs > 0 {
		return false, "Username already exists choose new one"
	}
	return true, "Username is available"
}

func VerifyEMail(email string) (bool, string) {
	ret, err := verifer.Verify(email)
	if err != nil {
		return false, "verify email address failed"
	}
	if !ret.Syntax.Valid {
		return false, "email address syntax is invalid "
	}
	return true, "email is valid"
}
