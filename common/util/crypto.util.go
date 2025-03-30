package commonutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSHA256Hash(data []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))

	mac.Write(data)

	sha := hex.EncodeToString(mac.Sum(nil))

	return sha
}

func VerifySHA256Hash(data []byte, secret, hash string) (bool, error) {
	sha, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(data)

	return hmac.Equal(sha, mac.Sum(nil)), nil
}
