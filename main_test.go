package main

import (
	"testing"
)

func TestSliceEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantUser string
		wantDom  string
		wantOK   bool
		wantErr  string
	}{
		{
			name:     "valid simple email",
			input:    "avimax37@gmail.com",
			wantUser: "avimax37",
			wantDom:  "gmail.com",
			wantOK:   true,
			wantErr:  "",
		},
		{
			name:     "valid with whitespace",
			input:    "  user@domain.com  ",
			wantUser: "user",
			wantDom:  "domain.com",
			wantOK:   true,
			wantErr:  "",
		},
		{
			name:     "valid with subdomain",
			input:    "john@mail.example.org",
			wantUser: "john",
			wantDom:  "mail.example.org",
			wantOK:   true,
			wantErr:  "",
		},
		{
			name:    "no at symbol",
			input:   "not-an-email",
			wantOK:  false,
			wantErr: "Please enter a valid Email Id.",
		},
		{
			name:    "empty string",
			input:   "",
			wantOK:  false,
			wantErr: "Please enter a valid Email Id.",
		},
		{
			name:    "whitespace only",
			input:   "   ",
			wantOK:  false,
			wantErr: "Please enter a valid Email Id.",
		},
		{
			name:    "at only",
			input:   "@",
			wantOK:  false,
			wantErr: "Please enter a valid Email Id.",
		},
		{
			name:    "missing domain",
			input:   "user@",
			wantOK:  false,
			wantErr: "Please enter a valid Email Id.",
		},
		{
			name:    "missing username",
			input:   "@domain.com",
			wantOK:  false,
			wantErr: "Please enter a valid Email Id.",
		},
		{
			name:     "multiple at symbols",
			input:    "user@sub@domain.com",
			wantUser: "user",
			wantDom:  "sub@domain.com",
			wantOK:   true,
			wantErr:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, dom, ok, errMsg := sliceEmail(tt.input)
			if user != tt.wantUser || dom != tt.wantDom || ok != tt.wantOK || errMsg != tt.wantErr {
				t.Fatalf("sliceEmail(%q) = (%q, %q, %v, %q), want (%q, %q, %v, %q)",
					tt.input, user, dom, ok, errMsg,
					tt.wantUser, tt.wantDom, tt.wantOK, tt.wantErr,
				)
			}
		})
	}
}
