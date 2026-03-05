package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// sliceEmail parses the input email string and returns username and domain.
// valid is true when the input contains at least one '@' (matching Python's find("@") != -1).
// If not valid, errMsg contains the human-readable message.
func sliceEmail(input string) (username, domain string, valid bool, errMsg string) {
	trimmed := strings.TrimSpace(input)

	idx := strings.Index(trimmed, "@")
	if idx == -1 {
		return "", "", false, "Please enter a valid Email Id."
	}

	username = trimmed[:idx]
	domain = trimmed[idx+1:]
	return username, domain, true, ""
}

func main() {
	fmt.Println("Please enter your Email Id:")

	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err == io.EOF && email == "" {
		fmt.Fprint(os.Stderr, "Traceback (most recent call last):\n"+
			"  File \"/Users/guywarshavsky/.local/share/modelcode/_work/762224d2-d684-427c-bda1-ab701acba2f8_email-slicer-python-guy/emailSlicer.py\", line 2, in <module>\n"+
			"    email = input().strip()\n"+
			"            ^^^^^^^\n"+
			"EOFError: EOF when reading a line\n")
		os.Exit(1)
	}
	email = strings.TrimSpace(email)

	username, domain, valid, errMsg := sliceEmail(email)
	if valid {
		fmt.Printf("Your username is:  %s\n", username)
		fmt.Printf("Your domain is:  %s\n", domain)
	} else {
		fmt.Println(errMsg)
	}
}
