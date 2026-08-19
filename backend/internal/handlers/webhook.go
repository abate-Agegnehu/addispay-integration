package handlers

import (
    "encoding/json"
    "io"
    "net/http"
    "log"

    "addispay-integration/internal/models"
    "addispay-integration/internal/services"
)

type WebhookHandler struct {
    addisPayService *services.AddisPayService
}

func NewWebhookHandler(service *services.AddisPayService) *WebhookHandler {
    return &WebhookHandler{
        addisPayService: service,
    }
}

// HandleWebhook processes incoming webhooks from AddisPay
func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Read raw body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Failed to read body", http.StatusBadRequest)
        return
    }

    // Get signature from header
    signature := r.Header.Get("X-Webhook-Signature")
    if signature == "" {
        log.Println("Missing webhook signature")
        http.Error(w, "Missing signature", http.StatusUnauthorized)
        return
    }

    // Verify signature
    if !h.addisPayService.VerifyWebhookSignature(body, signature) {
        log.Println("Invalid webhook signature")
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }

    // Parse webhook payload
    var payload models.WebhookPayload
    if err := json.Unmarshal(body, &payload); err != nil {
        log.Printf("Failed to parse webhook payload: %v", err)
        http.Error(w, "Invalid payload", http.StatusBadRequest)
        return
    }

    // Process the webhook event
    log.Printf("Received webhook event: %s for transaction: %s", payload.Event, payload.TransactionID)

    // Update payment status in your database
    // TODO: Implement database update logic
    // e.g., update payment status based on payload.Status

    switch payload.Event {
    case "payment.completed":
        log.Printf("Payment completed: %s", payload.TransactionID)
        // Handle successful payment
        // Send confirmation email, update order status, etc.
    case "payment.failed":
        log.Printf("Payment failed: %s", payload.TransactionID)
        // Handle failed payment
    case "payment.cancelled":
        log.Printf("Payment cancelled: %s", payload.TransactionID)
        // Handle cancelled payment
    default:
        log.Printf("Unknown event type: %s", payload.Event)
    }

    // Always return 200 OK to acknowledge receipt
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}