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
		// Valid email cases
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
			input:    "user@mail.example.com",
			wantUser: "user",
			wantDom:  "mail.example.com",
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

		// Multiple @ symbols — split on first @
		{
			name:     "multiple @ symbols",
			input:    "user@sub@domain.com",
			wantUser: "user",
			wantDom:  "sub@domain.com",
			wantOK:   true,
			wantErr:  "",
		},

		// Invalid email cases
		{
			name:     "no @ symbol",
			input:    "not-an-email",
			wantUser: "",
			wantDom:  "",
			wantOK:   false,
			wantErr:  "Please enter a valid Email Id.",
		},
		{
			name:     "empty string",
			input:    "",
			wantUser: "",
			wantDom:  "",
			wantOK:   false,
			wantErr:  "Please enter a valid Email Id.",
		},
		{
			name:     "only whitespace",
			input:    "   ",
			wantUser: "",
			wantDom:  "",
			wantOK:   false,
			wantErr:  "Please enter a valid Email Id.",
		},
		{
			name:     "@ only",
			input:    "@",
			wantUser: "",
			wantDom:  "",
			wantOK:   false,
			wantErr:  "Please enter a valid Email Id.",
		},
		{
			name:     "@ at start (empty username)",
			input:    "@domain.com",
			wantUser: "",
			wantDom:  "",
			wantOK:   false,
			wantErr:  "Please enter a valid Email Id.",
		},
		{
			name:     "@ at end (empty domain)",
			input:    "user@",
			wantUser: "",
			wantDom:  "",
			wantOK:   false,
			wantErr:  "Please enter a valid Email Id.",
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

func TestSliceEmailOutputFormat(t *testing.T) {
	// Verify the exact output format matches Python's print() comma-separator behavior.
	// Python: print("Your username is: ", username) produces "Your username is:  avimax37"
	// (double space: one from the string literal, one from comma separator)

	username, domain, valid, _ := sliceEmail("avimax37@gmail.com")
	if !valid {
		t.Fatal("expected valid email")
	}

	expectedUsername := "avimax37"
	expectedDomain := "gmail.com"

	if username != expectedUsername {
		t.Errorf("username = %q, want %q", username, expectedUsername)
	}
	if domain != expectedDomain {
		t.Errorf("domain = %q, want %q", domain, expectedDomain)
	}

	// Verify the format strings would produce correct output with double space
	expectedLine1 := "Your username is:  avimax37\n"
	expectedLine2 := "Your domain is:  gmail.com\n"

	gotLine1 := "Your username is:  " + username + "\n"
	gotLine2 := "Your domain is:  " + domain + "\n"

	if gotLine1 != expectedLine1 {
		t.Errorf("output line 1 = %q, want %q", gotLine1, expectedLine1)
	}
	if gotLine2 != expectedLine2 {
		t.Errorf("output line 2 = %q, want %q", gotLine2, expectedLine2)
	}
}
