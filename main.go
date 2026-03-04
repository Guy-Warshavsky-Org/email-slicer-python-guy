package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// sliceEmail trims whitespace from the input and splits it at the first '@'.
// It returns username, domain, and a boolean indicating whether the input was valid.
func sliceEmail(input string) (username, domain string, valid bool) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", "", false
	}

	idx := strings.Index(trimmed, "@")
	if idx < 0 {
		return "", "", false
	}

	username = trimmed[:idx]
	domain = trimmed[idx+1:]
	return username, domain, true
}

func main() {
	fmt.Println("Please enter your Email Id:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	username, domain, valid := sliceEmail(input)
	if valid {
		fmt.Println("Your username is: ", username)
		fmt.Println("Your domain is: ", domain)
	} else {
		fmt.Println("Please enter a valid Email Id.")
	}
}
