package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Result holds the outcome of parsing an email address.
type Result struct {
	Valid    bool
	Username string
	Domain   string
}

// parseEmail trims whitespace and splits the input on the first '@'.
// If no '@' is found, Valid is false.
func parseEmail(input string) Result {
	trimmed := strings.TrimSpace(input)
	idx := strings.Index(trimmed, "@")
	if idx == -1 {
		return Result{Valid: false}
	}
	return Result{
		Valid:    true,
		Username: trimmed[:idx],
		Domain:   trimmed[idx+1:],
	}
}

func main() {
	fmt.Println("Please enter your Email Id:")

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	result := parseEmail(line)
	if result.Valid {
		fmt.Println("Your username is: ", result.Username)
		fmt.Println("Your domain is: ", result.Domain)
	} else {
		fmt.Println("Please enter a valid Email Id.")
	}
}
