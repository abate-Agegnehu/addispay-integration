# AddisPay Integration

A full-stack AddisPay hosted-checkout integration with a Go backend and React/Vite frontend.

## Project Structure

```text
addispay-integration/
├── backend/
│   ├── cmd/server/main.go
│   ├── internal/config/
│   ├── internal/handlers/
│   ├── internal/models/
│   ├── internal/services/
│   └── docs/
├── frontend/
│   ├── src/main.jsx
│   ├── src/App.jsx
│   ├── src/api.js
│   ├── src/components/
│   └── docs/
└── README.md
```

The `backend/docs/` and `frontend/docs/` folders contain one Markdown document for each documented project file.

## Architecture

The browser never calls AddisPay directly. The frontend sends payment data to the Go API, and the Go API communicates with AddisPay using the private API key.

```text
Customer
   |
   v
React/Vite frontend :5173
   |
   | POST /api/payments/initiate
   v
Go backend :8080
   |
   | POST /checkout-api/v2/create-order
   v
AddisPay UAT API
   |
   | checkout_url + uuid
   v
AddisPay hosted checkout
   |
   | redirect to success/cancel URL
   v
React callback page
```

## Payment Flow

1. The customer fills out the payment form.
2. `PaymentForm.jsx` converts the amount and sends the request to the backend.
3. The backend validates the request in `payment.go`.
4. `addispay.go` creates a unique `nonce` and `tx_ref`.
5. The service sends a wrapped order request to AddisPay's `/checkout-api/v2/create-order` endpoint.
6. AddisPay returns a nested `data.uuid` and `data.checkout_url`.
7. The backend normalizes those fields and returns them to the browser.
8. The frontend redirects to `checkout_url/uuid`.
9. AddisPay redirects the customer to `/payment-success` or `/payment-cancel`.
10. The success page displays callback values such as amount, currency, UUID, and transaction reference.

## Backend API

### `POST /api/payments/initiate`

Creates an AddisPay hosted checkout order.

Example request:

```json
{
  "amount": 1000,
  "currency": "ETB",
  "description": "test payment",
  "customer_email": "customer@example.com",
  "customer_name": "Customer Name",
  "phone_number": "251911111111",
  "return_url": "http://localhost:5173/payment-success",
  "cancel_url": "http://localhost:5173/payment-cancel"
}
```

### `GET /api/payments/status`

Requests payment status using the `transaction_id` query parameter.

### `POST /api/webhook`

Receives AddisPay webhook events and validates the `X-Webhook-Signature` header.

## Configuration

Create `backend/.env` locally:

```env
PORT=8080
ADDISPAY_BASE_URL=https://uat.api.addispay.et
ADDISPAY_API_KEY=your_uat_api_key
ADDISPAY_MERCHANT_ID=your_merchant_id
ADDISPAY_WEBHOOK_SECRET=your_webhook_secret
APP_BASE_URL=http://localhost:5173
```

Create or update `frontend/.env`:

```env
VITE_API_BASE_URL=http://localhost:8080
```

Do not commit real API keys, webhook secrets, or other credentials.

## Run Locally

Start the backend from `backend/`:

```powershell
go test ./...
go build -o server.exe ./cmd/server
.\server.exe
```

The compiled executable approach avoids Windows Application Control policies that may block Go's temporary `go run` executable.

Start the frontend from `frontend/`:

```powershell
npm install
npm run dev
```

Open:

- Frontend: `http://localhost:5173`
- Backend: `http://localhost:8080`

## Verification

Backend:

```powershell
cd backend
go test ./...
```

Frontend:

```powershell
cd frontend
npm run build
```

## Security Notes

- Keep `backend/.env` out of version control.
- Rotate credentials if they are exposed.
- Keep the AddisPay API key on the backend only.
- Verify webhook signatures before processing events.
- Replace the webhook TODO logging with persistent order/status updates before production use.
