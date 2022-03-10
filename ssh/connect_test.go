package ssh

import (
	"testing"
	"time"

	"golang.org/x/crypto/ssh"
)

func TestSSH_NewSession(t *testing.T) {
	type fields struct {
		User       string
		Password   string
		PkFile     string
		PkPassword string
		Timeout    time.Duration
	}
	type args struct {
		host string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ssh.Session
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				User:       "root",
				Password:   "Kn@999+.",
				PkFile:     "",
				PkPassword: "",
				Timeout:    time.Second * 10,
			},
			args: args{
				host: "10.67.15.212",
			},
			wantErr: false,
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
			got, err := ss.NewSession(tt.args.host)
			t.Logf("%v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("SSH.NewSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
