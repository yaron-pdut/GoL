# GoL

## Description

This is a simple Go project demonstrating:
- A basic "Hello, World!" program
- Use of goroutines and sync.WaitGroup
- Unit testing of goroutine output

## Project Structure

- `main.go`: Main application file
- `print_message_test.go`: Unit test for the goroutine function
- `go.mod`: Go module definition

## How to Run

1. Navigate to the project directory:
   ```sh
   cd GoL
   ```
2. Run the program:
   ```sh
   go run main.go
   ```
   Output:
   ```
   Hello, World!
   Hello from goroutine!
   ```

## How to Test

Run the unit tests with:
```sh
go test
```

## Requirements
- Go 1.18 or newer

## Notes
- Demonstrates basic concurrency in Go using goroutines and WaitGroup.
- Includes a unit test that captures and checks stdout output from a goroutine.