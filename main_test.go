package main

import (
	"bytes"
	"fmt"
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
			name:     "valid email with subdomain",
			input:    "user@mail.example.org",
			wantUser: "user",
			wantDom:  "mail.example.org",
			wantOK:   true,
			wantErr:  "",
		},
		{
			name:     "valid email with leading and trailing whitespace",
			input:    "  user@domain.com  ",
			wantUser: "user",
			wantDom:  "domain.com",
			wantOK:   true,
			wantErr:  "",
		},
		{
			name:     "valid email with tabs and spaces",
			input:    "\t user@domain.com \t",
			wantUser: "user",
			wantDom:  "domain.com",
			wantOK:   true,
			wantErr:  "",
		},
		{
			name:     "multiple at symbols splits on first",
			input:    "user@sub@domain.com",
			wantUser: "user",
			wantDom:  "sub@domain.com",
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
			name:    "only whitespace",
			input:   "   ",
			wantOK:  false,
			wantErr: "Please enter a valid Email Id.",
		},
		{
			name:    "at sign only",
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

func TestOutputFormat(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		wantOut  string
	}{
		{
			name:  "valid email output format",
			email: "avimax37@gmail.com",
			wantOut: "Your username is:  avimax37\nYour domain is:  gmail.com\n",
		},
		{
			name:  "invalid email output format",
			email: "not-an-email",
			wantOut: "Please enter a valid Email Id.\n",
		},
		{
			name:  "multiple at symbols output format",
			email: "user@sub@domain.com",
			wantOut: "Your username is:  user\nYour domain is:  sub@domain.com\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			username, domain, valid, errMsg := sliceEmail(tt.email)
			if valid {
				fmt.Fprintf(&buf, "Your username is:  %s\n", username)
				fmt.Fprintf(&buf, "Your domain is:  %s\n", domain)
			} else {
				fmt.Fprintln(&buf, errMsg)
			}

			got := buf.String()
			if got != tt.wantOut {
				t.Fatalf("output for %q:\ngot:  %q\nwant: %q", tt.email, got, tt.wantOut)
			}
		})
	}
}
