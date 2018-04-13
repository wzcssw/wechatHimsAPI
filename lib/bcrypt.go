package lib

import "golang.org/x/crypto/bcrypt"

// CheckPasswordHash Bcrypt验证
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
