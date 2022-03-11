package kssh

import (
	"bufio"
	"io"

	"github.com/zeromicro/go-zero/core/logx"
)

//Cmd is in host exec cmd
func (ss *SSH) Cmd(host string, cmd string) string {
	session, err := ss.NewSession(host)
	if err != nil {
		logx.Errorf("[ssh][%s]Error create ssh session failed,%s", host, err)
	}
	defer session.Close()
	b, err := session.CombinedOutput(cmd)
	if err != nil {
		logx.Errorf("[ssh][%s]Error exec command failed: %s", host, err)
	}

	if b != nil {
		str := string(b)
		// str = strings.ReplaceAll(str, "\r\n", spilt)
		return str
	}

	return ""
}

// 异步执行命令
func (ss *SSH) CmdAsync(host string, cmd string) error {
	logx.Infof("[%s] %s", host, cmd)
	session, err := ss.NewSession(host)
	if err != nil {
		logx.Error("[ssh][%s]Error create ssh session failed,%s", host, err)
		return err
	}
	defer session.Close()
	stdout, err := session.StdoutPipe()
	if err != nil {
		logx.Errorf("[ssh][%s]Unable to request StdoutPipe(): %s", host, err)
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		logx.Errorf("[ssh][%s]Unable to request StderrPipe(): %s", host, err)
		return err
	}
	if err := session.Start(cmd); err != nil {
		logx.Errorf("[ssh][%s]Unable to execute command: %s", host, err)
		return err
	}
	doneout := make(chan bool, 1)
	doneerr := make(chan bool, 1)
	go func() {
		readPipe(host, stderr, true)
		doneerr <- true
	}()
	go func() {
		readPipe(host, stdout, false)
		doneout <- true
	}()
	<-doneerr
	<-doneout
	return session.Wait()
}

func readPipe(host string, pipe io.Reader, isErr bool) {
	r := bufio.NewReader(pipe)
	for {
		line, _, err := r.ReadLine()
		if line == nil {
			return
		} else if err != nil {
			logx.Infof("[%s] %s", host, line)
			logx.Errorf("[ssh] [%s] %s", host, err)
			return
		} else {
			if isErr {
				logx.Errorf("[%s] %s", host, line)
			} else {
				logx.Infof("[%s] %s", host, line)
			}
		}
	}
}
