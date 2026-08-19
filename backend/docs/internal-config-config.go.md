# `internal/config/config.go`

Loads backend configuration from environment variables and `.env`.

## How It Works

`LoadConfig` reads AddisPay credentials, callback settings, and the server port. `getEnv` returns an environment value when present and otherwise uses a safe default.

## Main Variables

- `PORT`
- `ADDISPAY_BASE_URL`
- `ADDISPAY_API_KEY`
- `ADDISPAY_MERCHANT_ID`
- `ADDISPAY_WEBHOOK_SECRET`
- `APP_BASE_URL`

Keep real credentials in `backend/.env`; never commit them.
