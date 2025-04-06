package keys

import (
	"crypto/rsa"
	"os"
)

func LoadPublicKey() *rsa.PublicKey {
	key, err := os.ReadFile("./../keys/public.pem")
}