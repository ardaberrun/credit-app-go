package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 14);

	return string(hashedPwd), err;
}

func ComparePassword(hpwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hpwd), []byte(pwd));

	return err == nil;
}