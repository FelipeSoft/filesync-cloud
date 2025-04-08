package service

import (
	"fmt"
	"log"

	"github.com/FelipeSoft/filesync-cloud/internal/application/dto"
	keyloader "github.com/FelipeSoft/filesync-cloud/internal/infrastructure/crypto"
	"github.com/FelipeSoft/filesync-cloud/internal/domain"
	"github.com/FelipeSoft/filesync-cloud/internal/domain/entity"
	"github.com/FelipeSoft/filesync-cloud/internal/domain/repository"
	"github.com/FelipeSoft/filesync-cloud/internal/domain/vobj"
)

type FingerprintService struct {
	tokenManager          domain.TokenManager
	fingerprintRepository repository.FingerprintRepository
}

func NewFingerprintService(tokenManager domain.TokenManager, fingerprintRepository repository.FingerprintRepository) *FingerprintService {
	return &FingerprintService{
		tokenManager:          tokenManager,
		fingerprintRepository: fingerprintRepository,
	}
}

func (s *FingerprintService) VerifyFingerprint(input dto.FingerprintRequest) (string, error) {
	cpuId, err := vobj.NewCpuID(input.CpuID)
	if err != nil {
		return "", err
	}

	mac, err := vobj.NewMAC(input.MAC)
	if err != nil {
		return "", err
	}

	fingerprint := entity.NewFingerprint(
		input.Key,
		*cpuId,
		*mac,
		input.Hostname,
	)

	if fingerprint.Key != "abc" {
		return "", fmt.Errorf("the installation key does not match")
	}
	return "token", nil
}

func (s *FingerprintService) Refresh(bearerToken string) (string, error) {
	publicKey, err := keyloader.LoadPublicKey()
	if err != nil {
		return "", err
	}
	payload, err := s.tokenManager.VerifyRSA256(bearerToken, publicKey)
	if err != nil {
		return "", err
	}
	log.Printf("Token Payload: %v", payload)
	return "", nil
}
