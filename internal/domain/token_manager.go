package domain

import "crypto/rsa"

type TokenManager interface {
	AssignRSA256(sub string) (string, error)
	VerifyRSA256(tokenString string, publicKey *rsa.PublicKey) (any, error)
}
