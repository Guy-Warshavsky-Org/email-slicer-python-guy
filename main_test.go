package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// Unit tests for SliceEmail
// ---------------------------------------------------------------------------

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

func TestSliceEmail_Whitespace(t *testing.T) {
	username, domain, ok := SliceEmail("  user@example.com  ")
	if !ok {
		t.Fatal("expected ok=true after trimming whitespace")
	}
	if username != "user" {
		t.Errorf("username = %q, want %q", username, "user")
	}
	if domain != "example.com" {
		t.Errorf("domain = %q, want %q", domain, "example.com")
	}
}

func TestSliceEmail_Empty(t *testing.T) {
	_, _, ok := SliceEmail("")
	if ok {
		t.Fatal("expected ok=false for empty input")
	}
}

func TestSliceEmail_WhitespaceOnly(t *testing.T) {
	_, _, ok := SliceEmail("   ")
	if ok {
		t.Fatal("expected ok=false for whitespace-only input")
	}
}

// ---------------------------------------------------------------------------
// CLI-level integration tests
// ---------------------------------------------------------------------------

// runMain feeds input to main() via os.Stdin and captures os.Stdout output.
func runMain(input string) string {
	// Save originals
	origStdin := os.Stdin
	origStdout := os.Stdout
	defer func() {
		os.Stdin = origStdin
		os.Stdout = origStdout
	}()

	// Create a pipe for stdin
	rIn, wIn, _ := os.Pipe()
	os.Stdin = rIn
	go func() {
		io.WriteString(wIn, input)
		wIn.Close()
	}()

	// Create a pipe for stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	main()

	wOut.Close()
	var buf bytes.Buffer
	io.Copy(&buf, rOut)

	return buf.String()
}

func TestMain_ValidInput_Output(t *testing.T) {
	output := runMain("avimax37@gmail.com\n")
	if !strings.Contains(output, "Your username is:  avimax37") {
		t.Errorf("output missing username line, got:\n%s", output)
	}
	if !strings.Contains(output, "Your domain is:  gmail.com") {
		t.Errorf("output missing domain line, got:\n%s", output)
	}
}

func TestMain_InvalidInput_Output(t *testing.T) {
	output := runMain("invalidemail\n")
	if !strings.Contains(output, "Please enter a valid Email Id.") {
		t.Errorf("output missing error message, got:\n%s", output)
	}
}
