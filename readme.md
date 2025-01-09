# Wallet SDK

A Go-based wallet management system that handles user accounts, transactions, and balances.

## Project Structure

```
wallet-sdk/
├── cmd/                # Main application entry point
│   └── main.go
├── pkg/                # Core business logic
│   ├── users/
│   ├── wallets/
│   ├── transactions/
│   ├── fincrime/
│   └── admin/
├── internal/           # Internal utilities
│   ├── config/
│   ├── database/
│   └── logger/
├── migrations/         # Database migration scripts
├── test/               # Test cases
└── go.mod

```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- Docker (optional)

### Environment Setup

Create a `.env` file in the root directory:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=wallet_user
DB_PASSWORD=wallet_password
DB_NAME=wallet_db
```

### Running with Docker (to run the psql database)

```bash
docker-compose up -d
```

### Running without Docker

```bash
go run cmd/main.go
```

### Running Tests

```bash
go test -v ./...
```

### Run specific package tests

```bash
go test ./pkg/users/...
```

## API Examples

### Create a user

```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{
"tenant_id": 123,
"name": "John Doe",
"email": "john@example.com"
}'
```

### Get a user

```bash
curl http://localhost:8080/users/1
```

## Development

### Adding New Features

1. Create new package under `pkg/`
2. Add tests for new functionality
3. Update migrations if needed
4. Add API endpoints in `cmd/main.go`

### Code Style

- Follow Go standard formatting (`go fmt`)
- Use interfaces for dependency injection
- Write tests for new functionality
- Document public functions and types

```

```
