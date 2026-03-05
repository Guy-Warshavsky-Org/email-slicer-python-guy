# Email Slicer

Email Slicer is a simple command-line tool that takes an email address as input and returns the username and domain as output. This is a Go port of the original [Python Email Slicer](https://github.com/gwarshavsky/email-slicer-python-guy).

## Built With

- [Go](https://go.dev/) 1.24+

## Example

Input:
```
Please enter your Email Id:
avimax37@gmail.com
```

Output:
```
Your username is:  avimax37
Your domain is:  gmail.com
```

Here we got **`avimax37`** as username and **`gmail.com`** as domain.

## Prerequisites

- [Go 1.24](https://go.dev/dl/) or later installed and available on your PATH.

## Building

```bash
go build -o email-slicer
```

## Running

Run interactively:
```bash
go run .
```

Or pipe input directly:
```bash
echo "avimax37@gmail.com" | go run .
```

## Testing

```bash
go test -v ./...
```

## License

Distributed under the MIT License. See **`LICENSE.md`** for more information.
