# Bankr

A CLI tool for analyzing bank transactions across multiple accounts. It processes CSV transaction files and provides summary statistics and categorized spending analysis.

## Features

- Multi-account transaction processing
- Concurrent CSV file reading
- Transaction categorization
- Summary statistics including:
  - Total amount spent/received
  - Number of transactions
  - Date range analysis
  - Category-based summaries

## Prerequisites

- Go 1.24.2 or higher

## Installation

1. Clone the repository
2. Install dependencies:

```sh
go mod download
```

## Usage

1. Place your bank transaction CSV files in the `internal/resources` directory
2. CSV files should have the following header format:
   ```
   Type,Details,Particulars,Code,Reference,Amount,Date,ForeignCurrencyAmount,ConversionCharge
   ```
3. Run the application:
   ```sh
   go run .
   ```

## Project Structure

```
├── bankr.go           # Main application setup
├── main.go           # Entry point
├── internal/
│   ├── processor.go  # Transaction processing logic
│   ├── util.go      # Utility functions
│   ├── io/
│   │   └── reader.go # File reading operations
│   └── model/
│       ├── operations.go  # Generic operations
│       ├── summary.go     # Summary data structures
│       └── transaction.go # Transaction data structures
```

## Dependencies

- [github.com/fatih/color](https://github.com/fatih/color) - Terminal color output

## License

MIT