package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"addispay-integration/internal/models"
	"addispay-integration/internal/services"
)

type PaymentHandler struct {
	addisPayService *services.AddisPayService
}

func NewPaymentHandler(service *services.AddisPayService) *PaymentHandler {
	return &PaymentHandler{
		addisPayService: service,
	}
}

// InitiatePayment handles the payment initiation request
func (h *PaymentHandler) InitiatePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var req models.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := validatePaymentRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Initiate payment with AddisPay
	resp, err := h.addisPayService.InitiatePayment(&req)
	if err != nil {
		http.Error(w, "Failed to initiate payment: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

// GetPaymentStatus handles payment status check
func (h *PaymentHandler) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	transactionID := r.URL.Query().Get("transaction_id")
	if transactionID == "" {
		http.Error(w, "transaction_id is required", http.StatusBadRequest)
		return
	}

	status, err := h.addisPayService.GetPaymentStatus(transactionID)
	if err != nil {
		http.Error(w, "Failed to get payment status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func validatePaymentRequest(req *models.PaymentRequest) error {
	if req.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	if req.Currency == "" {
		return fmt.Errorf("currency is required")
	}
	// Add more validation as needed
	return nil
}
