# `cmd/server/main.go`

Application entry point for the Go API.

## How It Works

- Loads environment configuration with `config.LoadConfig()`.
- Creates the AddisPay service and HTTP handlers.
- Registers payment initiation, payment status, and webhook routes.
- Wraps all routes with CORS middleware so browser preflight requests succeed.
- Starts the server on the configured `PORT`.

## Routes

- `POST /api/payments/initiate`
- `GET /api/payments/status`
- `POST /api/webhook`
