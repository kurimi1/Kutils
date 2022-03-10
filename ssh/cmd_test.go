package ssh

import (
	"testing"
	"time"
)

func TestSSH_Cmd(t *testing.T) {
	type fields struct {
		User       string
		Password   string
		PkFile     string
		PkPassword string
		Timeout    time.Duration
	}
	type args struct {
		host string
		cmd  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				User:       "c",
				Password:   "Kn@999+.",
				PkFile:     "",
				PkPassword: "",
				Timeout:    time.Second * 10,
			},
			args: args{
				host: "10.67.15.212",
				cmd:  "ls",
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
			if got := ss.Cmd(tt.args.host, tt.args.cmd); got != tt.want {
				t.Errorf("SSH.Cmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
