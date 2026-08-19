# `internal/models/payment.go`

Defines the data structures shared by HTTP handlers and the AddisPay service.

## Main Types

- `PaymentRequest`: data received from the frontend.
- `AddisPayOrderRequest`: wrapped request sent to AddisPay.
- `AddisPayOrderData`: hosted-checkout order fields.
- `AddisPayPaymentResponse`: provider order response.
- `PaymentStatusResponse`: payment status data.
- `WebhookPayload`: incoming webhook event data.

JSON tags keep Go field names aligned with the browser and AddisPay API contracts.
