package handlers

import (
	"log"
	"telegram-bot/internal/services"

	tele "gopkg.in/telebot.v3"
)

// RegisterHandlers registers all bot handlers
func RegisterHandlers(bot *tele.Bot, storyService *services.StoryService) {
	// Handle /start command
	bot.Handle("/start", handleStart)

	// Handle /help command
	bot.Handle("/help", handleHelp)

	// Use the reverse handler for all text messages
	bot.Handle(tele.OnText, HandleReverse)

	// Create and register the story handler
	storyHandler := NewStoryHandler(storyService)
	bot.Handle("/story", storyHandler.HandleStory)

	log.Println("Handlers registered successfully")
}

// handleStart handles the /start command
func handleStart(c tele.Context) error {
	return c.Send("👋 Hello! I'm your Telegram bot. Use /help to see available commands.")
}

// handleHelp handles the /help command
func handleHelp(c tele.Context) error {
	helpText := "Available commands:\n" +
		"/start - Start the bot\n" +
		"/help - Show this help message\n" +
		"/story - Get a short sci-fi story\n" +
		"Any text message - I'll reverse it for you!"
	return c.Send(helpText)
}

// handleText handles text messages - this is now replaced by HandleReverse
func handleText(c tele.Context) error {
	log.Printf("Received message: %s", c.Message().Text)
	return c.Send("I received your message: " + c.Message().Text)
}
