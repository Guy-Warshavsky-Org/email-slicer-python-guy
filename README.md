# Email Slicer

A simple command-line tool that takes an email address as input and extracts the username and domain parts.

Migrated from the original [Python email slicer](https://github.com/gwarshavsky/email-slicer-python-guy) to Go.

## Prerequisites

- [Go 1.24](https://go.dev/dl/) or later

## Build

```bash
go build -o email-slicer
```

## Run

```bash
./email-slicer
```

The program will prompt you:

```
Please enter your Email Id:
```

Type an email address and press Enter. For example:

```
Please enter your Email Id:
user@example.com
Your username is:  user
Your domain is:  example.com
```

If the input does not contain `@`, the program prints:

```
Please enter a valid Email Id.
```

You can also pipe input directly:

```bash
echo "user@example.com" | ./email-slicer
```

## Test

```bash
go test -v ./...
```

This runs both unit tests for the `parseEmail` function and end-to-end tests that build and invoke the compiled binary.
