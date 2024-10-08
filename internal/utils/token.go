package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomToken generates a secure random token of a specified length.
func GenerateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
