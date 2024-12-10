package password

import "golang.org/x/crypto/bcrypt"

// Hash returns the hash of the given password.
func Hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// Compare returns true if the password and hash match.
func Compare(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
