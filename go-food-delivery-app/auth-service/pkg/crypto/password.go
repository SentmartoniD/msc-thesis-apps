package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// IsPasswordValid checks if passwords match
func IsPasswordValid(currentPassword, typedPassword string) bool {
	splits := strings.Split(currentPassword, ".")
	salt := splits[0]
	typedPassword = fmt.Sprintf("%s.%s", salt, typedPassword)
	typedPassword = SHA256(typedPassword)
	if len(splits) < 2 || typedPassword != splits[1] {
		return false
	}

	return true
}

// RandomPassword generates random password
func RandomPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	var password []byte
	for i := 0; i < length; i++ {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password = append(password, charset[charIndex.Int64()])
	}
	return string(password), nil
}
