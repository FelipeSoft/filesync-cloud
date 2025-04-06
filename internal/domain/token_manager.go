package domain

import "crypto/rsa"

type TokenManager interface {
	AssignRSA(sub string) (string, error)
	VerifyRSA(tokenString string, publicKey *rsa.PublicKey)
}
