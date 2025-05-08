package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

// StoryService handles story generation using OpenAI
type StoryService struct {
	client *openai.Client
}

// NewStoryService creates a new StoryService
func NewStoryService(apiKey string) (*StoryService, error) {
	if apiKey == "" {
		return nil, errors.New("OpenAI API key is required")
	}

	client := openai.NewClient(apiKey)
	return &StoryService{
		client: client,
	}, nil
}

// GenerateSciFiStory generates a short sci-fi story using OpenAI
func (s *StoryService) GenerateSciFiStory() (string, error) {
	prompt := "Generate a short sci-fi story (maximum 400 characters). Make it creative and engaging with a twist at the end."

	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a creative sci-fi writer. Write very short stories (400 characters max) with unexpected twists.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens:   150, // Approximately 400 characters
			Temperature: 0.8, // More creative
		},
	)

	if err != nil {
		log.Printf("OpenAI API error: %v", err)
		return "", fmt.Errorf("failed to generate story: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", errors.New("no response from OpenAI")
	}

	story := resp.Choices[0].Message.Content
	return story, nil
}
