package utils

// ReverseString reverses the input string
// It properly handles UTF-8 characters including multi-byte sequences
func ReverseString(text string) string {
	// Convert string to rune slice to properly handle UTF-8 characters
	runes := []rune(text)

	// Reverse the rune slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert back to string
	return string(runes)
}
