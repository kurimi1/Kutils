package ssh

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pkg/sftp"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/ssh"
)

const oneKBByte = 1024
const oneMBByte = 1024 * 1024

func (ss *SSH) Copy(host, localFilePath, remoteFilePath string) {
	sftpClient, err := ss.sftpConnect(host)
	if err != nil {
		logx.Errorf("[ssh][%s]scpCopy: %s", host, err)
	}
	defer sftpClient.Close()

	srcFile, err := os.Open(localFilePath)
	if err != nil {
		logx.Errorf("[ssh][%s]scpCopy: %s", host, err)
	}
	defer srcFile.Close()

	dstFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		logx.Errorf("[ssh][%s]scpCopy: %s", host, err)
	}
	defer dstFile.Close()
	
	buf := make([]byte, 100*oneMBByte) //100mb
	total := 0
	unit := ""
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		length, _ := dstFile.Write(buf[0:n])
		isKb := length/oneMBByte < 1
		speed := 0
		if isKb {
			total += length
			unit = "KB"
			speed = length / oneKBByte
		} else {
			total += length
			unit = "MB"
			speed = length / oneMBByte
		}
		totalLength, totalUnit := toSizeFromInt(total)
		logx.Infof("[ssh][%s]transfer total size is: %.2f%s ;speed is %d%s", host, totalLength, totalUnit, speed, unit)
	}
}

func toSizeFromInt(length int) (float64, string) {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(length)/oneMBByte), 64)
	if isMb := length/oneMBByte > 1; isMb {
		return value, "MB"
	}
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(length)/oneKBByte), 64)
	return value, "KB"
}

func (ss *SSH) sftpConnect(host string) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = ss.authMethod(ss.Password, ss.PkFile, ss.PkPassword)

	clientConfig = &ssh.ClientConfig{
		User:    ss.User,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		},
	}

	// connet to ssh
	addr = ss.addrReformat(host)

	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}

	return sftpClient, nil
}
