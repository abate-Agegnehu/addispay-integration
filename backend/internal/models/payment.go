package models

import "time"

// PaymentRequest represents the request to initiate a payment
type PaymentRequest struct {
	Amount        float64 `json:"amount" validate:"required,gt=0"`
	Currency      string  `json:"currency" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	CustomerEmail string  `json:"customer_email" validate:"required,email"`
	CustomerName  string  `json:"customer_name" validate:"required"`
	PhoneNumber   string  `json:"phone_number" validate:"required"`
	ReturnURL     string  `json:"return_url" validate:"required,url"`
	CancelURL     string  `json:"cancel_url" validate:"required,url"`
}

// AddisPayOrderRequest is the request sent to AddisPay's hosted checkout API.
type AddisPayOrderRequest struct {
	Data    AddisPayOrderData `json:"data"`
	Message string            `json:"message"`
}

type AddisPayOrderData struct {
	RedirectURL    string                 `json:"redirect_url"`
	CancelURL      string                 `json:"cancel_url"`
	SuccessURL     string                 `json:"success_url"`
	ErrorURL       string                 `json:"error_url"`
	OrderReason    string                 `json:"order_reason"`
	Currency       string                 `json:"currency"`
	Email          string                 `json:"email"`
	PhoneNumber    string                 `json:"phone_number"`
	FirstName      string                 `json:"first_name"`
	LastName       string                 `json:"last_name"`
	Nonce          string                 `json:"nonce"`
	OrderDetail    map[string]interface{} `json:"order_detail"`
	SessionExpired string                 `json:"session_expired"`
	TotalAmount    float64                `json:"total_amount"`
	TxRef          string                 `json:"tx_ref"`
}

// AddisPayPaymentResponse is the response from AddisPay API
type AddisPayPaymentResponse struct {
	Amount      string `json:"amount"`
	CheckoutURL string `json:"checkout_url"`
	UUID        string `json:"uuid"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	Error       string `json:"error,omitempty"`
	Data        struct {
		UUID        string  `json:"uuid"`
		TxRef       string  `json:"tx_ref"`
		TotalAmount float64 `json:"total_amount"`
		Currency    string  `json:"currency"`
		CheckoutURL string  `json:"checkout_url"`
	} `json:"data"`
}

// PaymentStatusResponse represents the payment status
type PaymentStatusResponse struct {
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"` // pending, completed, failed, cancelled
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	PaymentMethod string    `json:"payment_method,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// WebhookPayload represents the webhook payload from AddisPay
type WebhookPayload struct {
	Event         string    `json:"event"` // payment.completed, payment.failed, etc.
	TransactionID string    `json:"transaction_id"`
	PaymentID     string    `json:"payment_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	CustomerEmail string    `json:"customer_email"`
	Timestamp     time.Time `json:"timestamp"`
	Signature     string    `json:"signature"`
}
