package utils

import (
	"crowFunding/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"github.com/badoux/checkmail"
)

func Validate(action string, user *model.User) error {
	switch strings.ToLower(action) {
	case "signup":
		if user.Name == "" {
			return errors.New("Required Name!")
		}
		if user.Occupation == "" {
			return errors.New("Required Occupation!")
		}
		if user.Email == "" {
			return errors.New("Required Email!")
		}
		if user.Password == "" {
			return errors.New("Required Password!")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email address")
		}
		return nil
	default:
		if user.Name == "" {
			return errors.New("Required Name!")
		}
		if user.Occupation == "" {
			return errors.New("Required Occupation!")
		}
		if user.Email == "" {
			return errors.New("Required Email!")
		}
		if user.Password == "" {
			return errors.New("Required Password!")
		}
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return errors.New("Invalid email address")
		}
		return nil
	}
}

// VerifyPassword use for validasi password
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}