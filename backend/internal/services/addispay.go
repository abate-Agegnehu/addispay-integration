package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"addispay-integration/internal/config"
	"addispay-integration/internal/models"
)

type AddisPayService struct {
	config     *config.Config
	httpClient *http.Client
}

func NewAddisPayService(cfg *config.Config) *AddisPayService {
	return &AddisPayService{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// InitiatePayment creates a new payment session
func (s *AddisPayService) InitiatePayment(req *models.PaymentRequest) (*models.AddisPayPaymentResponse, error) {
	// Prepare request to AddisPay
	firstName, lastName := splitName(req.CustomerName)
	reference := strconv.FormatInt(time.Now().UnixNano(), 10)
	addisReq := models.AddisPayOrderRequest{
		Data: models.AddisPayOrderData{
			RedirectURL:    req.ReturnURL,
			CancelURL:      req.CancelURL,
			SuccessURL:     req.ReturnURL,
			ErrorURL:       req.CancelURL,
			OrderReason:    req.Description,
			Currency:       req.Currency,
			Email:          req.CustomerEmail,
			FirstName:      firstName,
			LastName:       lastName,
			PhoneNumber:    req.PhoneNumber,
			Nonce:          reference,
			OrderDetail:    map[string]interface{}{"amount": req.Amount, "description": req.Description},
			SessionExpired: "5000",
			TotalAmount:    req.Amount,
			TxRef:          reference,
		},
		Message: "Payment initiated",
	}

	jsonData, err := json.Marshal(addisReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	httpReq, err := http.NewRequest("POST", strings.TrimRight(s.config.AddisPayBaseURL, "/")+"/checkout-api/v2/create-order", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Auth", s.config.AddisPayAPIKey)

	// Send request
	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("AddisPay API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var addisResp models.AddisPayPaymentResponse
	if err := json.Unmarshal(body, &addisResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if addisResp.UUID == "" {
		addisResp.UUID = addisResp.Data.UUID
	}
	if addisResp.CheckoutURL == "" {
		addisResp.CheckoutURL = addisResp.Data.CheckoutURL
	}
	if addisResp.Amount == "" && addisResp.Data.TotalAmount != 0 {
		addisResp.Amount = strconv.FormatFloat(addisResp.Data.TotalAmount, 'f', -1, 64)
	}

	if addisResp.CheckoutURL == "" || addisResp.UUID == "" {
		return nil, fmt.Errorf("AddisPay error: %s", providerError(addisResp))
	}

	return &addisResp, nil
}

// VerifyWebhookSignature verifies the webhook signature
func (s *AddisPayService) VerifyWebhookSignature(payload []byte, signature string) bool {
	// Create HMAC-SHA256 hash of the payload
	h := hmac.New(sha256.New, []byte(s.config.AddisPayWebhookSecret))
	h.Write(payload)
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare signatures (constant time)
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

// GetPaymentStatus retrieves the status of a payment
func (s *AddisPayService) GetPaymentStatus(transactionID string) (*models.PaymentStatusResponse, error) {
	url := fmt.Sprintf("%s/v1/payments/%s", s.config.AddisPayBaseURL, transactionID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.config.AddisPayAPIKey)
	req.Header.Set("X-Merchant-ID", s.config.AddisPayMerchantID)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AddisPay API error (status %d): %s", resp.StatusCode, string(body))
	}

	var statusResp models.PaymentStatusResponse
	if err := json.Unmarshal(body, &statusResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &statusResp, nil
}

func splitName(name string) (string, string) {
	parts := strings.Fields(name)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], parts[0]
	}
	return parts[0], strings.Join(parts[1:], " ")
}

func providerError(resp models.AddisPayPaymentResponse) string {
	if resp.Message != "" {
		return resp.Message
	}
	if resp.Error != "" {
		return resp.Error
	}
	return "invalid response"
}
