# Portfolio Profitability Calculator

This Go project simulates a portfolio with buy, sell, and calculates the overall profitability, taking transaction fees into account.

## Requirements

- Go 1.18 or newer

## Running the Project

1. **Clone or download the repository.**
2. Open a terminal and navigate to the project directory.

### Run the Main Program

To run the main program:

```sh
go run main.go
```

### Run Unit Tests

To run all unit tests:

```sh
go test
```

To run a specific test:

```sh
go test -run TestOnlySellTransactions
```

## Project Structure

- `main.go` — Contains the portfolio logic and transaction processing.
- `validate_test.go` — Contains unit tests for edge cases and core functionality.

## Notes

- All transaction fees are deducted from profitability.
