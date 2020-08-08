package bd

import "golang.org/x/crypto/bcrypt"

/* function to encrypt password*/
func EncryptPassword(pass string) (string, error) {
	cost := 8 // password security
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
