package kssh

import (
	"fmt"
	"strings"
	"time"

	"github.com/kurimi1/Kutils/kfile"
	"golang.org/x/crypto/ssh"
)

// privateKeyMethod returns AuthMethod that uses the pk.
func (ss *SSH) privateKeyMethod(pkFile, pkPassword string) (am ssh.AuthMethod, err error) {
	pkData, err := kfile.ReadFile(pkFile)
	if err != nil {
		return nil, err
	}
	var pk ssh.Signer
	if pkPassword == "" {
		pk, err = ssh.ParsePrivateKey(pkData)
		if err != nil {
			return nil, err
		}
	} else {
		bufPwd := []byte(pkPassword)
		pk, err = ssh.ParsePrivateKeyWithPassphrase(pkData, bufPwd)
		if err != nil {
			return nil, err
		}
	}

	return ssh.PublicKeys(pk), nil
}

// passwordMethod returns an AuthMethod using the given password.
func (ss *SSH) passwordMethod(passwd string) ssh.AuthMethod {
	return ssh.Password(passwd)
}

// authMethods returns a list of authentication methods.
func (ss *SSH) authMethod(passwd, pkFile, pkPasswd string) (auth []ssh.AuthMethod) {
	// if pkfile is not empty
	if kfile.FileExist(pkFile) {
		am, err := ss.privateKeyMethod(pkFile, pkPasswd)
		if err == nil {
			auth = append(auth, am)
		}
	}
	// if password is not empty
	if passwd != "" {
		auth = append(auth, ss.passwordMethod(passwd))
	}

	return auth
}

// newClient returns a new SSH client.
func (ss *SSH) connect(host string) (*ssh.Client, error) {
	auth := ss.authMethod(ss.Password, ss.PkFile, ss.PkPassword)
	config := ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}
	DefaultTimeout := time.Minute
	if ss.Timeout == 0 {
		ss.Timeout = DefaultTimeout
	}
	clientConfig := &ssh.ClientConfig{
		User:            ss.User,
		Auth:            auth,
		Timeout:         ss.Timeout,
		Config:          config,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := ss.addrReformat(host)

	return ssh.Dial("tcp", addr, clientConfig)
}

// ssh port
func (ss *SSH) addrReformat(host string) string {
	if !strings.Contains(host, ":") {
		host = fmt.Sprintf("%s:22", host)
	}

	return host
}

// NewSession returns a new session.
func (ss *SSH) NewSession(host string) (*ssh.Session, error) {
	client, err := ss.connect(host)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}

	return session, nil
}
