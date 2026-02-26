package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SliceEmail takes a trimmed email string and splits it into username and domain
// based on the first '@' character. It returns ok=false if no '@' is present
// or the input is empty.
func SliceEmail(input string) (username, domain string, ok bool) {
	if input == "" {
		return "", "", false
	}
	idx := strings.Index(input, "@")
	if idx == -1 {
		return "", "", false
	}
	return input[:idx], input[idx+1:], true
}

func main() {
	fmt.Println("Please enter your Email Id:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	username, domain, ok := SliceEmail(email)
	if ok {
		fmt.Printf("Your username is:  %s\n", username)
		fmt.Printf("Your domain is:  %s\n", domain)
	} else {
		fmt.Println("Please enter a valid Email Id.")
	}
}
