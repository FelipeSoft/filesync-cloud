package repository

import (
	"github.com/FelipeSoft/filesync-cloud/internal/domain/entity"
)

type FingerprintRepository interface {
	Save(fingerprint *entity.Fingerprint) error
	GetFingerprintByAgentId(agentId string) (*entity.Fingerprint, error)
}
