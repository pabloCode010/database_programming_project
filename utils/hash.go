package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// generate a hash using SHA256
func HashSHA256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}
