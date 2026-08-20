package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortCodeUrl(originalURL string) string {
	hash := sha256.Sum256([]byte(originalURL))
	return base64.RawURLEncoding.EncodeToString(hash[:])[:8]
}
