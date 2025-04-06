package keys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func LoadPublicKey() (*rsa.PublicKey, error) {
	data, err := os.ReadFile("./keys/public.pem")
	if err != nil {
		return nil, fmt.Errorf("error during loading public key: %v", err)
	}	
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid PEM block inside PUBLIC KEY")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}