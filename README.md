# Learn Go with Tests

A comprehensive collection of Go exercises following the Test-Driven Development (TDD) approach. This repository contains my journey through learning Go by writing tests first.

## Project Structure

```
learn-go-with-tests/
├── arrays/                  # Working with arrays and slices
├── blogrenderer/           # HTML template rendering with snapshots
├── concurrency/            # Goroutines and concurrent operations
├── context/                # Context package for cancellation
├── di/                     # Dependency injection patterns
├── generics/               # Go generics with type parameters
├── helloworld/             # Basic Hello World implementation
├── integers/               # Integer operations and examples
├── iteration/              # Loops and iteration patterns
├── maps/                   # Dictionary/map operations
├── math/                   # Mathematical calculations (clockface SVG)
├── mocking/                # Mocking and test doubles
├── pointers/               # Pointer operations and wallet example
├── reading-files/          # File I/O and blog post parsing
├── reflection/             # Reflection package usage
├── roman-numerals/         # Roman numeral conversion
├── select/                 # Select statement and racing
├── structs/                # Struct definitions and methods
└── sync/                   # Synchronization primitives
```

## Topics Covered

### Fundamentals
- **Hello World**: Basic Go syntax and testing
- **Integers**: Working with integer types and operations
- **Iteration**: Loops and repetition
- **Arrays & Slices**: Collection data structures
- **Structs**: Custom types and methods
- **Pointers**: Memory addresses and pointer operations
- **Maps**: Key-value data structures

### Intermediate Concepts
- **Dependency Injection**: Decoupling code for testability
- **Mocking**: Creating test doubles and mocks
- **Concurrency**: Goroutines and concurrent execution
- **Select**: Managing multiple channel operations
- **Reflection**: Runtime type inspection
- **Sync**: Mutexes and synchronization primitives
- **Context**: Request cancellation and timeouts

### Advanced Topics
- **Generics**: Type parameters and generic functions
- **Reading Files**: File I/O and parsing
- **Blog Renderer**: Template rendering with snapshot testing
- **Roman Numerals**: Algorithm implementation with TDD
- **Math/SVG**: Generating SVG graphics (clockface)

## Getting Started

### Prerequisites
- Go 1.18 or higher (for generics support)
- Basic understanding of programming concepts

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test ./arrays
go test ./concurrency
```

Run tests with verbose output:
```bash
go test -v ./...
```

Run benchmarks:
```bash
go test -bench=. ./concurrency
```

### Running Examples

Some packages include executable examples:

**Clockface SVG Generator:**
```bash
go run ./math/clockface
```

## Key Learning Points

- **Test-Driven Development (TDD)**: Write tests before implementation
- **Go Testing**: Using the `testing` package effectively
- **Benchmarking**: Performance testing with Go's benchmark tools
- **Snapshot Testing**: Testing rendered output (blogrenderer)
- **Concurrent Programming**: Safe concurrent code with goroutines
- **Interface Design**: Writing flexible, testable code
- **Error Handling**: Go's explicit error handling patterns

## Project Highlights

### Concurrency Package
Demonstrates safe concurrent operations with benchmark tests showing performance improvements through parallelization.

### Blog Renderer
Uses Go templates to render HTML with snapshot testing to verify output consistency.

### Generics
Shows practical use of Go generics with a type-safe stack implementation and custom assertion helpers.

### Context Package
Illustrates proper context usage for cancellation and timeouts in HTTP handlers.

## Resources

This project follows the [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) guide by Chris James.

## License

This project is licensed under the terms specified in the LICENSE file.

## Development

### IDE
This project uses GoLand/IntelliJ IDEA (configuration in `.idea/` directory).

### Dependencies
Dependencies are managed with Go modules. See `go.mod` for the complete list.

```bash
# Install dependencies
go mod download

# Tidy dependencies
go mod tidy
```

## Contributing

This is a personal learning repository, but feel free to use it as a reference for your own Go journey!

## Progress Tracking

- ✅ Hello World
- ✅ Integers
- ✅ Iteration
- ✅ Arrays and Slices
- ✅ Structs, Methods & Interfaces
- ✅ Pointers & Errors
- ✅ Maps
- ✅ Dependency Injection
- ✅ Mocking
- ✅ Concurrency
- ✅ Select
- ✅ Reflection
- ✅ Sync
- ✅ Context
- ✅ Reading Files
- ✅ Generics
- ✅ Blog Renderer
- ✅ Math/SVG Graphics
- ✅ Roman Numerals

---

**Happy Testing! 🚀**