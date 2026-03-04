package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func run(r io.Reader, w io.Writer) error {
	fmt.Fprintln(w, "Please enter your Email Id:")

	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return fmt.Errorf("EOFError: EOF when reading a line")
	}
	email := strings.TrimSpace(scanner.Text())

	if strings.Contains(email, "@") {
		idx := strings.Index(email, "@")
		username := email[:idx]
		domain := email[idx+1:]
		fmt.Fprintln(w, "Your username is: ", username)
		fmt.Fprintln(w, "Your domain is: ", domain)
	} else {
		fmt.Fprintln(w, "Please enter a valid Email Id.")
	}

	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
