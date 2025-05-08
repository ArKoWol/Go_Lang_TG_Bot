package handlers

import (
	"log"
	"telegram-bot/internal/utils"

	tele "gopkg.in/telebot.v3"
)

// ReverseText reverses the input string
func ReverseText(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// HandleReverse handles text messages and responds with reversed text
func HandleReverse(c tele.Context) error {
	originalText := c.Message().Text
	log.Printf("Received message to reverse: %s", originalText)

	// Use the utils.ReverseString function
	reversedText := utils.ReverseString(originalText)

	return c.Send(reversedText)
}
