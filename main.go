package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// sliceEmail parses the input email string and returns username and domain.
// valid is true when the split at the first '@' yields non-empty username and domain.
// If not valid, errMsg contains the human-readable message (matching Python behavior).
func sliceEmail(input string) (username, domain string, valid bool, errMsg string) {
	trimmed := strings.TrimSpace(input)

	idx := strings.Index(trimmed, "@")
	if idx == -1 || idx == 0 || idx == len(trimmed)-1 {
		return "", "", false, "Please enter a valid Email Id."
	}

	username = trimmed[:idx]
	domain = trimmed[idx+1:]
	return username, domain, true, ""
}

func main() {
	fmt.Println("Please enter your Email Id:")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	email := strings.TrimSpace(line)

	username, domain, valid, errMsg := sliceEmail(email)
	if valid {
		fmt.Printf("Your username is:  %s\n", username)
		fmt.Printf("Your domain is:  %s\n", domain)
	} else {
		fmt.Println(errMsg)
	}
}
