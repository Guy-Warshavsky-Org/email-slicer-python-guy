# Email Slicer

A simple CLI tool that takes an email address as input and extracts the username and domain parts.

This is a Go port of the original Python [emailSlicer.py](https://github.com/gwarshavsky/email-slicer-python-guy) script, with behavior verified by automated tests.

## Prerequisites

- [Go](https://go.dev/) 1.21 or later

## Build

```bash
go build -o email-slicer
```

## Run

Run interactively:

```bash
./email-slicer
```

Or pipe input directly:

```bash
echo "avimax37@gmail.com" | ./email-slicer
```

### Example Output

**Valid email:**

```
Please enter your Email Id:
Your username is:  avimax37
Your domain is:  gmail.com
```

**Invalid email (no `@`):**

```
Please enter your Email Id:
Please enter a valid Email Id.
```

## Test

```bash
go test ./...
```

Run with verbose output:

```bash
go test -v ./...
```
