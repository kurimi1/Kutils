package ssh

import (
	"time"
)

type SSH struct {
	User       string
	Password   string
	PkFile     string
	PkPassword string
	Timeout    time.Duration
}

// DefaultTimeout is the 1 minute for ssh.
const DefaultTimeout = time.Minute

// NewSSH returns a new SSH struct.
func NewSSH(user, password, pkFile, pkPassword string, timeout time.Duration) *SSH {
	// if timeou is 0, use default timeout.
	if timeout == 0 {
		timeout = DefaultTimeout
	}
	return &SSH{
		User:       user,
		Password:   password,
		PkFile:     pkFile,
		PkPassword: pkPassword,
		Timeout:    timeout,
	}
}
