# `src/components/PaymentForm.jsx`

Collects the information required to create an AddisPay hosted checkout order.

## How It Works

Maintains controlled form state for amount, currency, description, customer name, email, and phone number. On submit it converts the amount to a number, maps camelCase fields to the backend's snake_case contract, adds success/cancel callback URLs, and calls `initiatePayment`.

When the backend returns `checkout_url` and `uuid`, the browser redirects to `checkout_url/uuid`.
