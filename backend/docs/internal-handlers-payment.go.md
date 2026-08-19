# `internal/handlers/payment.go`

HTTP handlers for payment initiation and payment status requests.

## How It Works

`InitiatePayment` accepts JSON, validates the basic amount and currency fields, calls the AddisPay service, and returns the provider response. `GetPaymentStatus` reads `transaction_id` from the query string and requests the status from the service.

Validation failures return `400`. Provider or service failures return `500`.
