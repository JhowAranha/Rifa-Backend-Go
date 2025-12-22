package hash

import "golang.org/x/crypto/bcrypt"

func CreateNewHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	check := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return check == nil
}
