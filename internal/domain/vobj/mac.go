package vobj

import (
	"fmt"
	"regexp"
)

type MAC struct {
	token string
}

func NewMAC(token string) (*MAC, error) {
	mac := &MAC{
		token: token,
	}
	isValid := mac.validate()
	if isValid {
		return nil, fmt.Errorf("invalid mac address provided")
	}
	return mac, nil
}

func (c *MAC) validate() bool {
	// Possible formats
	// 00:1A:2B:3C:4D:5E
	// 00-1A-2B-3C-4D-5E
	// 001A2B3C4D5E
	// 001A.2B3C.4D5E
	macPattern := `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$|^([0-9A-Fa-f]{4}\.){2}([0-9A-Fa-f]{4})$|^([0-9A-Fa-f]{12})$`
	match, _ := regexp.MatchString(macPattern, c.token)
	return match
}

func (c *MAC) ToString() string {
	return c.token
}
