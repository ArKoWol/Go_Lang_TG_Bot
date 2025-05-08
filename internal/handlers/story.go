package handlers

import (
	"log"
	"telegram-bot/internal/services"

	tele "gopkg.in/telebot.v3"
)

// StoryHandler handles requests for sci-fi stories
type StoryHandler struct {
	storyService *services.StoryService
}

// NewStoryHandler creates a new StoryHandler
func NewStoryHandler(storyService *services.StoryService) *StoryHandler {
	return &StoryHandler{
		storyService: storyService,
	}
}

// HandleStory generates and sends a short sci-fi story
func (h *StoryHandler) HandleStory(c tele.Context) error {
	log.Println("Generating a sci-fi story...")

	// Let the user know the bot is working on the request
	_ = c.Notify(tele.Typing)

	// Generate the story
	story, err := h.storyService.GenerateSciFiStory()
	if err != nil {
		log.Printf("Error generating story: %v", err)
		return c.Send("Sorry, I couldn't generate a story right now. Please try again later.")
	}

	// Send the story to the user
	return c.Send(story)
}
