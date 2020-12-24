package common

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test upload directory", args{"UPLOADS_DIRECTORY_PATH", "uploads"},
			"uploads"},
		{"dummy test", args{"DUMMY_VARIABLE", "dummyValue"}, "dummyValue"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
