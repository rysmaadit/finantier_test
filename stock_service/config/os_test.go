package config

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	type args struct {
		key string
	}

	tests := []struct {
		name string
		mock func()
		args args
		want string
	}{
		{
			name: "given set environment value string, should return string",
			mock: func() {
				_ = os.Setenv("key", "value")
			},
			args: args{key: "key"},
			want: "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := GetString(tt.args.key); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		mock func()
		args args
		want int
	}{
		{
			name: "given set environment value int, should return int",
			mock: func() {
				_ = os.Setenv("id", "1")
			},
			args: args{key: "id"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := GetInt(tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
