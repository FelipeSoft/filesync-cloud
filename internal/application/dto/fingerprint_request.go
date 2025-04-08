package dto

import (
	_ "github.com/go-playground/validator/v10"
)

type FingerprintRequest struct {
	Key      string `json:"key" validate:"required"`
	CpuID    string `json:"cpu_id" validate:"required"`
	MAC      string `json:"mac_address" validate:"required"`
	Hostname string `json:"hostname" validate:"required"`
}