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

func (j *JwtTokenManager) VerifyRSA(tokenString string, publicKey *rsa.PublicKey) (any, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error verifying token: %v", err)
	}

	if !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return parsedToken.Claims, nil
}
