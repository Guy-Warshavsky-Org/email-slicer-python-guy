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

	out := output.String()

	if !strings.Contains(out, "Your username is:  avimax37") {
		t.Errorf("expected output to contain 'Your username is:  avimax37', got:\n%s", out)
	}
	if !strings.Contains(out, "Your domain is:  gmail.com") {
		t.Errorf("expected output to contain 'Your domain is:  gmail.com', got:\n%s", out)
	}
}

func TestInvalidEmail(t *testing.T) {
	input := strings.NewReader("invalidemail\n")
	var output bytes.Buffer

	err := run(input, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := output.String()

	if !strings.Contains(out, "Please enter a valid Email Id.") {
		t.Errorf("expected output to contain 'Please enter a valid Email Id.', got:\n%s", out)
	}
}

func TestMultipleAtSymbols(t *testing.T) {
	input := strings.NewReader("user@sub@example.com\n")
	var output bytes.Buffer

	err := run(input, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := output.String()

	if !strings.Contains(out, "Your username is:  user") {
		t.Errorf("expected output to contain 'Your username is:  user', got:\n%s", out)
	}
	if !strings.Contains(out, "Your domain is:  sub@example.com") {
		t.Errorf("expected output to contain 'Your domain is:  sub@example.com', got:\n%s", out)
	}
}

func TestLeadingTrailingWhitespace(t *testing.T) {
	input := strings.NewReader("   user@example.com  \n")
	var output bytes.Buffer

	err := run(input, &output)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := output.String()

	if !strings.Contains(out, "Your username is:  user") {
		t.Errorf("expected output to contain 'Your username is:  user', got:\n%s", out)
	}
	if !strings.Contains(out, "Your domain is:  example.com") {
		t.Errorf("expected output to contain 'Your domain is:  example.com', got:\n%s", out)
	}
}

func TestExactMessageConformity(t *testing.T) {
	// Test that exact output lines match the Python script's format
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "prompt message",
			input: "test@example.com\n",
			expected: []string{
				"Please enter your Email Id:\n",
			},
		},
		{
			name:  "valid email output format",
			input: "avimax37@gmail.com\n",
			expected: []string{
				"Your username is:  avimax37\n",
				"Your domain is:  gmail.com\n",
			},
		},
		{
			name:  "invalid email output format",
			input: "invalidemail\n",
			expected: []string{
				"Please enter a valid Email Id.\n",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := strings.NewReader(tc.input)
			var output bytes.Buffer

			err := run(input, &output)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			out := output.String()
			for _, exp := range tc.expected {
				if !strings.Contains(out, exp) {
					t.Errorf("expected output to contain %q, got:\n%q", exp, out)
				}
			}
		})
	}
}
