package main

import (
	"log"
	"time"

	"telegram-bot/config"
	"telegram-bot/internal/handlers"
	"telegram-bot/internal/services"

	tele "gopkg.in/telebot.v3"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	// Set up bot configuration
	pref := tele.Settings{
		Token:  cfg.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// Create new bot
	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Failed to create bot: ", err)
	}

	// Initialize the story service
	var storyService *services.StoryService
	if cfg.OpenAIAPIKey != "" {
		storyService, err = services.NewStoryService(cfg.OpenAIAPIKey)
		if err != nil {
			log.Printf("Warning: Failed to initialize story service: %v", err)
			log.Println("The /story command will not work until you set OPENAI_API_KEY in .env")
		}
	} else {
		log.Println("Warning: OPENAI_API_KEY not set in .env")
		log.Println("The /story command will not work until you set this value")
	}

	// Register handlers
	handlers.RegisterHandlers(bot, storyService)

	// Start the bot
	log.Println("Bot started!")
	bot.Start()
}
