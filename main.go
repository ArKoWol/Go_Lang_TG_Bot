package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get bot token from environment variables
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is not set in .env file")
	}

	// Set up bot configuration
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// Create new bot
	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Handle /start command
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! I'm your Telegram bot.")
	})

	// Handle /help command
	bot.Handle("/help", func(c tele.Context) error {
		return c.Send("Available commands:\n/start - Start the bot\n/help - Show this help message")
	})

	// Start the bot
	log.Println("Bot started!")
	bot.Start()
} 