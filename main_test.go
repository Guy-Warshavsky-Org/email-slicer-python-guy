package main

import (
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

// buildBinary builds the email-slicer binary for integration tests
// and returns the path to the built binary.
func buildBinary(t *testing.T) string {
	t.Helper()

	// Build binary in a temp directory
	dir := t.TempDir()
	binary := filepath.Join(dir, "email-slicer")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}

	cmd := exec.Command("go", "build", "-o", binary, ".")
	cmd.Dir = "."
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build binary: %v\n%s", err, out)
	}
	return binary
}

func TestCLI(t *testing.T) {
	binary := buildBinary(t)

	tests := []struct {
		name   string
		input  string
		want   string
	}{
		{
			name:  "valid email",
			input: "avimax37@gmail.com\n",
			want:  "Please enter your Email Id:\nYour username is:  avimax37\nYour domain is:  gmail.com\n",
		},
		{
			name:  "invalid email no @",
			input: "invalidemail\n",
			want:  "Please enter your Email Id:\nPlease enter a valid Email Id.\n",
		},
		{
			name:  "multiple @ characters",
			input: "user@sub@example.com\n",
			want:  "Please enter your Email Id:\nYour username is:  user\nYour domain is:  sub@example.com\n",
		},
		{
			name:  "email with leading and trailing spaces",
			input: "  user@example.com  \n",
			want:  "Please enter your Email Id:\nYour username is:  user\nYour domain is:  example.com\n",
		},
		{
			name:  "empty input",
			input: "\n",
			want:  "Please enter your Email Id:\nPlease enter a valid Email Id.\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(binary)
			cmd.Stdin = strings.NewReader(tt.input)

			out, err := cmd.Output()
			if err != nil {
				t.Fatalf("command failed: %v", err)
			}

			got := string(out)
			if got != tt.want {
				t.Fatalf("output mismatch:\ngot:  %q\nwant: %q", got, tt.want)
			}
		})
	}
}
