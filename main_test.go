package main

import (
	"bytes"
	"strings"
	"testing"
)

// helper runs the email slicer with the given input and returns the captured output.
func helper(t *testing.T, input string) string {
	t.Helper()
	r := strings.NewReader(input + "\n")
	var w bytes.Buffer
	err := run(r, &w)
	if err != nil {
		t.Fatalf("run() returned an error: %v", err)
	}
	return w.String()
}

func TestValidEmail(t *testing.T) {
	output := helper(t, "avimax37@gmail.com")

	if !strings.Contains(output, "Your username is:  avimax37") {
		t.Errorf("expected output to contain 'Your username is:  avimax37', got:\n%s", output)
	}
	if !strings.Contains(output, "Your domain is:  gmail.com") {
		t.Errorf("expected output to contain 'Your domain is:  gmail.com', got:\n%s", output)
	}
}

func TestInvalidEmail(t *testing.T) {
	output := helper(t, "invalidemail")

	if !strings.Contains(output, "Please enter a valid Email Id.") {
		t.Errorf("expected output to contain 'Please enter a valid Email Id.', got:\n%s", output)
	}
	// Should NOT contain username/domain output
	if strings.Contains(output, "Your username is:") {
		t.Errorf("unexpected username output for invalid email, got:\n%s", output)
	}
}

func TestMultipleAtSymbols(t *testing.T) {
	output := helper(t, "user@sub@example.com")

	if !strings.Contains(output, "Your username is:  user") {
		t.Errorf("expected username 'user' (split at first @), got:\n%s", output)
	}
	if !strings.Contains(output, "Your domain is:  sub@example.com") {
		t.Errorf("expected domain 'sub@example.com' (everything after first @), got:\n%s", output)
	}
}

func TestLeadingTrailingWhitespace(t *testing.T) {
	output := helper(t, "   user@example.com  ")

	if !strings.Contains(output, "Your username is:  user") {
		t.Errorf("expected whitespace to be trimmed, username should be 'user', got:\n%s", output)
	}
	if !strings.Contains(output, "Your domain is:  example.com") {
		t.Errorf("expected whitespace to be trimmed, domain should be 'example.com', got:\n%s", output)
	}
}

func TestExactMessageConformity(t *testing.T) {
	// Test that the prompt message is exact
	output := helper(t, "test@example.com")
	if !strings.Contains(output, "Please enter your Email Id:") {
		t.Errorf("expected prompt 'Please enter your Email Id:', got:\n%s", output)
	}

	// Test exact format of username line (double space before value)
	if !strings.Contains(output, "Your username is:  test") {
		t.Errorf("expected 'Your username is:  test' (with double space), got:\n%s", output)
	}

	// Test exact format of domain line (double space before value)
	if !strings.Contains(output, "Your domain is:  example.com") {
		t.Errorf("expected 'Your domain is:  example.com' (with double space), got:\n%s", output)
	}

	// Test invalid email message is exact
	invalidOutput := helper(t, "noemail")
	if !strings.Contains(invalidOutput, "Please enter a valid Email Id.") {
		t.Errorf("expected exact error message 'Please enter a valid Email Id.', got:\n%s", invalidOutput)
	}
}
