package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenManager struct {
	signingMethodRSA jwt.SigningMethodRSA
}

func NewJwtTokenManager(signingMethodRSA jwt.SigningMethodRSA) *JwtTokenManager {
	return &JwtTokenManager{
		signingMethodRSA: signingMethodRSA,
	}
}

func (j *JwtTokenManager) AssignRSA(sub string) (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", fmt.Errorf("error during generation of private key: %v", err)
	}
	token := jwt.NewWithClaims(&j.signingMethodRSA, jwt.MapClaims{
		"sub": sub,
		"exp": jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
	})
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error during token signing: %v", err)
	}
	return signedToken, nil
}

func (j *JwtTokenManager) VerifyRSA(token string) (any, error) {
	decoded, err := j.signingMethodRSA.Verify(token, )
	return nil, nil
}