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
		wantErr    bool
	}{
		{
			name:       "valid email",
			input:      "avimax37@gmail.com",
			wantUser:   "avimax37",
			wantDomain: "gmail.com",
			wantErr:    false,
		},
		{
			name:       "valid email with newline",
			input:      "avimax37@gmail.com\n",
			wantUser:   "avimax37",
			wantDomain: "gmail.com",
			wantErr:    false,
		},
		{
			name:    "missing @ sign",
			input:   "invalidemail",
			wantErr: true,
		},
		{
			name:    "missing @ with newline",
			input:   "invalidemail\n",
			wantErr: true,
		},
		{
			name:       "leading and trailing spaces",
			input:      "  user@domain.com  ",
			wantUser:   "user",
			wantDomain: "domain.com",
			wantErr:    false,
		},
		{
			name:       "multiple @ signs",
			input:      "user@sub@domain.com",
			wantUser:   "user",
			wantDomain: "sub@domain.com",
			wantErr:    false,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "only whitespace",
			input:   "   \t  ",
			wantErr: true,
		},
		{
			name:    "@ at start - empty username",
			input:   "@domain.com",
			wantErr: true,
		},
		{
			name:    "@ at end - empty domain",
			input:   "user@",
			wantErr: true,
		},
		{
			name:    "only @ sign",
			input:   "@",
			wantErr: true,
		},
		{
			name:       "email with subdomain",
			input:      "user@mail.example.com",
			wantUser:   "user",
			wantDomain: "mail.example.com",
			wantErr:    false,
		},
		{
			name:       "email with plus addressing",
			input:      "user+tag@example.com",
			wantUser:   "user+tag",
			wantDomain: "example.com",
			wantErr:    false,
		},
		{
			name:       "email with dots in username",
			input:      "first.last@example.com",
			wantUser:   "first.last",
			wantDomain: "example.com",
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, domain, err := sliceEmail(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("sliceEmail(%q): wantErr=%v, got err=%v", tt.input, tt.wantErr, err)
			}
			if !tt.wantErr {
				if user != tt.wantUser {
					t.Errorf("sliceEmail(%q): username = %q, want %q", tt.input, user, tt.wantUser)
				}
				if domain != tt.wantDomain {
					t.Errorf("sliceEmail(%q): domain = %q, want %q", tt.input, domain, tt.wantDomain)
				}
			}
		})
	}
}
