package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 8
	password := []byte(p)

	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePassword(hash, password []byte) bool {
	tempHash, tempPassword := []byte(hash), []byte(password)

	err := bcrypt.CompareHashAndPassword(tempHash, tempPassword)

	return err == nil
}
