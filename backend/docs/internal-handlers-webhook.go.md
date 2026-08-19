# `internal/handlers/webhook.go`

Handles AddisPay webhook notifications.

## How It Works

The handler accepts only `POST`, reads the raw body, verifies `X-Webhook-Signature`, decodes `WebhookPayload`, and logs the payment event. It returns `200 OK` after acknowledgement.

Database updates, emails, and other business actions are marked as future work in the source file.
