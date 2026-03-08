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
// To match the original Python behavior, empty username or domain parts
// are allowed (e.g. "@domain.com", "user@", or just "@").
func sliceEmail(raw string) (string, string, error) {
	trimmed := strings.TrimSpace(raw)

	at := strings.Index(trimmed, "@")
	if at == -1 {
		return "", "", errors.New("missing @")
	}

	username := trimmed[:at]
	domain := trimmed[at+1:]

	return username, domain, nil
}

func main() {
	fmt.Println("Please enter your Email Id:")
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

	fmt.Printf("Your username is:  %s\n", username)
	fmt.Printf("Your domain is:  %s\n", domain)
}
