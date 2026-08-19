# `src/components/PaymentSuccess.jsx`

Displays the AddisPay payment-success callback page.

## How It Works

Reads `uuid`, `tx_ref`, `transaction_id`, `amount`, `total_amount`, and `currency` from the callback query string. It displays whichever payment details AddisPay provides and links the customer back to the payment form for another transaction.
