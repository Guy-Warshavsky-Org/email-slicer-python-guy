package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// sliceEmail parses an email string and returns username and domain.
// It trims leading and trailing whitespace before validation.
// The first "@" is used as the separator, so additional "@" characters
// remain in the domain part.
func sliceEmail(raw string) (string, string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", "", errors.New("empty input")
	}

	at := strings.Index(trimmed, "@")
	if at == -1 {
		return "", "", errors.New("missing @")
	}

	username := trimmed[:at]
	domain := trimmed[at+1:]

	if username == "" || domain == "" {
		return "", "", errors.New("incomplete email")
	}

	return username, domain, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("Please enter a valid Email Id.")
		return
	}

	username, domain, err := sliceEmail(input)
	if err != nil {
		fmt.Println("Please enter a valid Email Id.")
		return
	}

	fmt.Printf("Your username is: %s\n", username)
	fmt.Printf("Your domain is: %s\n", domain)
}
