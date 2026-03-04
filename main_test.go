package main

import (
	"testing"
)

func TestSliceEmail(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantUser   string
		wantDomain string
		wantValid  bool
	}{
		{
			name:       "standard valid email",
			input:      "avimax37@gmail.com",
			wantUser:   "avimax37",
			wantDomain: "gmail.com",
			wantValid:  true,
		},
		{
			name:       "no @ present",
			input:      "invalidemail",
			wantUser:   "",
			wantDomain: "",
			wantValid:  false,
		},
		{
			name:       "leading and trailing spaces",
			input:      "  user@example.com  ",
			wantUser:   "user",
			wantDomain: "example.com",
			wantValid:  true,
		},
		{
			name:       "multiple @ characters",
			input:      "user@sub@example.com",
			wantUser:   "user",
			wantDomain: "sub@example.com",
			wantValid:  true,
		},
		{
			name:       "empty input",
			input:      "",
			wantUser:   "",
			wantDomain: "",
			wantValid:  false,
		},
		{
			name:       "whitespace only",
			input:      "   ",
			wantUser:   "",
			wantDomain: "",
			wantValid:  false,
		},
		{
			name:       "input with trailing newline",
			input:      "test@domain.org\n",
			wantUser:   "test",
			wantDomain: "domain.org",
			wantValid:  true,
		},
		{
			name:       "@ at the start",
			input:      "@domain.com",
			wantUser:   "",
			wantDomain: "domain.com",
			wantValid:  true,
		},
		{
			name:       "@ at the end",
			input:      "user@",
			wantUser:   "user",
			wantDomain: "",
			wantValid:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, domain, valid := sliceEmail(tt.input)
			if user != tt.wantUser || domain != tt.wantDomain || valid != tt.wantValid {
				t.Fatalf("sliceEmail(%q) = (%q, %q, %v), want (%q, %q, %v)",
					tt.input, user, domain, valid, tt.wantUser, tt.wantDomain, tt.wantValid)
			}
		})
	}
}
