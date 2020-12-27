package validate

import (
	"testing"
)

func TestValidatePEM(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Validate test certificate", args{"../../test/github.crt.pem"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ValidatePEM(tt.args.f)
		})
	}
}
