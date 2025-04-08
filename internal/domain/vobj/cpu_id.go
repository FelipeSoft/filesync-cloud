package vobj

import (
	"fmt"
	"regexp"
)

type CpuID struct {
	token string
}

func NewCpuID(token string) (*CpuID, error) {
	cpuId := &CpuID{
		token: token,
	}
	isValid := cpuId.validate()
	if isValid {
		return nil, fmt.Errorf("invalid cpu id provided")
	}
	return cpuId, nil
}

func (c *CpuID) validate() bool {
	matchLinux, _ := regexp.MatchString(`^(AuthenticAMD|GenuineIntel|ARM|CPU).*`, c.token)
	matchWindows, _ := regexp.MatchString(`^[A-Fa-f0-9]{8,16}$`, c.token)
	return matchLinux || matchWindows
}

func (c *CpuID) ToString() string {
	return c.token
}
