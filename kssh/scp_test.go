package kssh

import (
	"testing"
	"time"
)

func TestSSH_Copy(t *testing.T) {
	type fields struct {
		User       string
		Password   string
		PkFile     string
		PkPassword string
		Timeout    time.Duration
	}
	type args struct {
		host           string
		localFilePath  string
		remoteFilePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				User:       "g",
				Password:   "Kn@999+.",
				PkFile:     "",
				PkPassword: "",
				Timeout:    time.Second * 10,
			},
			args: args{
				host:           "10.67.15.216",
				localFilePath:  "./ssh.go",
				remoteFilePath: "/home/g/ssh.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &SSH{
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				PkFile:     tt.fields.PkFile,
				PkPassword: tt.fields.PkPassword,
				Timeout:    tt.fields.Timeout,
			}
			ss.Copy(tt.args.host, tt.args.localFilePath, tt.args.remoteFilePath)
		})
	}
}
