package rmysql

import (
	"database/sql/driver"

	"github.com/FelipeSoft/filesync-cloud/internal/domain/entity"
)

type MySQLFingerprintRepository struct {
	conn *driver.Conn
}

func NewMySQLFingerprintRepository(conn *driver.Conn) *MySQLFingerprintRepository {
	return &MySQLFingerprintRepository{
		conn: conn,
	}
}

func (r *MySQLFingerprintRepository) Save(fingerprint *entity.Fingerprint) error {
	return nil
}

func (r *MySQLFingerprintRepository) GetFingerprintByAgentId(agentId string) (*entity.Fingerprint, error) {
	return nil, nil
}
