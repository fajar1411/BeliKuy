package scripts

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Bcript(y string) string {
	password := []byte(y)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)

}

func CheckPassword(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		log.Println("login compare", err.Error())
		return errors.New("password tidak sesuai ")
	}
	return nil
}
func ComparePassword(hashedPassword, password string) error {
	password = Bcript(password)
	if hashedPassword != password {
		return errors.New("password not match")
	}
	return nil
}
