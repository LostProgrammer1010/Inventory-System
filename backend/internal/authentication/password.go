package authentication

import "golang.org/x/crypto/bcrypt"

// Turns a plain text password into a hashed password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

// Checks the incoming password against the encrypted password
func VerifePassword(encryptedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
