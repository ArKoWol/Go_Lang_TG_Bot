package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	BotToken     string
	OpenAIAPIKey string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	// Get bot token from environment variables
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, err
	}

	// Get OpenAI API key from environment variables
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")

	return &Config{
		BotToken:     botToken,
		OpenAIAPIKey: openAIAPIKey,
	}, nil
}
