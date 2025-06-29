package utils

import "golang.org/x/crypto/bcrypt"

func Bcryp_Encryption(str string) (result string, err error) {
	res, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", err
	}
	result = string(res)
	return result, nil
}

func Bcryp_Compare(str1 string, str2 string) (result bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(str1), []byte(str2)); err != nil {
		return false
	}
	return true
}
