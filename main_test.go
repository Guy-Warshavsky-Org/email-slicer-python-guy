package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// Unit tests for parseEmail
// ---------------------------------------------------------------------------

func TestParseEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		valid    bool
		username string
		domain   string
	}{
		{
			name:     "valid email simple",
			input:    "user@example.com",
			valid:    true,
			username: "user",
			domain:   "example.com",
		},
		{
			name:     "whitespace trimming",
			input:    "  user@example.com  ",
			valid:    true,
			username: "user",
			domain:   "example.com",
		},
		{
			name:  "no at symbol",
			input: "userexample.com",
			valid: false,
		},
		{
			name:     "multiple at symbols",
			input:    "a@b@c",
			valid:    true,
			username: "a",
			domain:   "b@c",
		},
		{
			name:     "empty username",
			input:    "@example.com",
			valid:    true,
			username: "",
			domain:   "example.com",
		},
		{
			name:     "empty domain",
			input:    "user@",
			valid:    true,
			username: "user",
			domain:   "",
		},
		{
			name:  "whitespace only input",
			input: "   ",
			valid: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := parseEmail(tc.input)
			if result.Valid != tc.valid {
				t.Errorf("parseEmail(%q).Valid = %v, want %v", tc.input, result.Valid, tc.valid)
			}
			if result.Valid {
				if result.Username != tc.username {
					t.Errorf("parseEmail(%q).Username = %q, want %q", tc.input, result.Username, tc.username)
				}
				if result.Domain != tc.domain {
					t.Errorf("parseEmail(%q).Domain = %q, want %q", tc.input, result.Domain, tc.domain)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// E2E tests — build the binary and verify exact stdout output
// ---------------------------------------------------------------------------

// buildBinary compiles the binary into a temporary directory and returns
// the path to the executable.
func buildBinary(t *testing.T) string {
	t.Helper()

	tmpDir := t.TempDir()
	binName := "email-slicer"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	binPath := filepath.Join(tmpDir, binName)

	// Find the module root (directory containing go.mod).
	// When tests run, the working directory is the package directory,
	// which for package main in the root is the module root itself.
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working directory: %v", err)
	}

	cmd := exec.Command("go", "build", "-o", binPath, ".")
	cmd.Dir = wd
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}
	return binPath
}

func TestE2E_ValidEmail(t *testing.T) {
	bin := buildBinary(t)

	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader("user@example.com\n")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Fatalf("binary exited with error: %v", err)
	}

	expected := "Please enter your Email Id:\nYour username is:  user\nYour domain is:  example.com\n"
	if stdout.String() != expected {
		t.Errorf("stdout mismatch\ngot:\n%s\nwant:\n%s", stdout.String(), expected)
	}
}

func TestE2E_InvalidEmail(t *testing.T) {
	bin := buildBinary(t)

	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader("userexample.com\n")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Fatalf("binary exited with error: %v", err)
	}

	expected := "Please enter your Email Id:\nPlease enter a valid Email Id.\n"
	if stdout.String() != expected {
		t.Errorf("stdout mismatch\ngot:\n%s\nwant:\n%s", stdout.String(), expected)
	}
}

func TestE2E_MultipleAtSymbols(t *testing.T) {
	bin := buildBinary(t)

	cmd := exec.Command(bin)
	cmd.Stdin = strings.NewReader("a@b@c\n")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		t.Fatalf("binary exited with error: %v", err)
	}

	expected := "Please enter your Email Id:\nYour username is:  a\nYour domain is:  b@c\n"
	if stdout.String() != expected {
		t.Errorf("stdout mismatch\ngot:\n%s\nwant:\n%s", stdout.String(), expected)
	}
}
