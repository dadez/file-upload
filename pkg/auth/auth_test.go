package auth

import (
	"os"
	"reflect"
	"testing"
)

func Test_validateFromFile(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test credentials from file", args{"user", "secret"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// override AUTH_FILES_PATH as test runs from pkg
			os.Setenv("AUTH_FILES_PATH", "../../test/")
			if got := validateFromFile(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("validateFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicAuth(t *testing.T) {
	type args struct {
		pass Handler
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BasicAuth(tt.args.pass); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
