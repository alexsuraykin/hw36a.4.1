package conf

import (
	"testing"
)

const (
	pathRSSConfig = "../../cmd/server/config.json"
	pathBDConfig  = "../../cmd/server/BD.json"
)

func TestNewBD(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Wrong filePath",
			args: args{
				filePath: "/wrong/wrong.json",
			},
			wantErr: true,
		},
		{
			name: "Right filePath",
			args: args{
				filePath: pathRSSConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBD(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && err == nil {
				t.Errorf("NewBD() got = %v, want not nil", got)
			}
		})
	}
}

func TestNewRSS(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Wrong filePath",
			args: args{
				filePath: "/wrong/wrong.json",
			},
			wantErr: true,
		},
		{
			name: "Right filePath",
			args: args{
				filePath: pathBDConfig,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRSS(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRSS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && err == nil {
				t.Errorf("NewBD() got = %v, want not nil", got)
			}
		})
	}
}
