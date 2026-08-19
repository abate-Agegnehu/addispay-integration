# `internal/services/addispay.go`

Contains the AddisPay API integration.

## How It Works

`InitiatePayment` creates unique transaction values, maps customer data into AddisPay's order format, and sends a request to `/checkout-api/v2/create-order` using the `Auth` header. It normalizes nested `data.uuid` and `data.checkout_url` fields before returning them to the frontend.

The service also verifies webhook signatures with HMAC-SHA256 and requests payment status from AddisPay.
