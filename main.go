package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SliceEmail takes a raw input string, trims whitespace, and splits it into
// username and domain on the first '@' character. Returns ok=false if the
// trimmed input does not contain '@'.
func SliceEmail(input string) (username, domain string, ok bool) {
	trimmed := strings.TrimSpace(input)
	idx := strings.Index(trimmed, "@")
	if idx == -1 {
		return "", "", false
	}
	return trimmed[:idx], trimmed[idx+1:], true
}

func main() {
	fmt.Println("Please enter your Email Id:")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	username, domain, ok := SliceEmail(line)
	if ok {
		fmt.Printf("Your username is:  %s\n", username)
		fmt.Printf("Your domain is:  %s\n", domain)
	} else {
		fmt.Println("Please enter a valid Email Id.")
	}
}
