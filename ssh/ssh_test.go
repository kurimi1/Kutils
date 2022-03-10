package ssh

import (
	"reflect"
	"testing"
	"time"
)

func TestNewSSH(t *testing.T) {
	type args struct {
		user       string
		password   string
		pkFile     string
		pkPassword string
		timeout    time.Duration
	}
	tests := []struct {
		name string
		args args
		want *SSH
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				user:       "user",
				password:   "password",
				pkFile:     "pkFile",
				pkPassword: "pkPassword",
			},
			want: &SSH{
				User:       "user",
				Password:   "password",
				PkFile:     "pkFile",
				PkPassword: "pkPassword",
				Timeout:    DefaultTimeout,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSSH(tt.args.user, tt.args.password, tt.args.pkFile, tt.args.pkPassword, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSSH() = %v, want %v", got, tt.want)
			} else {
				t.Logf("NewSSH() = %v", got)
			}

		})
	}
}
