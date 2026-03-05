# Email Slicer

A simple command-line tool that takes an email address as input and extracts the username and domain parts.

Migrated from the original [Python implementation](https://github.com/gwarshavsky/email-slicer-python-guy).

## Prerequisites

- Go 1.24 or later

## Building

```bash
go build -o email-slicer
```

## Running

Run interactively:

```bash
go run .
```

Or with piped input:

```bash
echo "avimax37@gmail.com" | go run .
```

### Example Output

For a valid email:

```
Please enter your Email Id:
Your username is:  avimax37
Your domain is:  gmail.com
```

For an invalid email:

```
Please enter your Email Id:
Please enter a valid Email Id.
```

## Testing

```bash
go test -v ./...
```
