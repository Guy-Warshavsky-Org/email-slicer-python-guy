package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// --- Unit tests for SliceEmail ---

func TestSliceEmail_Valid(t *testing.T) {
	username, domain, ok := SliceEmail("avimax37@gmail.com")
	if !ok {
		t.Fatal("expected ok=true for valid email")
	}
	if username != "avimax37" {
		t.Errorf("username = %q, want %q", username, "avimax37")
	}
	if domain != "gmail.com" {
		t.Errorf("domain = %q, want %q", domain, "gmail.com")
	}
}

func TestSliceEmail_NoAt(t *testing.T) {
	_, _, ok := SliceEmail("invalidemail")
	if ok {
		t.Fatal("expected ok=false for input without @")
	}
}

func TestSliceEmail_MultipleAt(t *testing.T) {
	username, domain, ok := SliceEmail("user@sub@domain.com")
	if !ok {
		t.Fatal("expected ok=true for input with multiple @")
	}
	if username != "user" {
		t.Errorf("username = %q, want %q", username, "user")
	}
	if domain != "sub@domain.com" {
		t.Errorf("domain = %q, want %q", domain, "sub@domain.com")
	}
}

func TestSliceEmail_Empty(t *testing.T) {
	_, _, ok := SliceEmail("")
	if ok {
		t.Fatal("expected ok=false for empty input")
	}
}

func TestSliceEmail_AtOnly(t *testing.T) {
	username, domain, ok := SliceEmail("@")
	if !ok {
		t.Fatal("expected ok=true for input that is just @")
	}
	if username != "" {
		t.Errorf("username = %q, want %q", username, "")
	}
	if domain != "" {
		t.Errorf("domain = %q, want %q", domain, "")
	}
}

func TestSliceEmail_AtStart(t *testing.T) {
	username, domain, ok := SliceEmail("@domain.com")
	if !ok {
		t.Fatal("expected ok=true for input starting with @")
	}
	if username != "" {
		t.Errorf("username = %q, want %q", username, "")
	}
	if domain != "domain.com" {
		t.Errorf("domain = %q, want %q", domain, "domain.com")
	}
}

func TestSliceEmail_AtEnd(t *testing.T) {
	username, domain, ok := SliceEmail("user@")
	if !ok {
		t.Fatal("expected ok=true for input ending with @")
	}
	if username != "user" {
		t.Errorf("username = %q, want %q", username, "user")
	}
	if domain != "" {
		t.Errorf("domain = %q, want %q", domain, "")
	}
}

// --- CLI integration tests ---

// captureMainOutput runs main() with the given input piped to stdin
// and captures whatever main() writes to stdout.
func captureMainOutput(input string) (string, error) {
	// Save original stdin/stdout
	origStdin := os.Stdin
	origStdout := os.Stdout
	defer func() {
		os.Stdin = origStdin
		os.Stdout = origStdout
	}()

	// Create a pipe for stdin
	stdinR, stdinW, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stdin = stdinR

	// Write input and close the write end
	_, err = stdinW.WriteString(input)
	if err != nil {
		stdinR.Close()
		stdinW.Close()
		return "", err
	}
	stdinW.Close()

	// Create a pipe for stdout
	stdoutR, stdoutW, err := os.Pipe()
	if err != nil {
		stdinR.Close()
		return "", err
	}
	os.Stdout = stdoutW

	// Run main
	main()

	// Close the write end so the reader can finish
	stdoutW.Close()

	// Read all captured output
	var buf bytes.Buffer
	_, err = io.Copy(&buf, stdoutR)
	stdoutR.Close()
	stdinR.Close()
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func TestMain_ValidInput_Output(t *testing.T) {
	output, err := captureMainOutput("avimax37@gmail.com\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 output lines, got %d: %q", len(lines), output)
	}

	if lines[0] != "Please enter your Email Id:" {
		t.Errorf("line 0 = %q, want %q", lines[0], "Please enter your Email Id:")
	}
	if lines[1] != "Your username is:  avimax37" {
		t.Errorf("line 1 = %q, want %q", lines[1], "Your username is:  avimax37")
	}
	if lines[2] != "Your domain is:  gmail.com" {
		t.Errorf("line 2 = %q, want %q", lines[2], "Your domain is:  gmail.com")
	}
}

func TestMain_InvalidInput_Output(t *testing.T) {
	output, err := captureMainOutput("invalidemail\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 output lines, got %d: %q", len(lines), output)
	}

	if lines[0] != "Please enter your Email Id:" {
		t.Errorf("line 0 = %q, want %q", lines[0], "Please enter your Email Id:")
	}
	if lines[1] != "Please enter a valid Email Id." {
		t.Errorf("line 1 = %q, want %q", lines[1], "Please enter a valid Email Id.")
	}
}

func TestMain_WhitespaceInput_Output(t *testing.T) {
	output, err := captureMainOutput("  avimax37@gmail.com  \n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 output lines, got %d: %q", len(lines), output)
	}

	if lines[1] != "Your username is:  avimax37" {
		t.Errorf("line 1 = %q, want %q", lines[1], "Your username is:  avimax37")
	}
	if lines[2] != "Your domain is:  gmail.com" {
		t.Errorf("line 2 = %q, want %q", lines[2], "Your domain is:  gmail.com")
	}
}

func TestMain_EmptyInput_Output(t *testing.T) {
	output, err := captureMainOutput("\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 output lines, got %d: %q", len(lines), output)
	}

	if lines[1] != "Please enter a valid Email Id." {
		t.Errorf("line 1 = %q, want %q", lines[1], "Please enter a valid Email Id.")
	}
}

func TestMain_MultipleAt_Output(t *testing.T) {
	output, err := captureMainOutput("user@sub@domain.com\n")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 output lines, got %d: %q", len(lines), output)
	}

	if lines[1] != "Your username is:  user" {
		t.Errorf("line 1 = %q, want %q", lines[1], "Your username is:  user")
	}
	if lines[2] != "Your domain is:  sub@domain.com" {
		t.Errorf("line 2 = %q, want %q", lines[2], "Your domain is:  sub@domain.com")
	}
}
