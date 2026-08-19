# `src/components/PaymentStatus.jsx`

Displays payment status information for the cancellation callback route and legacy status behavior.

## How It Works

Reads `transaction_id` and an optional `status` from query parameters. If a status is supplied, it displays it immediately; otherwise it calls `getPaymentStatus`. It renders loading, error, no-status, and status-specific states.
