package commonutil

import (
	"fmt"
	"testing"
)

func TestSHA256Hash(t *testing.T) {
	data := []byte("data")
	secret := "secret"

	hash := GenerateSHA256Hash(data, secret)

	fmt.Println(VerifySHA256Hash(data, secret, hash))
}
