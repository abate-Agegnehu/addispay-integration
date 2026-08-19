package main

import (
	"log"
	"net/http"

	"addispay-integration/internal/config"
	"addispay-integration/internal/handlers"
	"addispay-integration/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services
	addisPayService := services.NewAddisPayService(cfg)

	// Initialize handlers
	paymentHandler := handlers.NewPaymentHandler(addisPayService)
	webhookHandler := handlers.NewWebhookHandler(addisPayService)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/payments/initiate", paymentHandler.InitiatePayment)
	mux.HandleFunc("/api/payments/status", paymentHandler.GetPaymentStatus)
	mux.HandleFunc("/api/webhook", webhookHandler.HandleWebhook)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, corsMiddleware(mux)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
