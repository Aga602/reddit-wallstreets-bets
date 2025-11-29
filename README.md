# reddit-wallstreets-bets

Agent that crawls through the wall street bets subreddit and determines market sentiments

## Prerequisites

- Go 1.21 or later

## Project Structure

```
.
├── cmd/
│   └── greeter/          # Application entry point
│       └── main.go
├── internal/
│   └── greeter/          # Greeting logic package
│       ├── greeter.go
│       └── greeter_test.go
├── go.mod
├── Makefile
└── README.md
```

## Quick Start

### Build

```bash
make build
```

### Run

```bash
# Run with default greeting
make run

# Run with a name
./bin/greeter -name "Alice"
```

### Test

```bash
make test
```

### Lint

```bash
make lint
```

### Clean

```bash
make clean
```

## Usage

The greeter application accepts the following flags:

- `-name`: Name of the person to greet (optional, defaults to "World")
