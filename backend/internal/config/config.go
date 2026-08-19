package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    Port               string
    AddisPayBaseURL    string
    AddisPayAPIKey     string
    AddisPayMerchantID string
    AddisPayWebhookSecret string
    AppBaseURL         string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using environment variables")
    }

    return &Config{
        Port:               getEnv("PORT", "8080"),
        AddisPayBaseURL:    getEnv("ADDISPAY_BASE_URL", "https://uat.api.addispay.et"),
        AddisPayAPIKey:     getEnv("ADDISPAY_API_KEY", ""),
        AddisPayMerchantID: getEnv("ADDISPAY_MERCHANT_ID", ""),
        AddisPayWebhookSecret: getEnv("ADDISPAY_WEBHOOK_SECRET", ""),
        AppBaseURL:         getEnv("APP_BASE_URL", "http://localhost:5173"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        return value
    }
    return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        if intVal, err := strconv.Atoi(value); err == nil {
            return intVal
        }
    }
    return defaultValue
}