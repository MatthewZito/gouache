package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a given plaintext password into an encrypted, salted hash.
func HashPassword(password string) (string, error) {
	// GenerateFromPassword auto generates salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a given plaintext password to a given encrypted, salted hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
