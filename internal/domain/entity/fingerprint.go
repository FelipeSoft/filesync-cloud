package entity

import "github.com/FelipeSoft/filesync-cloud/internal/domain/vobj"

type Fingerprint struct {
	Key      string
	CpuID    vobj.CpuID
	MAC      vobj.MAC
	Hostname string
}

func NewFingerprint(Key string, CpuID vobj.CpuID, MAC vobj.MAC, Hostname string) *Fingerprint {
	return &Fingerprint{
		Key:      Key,
		CpuID:    CpuID,
		MAC:      MAC,
		Hostname: Hostname,
	}
}
