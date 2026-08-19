# `src/api.js`

Axios client for communication with the Go backend.

## How It Works

The base URL comes from `VITE_API_BASE_URL`, defaulting to `http://localhost:8080`. `initiatePayment` posts form data to `/api/payments/initiate`; `getPaymentStatus` requests `/api/payments/status`. Both functions log failures and rethrow a user-displayable error.
