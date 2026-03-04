package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestValidEmail(t *testing.T) {
	input := strings.NewReader("avimax37@gmail.com\n")
	var output bytes.Buffer

	err := run(input, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := output.String()

	if !strings.Contains(got, "Your username is:  avimax37") {
		t.Errorf("expected output to contain 'Your username is:  avimax37', got:\n%s", got)
	}
	if !strings.Contains(got, "Your domain is:  gmail.com") {
		t.Errorf("expected output to contain 'Your domain is:  gmail.com', got:\n%s", got)
	}
}

func TestInvalidEmail(t *testing.T) {
	input := strings.NewReader("invalidemail\n")
	var output bytes.Buffer

	err := run(input, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := output.String()

	if !strings.Contains(got, "Please enter a valid Email Id.") {
		t.Errorf("expected output to contain 'Please enter a valid Email Id.', got:\n%s", got)
	}
}

func TestMultipleAtSymbols(t *testing.T) {
	input := strings.NewReader("user@sub@example.com\n")
	var output bytes.Buffer

	err := run(input, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := output.String()

	if !strings.Contains(got, "Your username is:  user") {
		t.Errorf("expected output to contain 'Your username is:  user', got:\n%s", got)
	}
	if !strings.Contains(got, "Your domain is:  sub@example.com") {
		t.Errorf("expected output to contain 'Your domain is:  sub@example.com', got:\n%s", got)
	}
}

func TestLeadingTrailingWhitespace(t *testing.T) {
	input := strings.NewReader("   user@example.com  \n")
	var output bytes.Buffer

	err := run(input, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := output.String()

	if !strings.Contains(got, "Your username is:  user") {
		t.Errorf("expected output to contain 'Your username is:  user', got:\n%s", got)
	}
	if !strings.Contains(got, "Your domain is:  example.com") {
		t.Errorf("expected output to contain 'Your domain is:  example.com', got:\n%s", got)
	}
}

func TestExactMessageConformity(t *testing.T) {
	// Test that the exact prompt message is printed
	t.Run("prompt message", func(t *testing.T) {
		input := strings.NewReader("test@example.com\n")
		var output bytes.Buffer

		err := run(input, &output)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := output.String()
		lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

		if len(lines) < 1 || lines[0] != "Please enter your Email Id:" {
			t.Errorf("expected first line to be 'Please enter your Email Id:', got: %q", lines[0])
		}
	})

	// Test exact valid email output format (double space before values)
	t.Run("valid email format", func(t *testing.T) {
		input := strings.NewReader("avimax37@gmail.com\n")
		var output bytes.Buffer

		err := run(input, &output)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := output.String()
		lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

		if len(lines) != 3 {
			t.Fatalf("expected 3 lines of output, got %d: %q", len(lines), got)
		}
		if lines[1] != "Your username is:  avimax37" {
			t.Errorf("expected line 2 to be 'Your username is:  avimax37', got: %q", lines[1])
		}
		if lines[2] != "Your domain is:  gmail.com" {
			t.Errorf("expected line 3 to be 'Your domain is:  gmail.com', got: %q", lines[2])
		}
	})

	// Test exact invalid email output format
	t.Run("invalid email format", func(t *testing.T) {
		input := strings.NewReader("nope\n")
		var output bytes.Buffer

		err := run(input, &output)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		got := output.String()
		lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

		if len(lines) != 2 {
			t.Fatalf("expected 2 lines of output, got %d: %q", len(lines), got)
		}
		if lines[1] != "Please enter a valid Email Id." {
			t.Errorf("expected line 2 to be 'Please enter a valid Email Id.', got: %q", lines[1])
		}
	})
}
