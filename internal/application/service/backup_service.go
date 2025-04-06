package service

import (
	"fmt"

	"github.com/FelipeSoft/filesync-cloud/internal/domain"
)

type BackupService struct {
	tokenManager domain.TokenManager
}

func NewBackupService(tokenManager domain.TokenManager) *BackupService {
	return &BackupService{
		tokenManager: tokenManager,
	}
}

func (s *BackupService) VerifyInstallationKey(key string) error {
	if key != "abc" {
		return fmt.Errorf("the installation key does not match")
	}
	return nil
}