package main

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "success",
			args: args{"configuration.toml.example"},
			want: &Config{
				Server: Server{
					Address: "127.0.0.1",
					Port:    "8000",
				},
				Database: Database{
					User:    "superheroapi",
					Pass:    "superheroapi",
					Address: "127.0.0.1",
					Port:    "5432",
					Name:    "superheroapi",
				},
			},
			wantErr: false,
		},
		{
			name:    "error",
			args:    args{"configuration.toml.error"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
