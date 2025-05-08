package utils

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	// Test cases table
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		// Basic cases
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"Simple word", "hello", "olleh"},
		{"Sentence with spaces", "Hello World", "dlroW olleH"},

		// Numbers and special characters
		{"Numbers only", "12345", "54321"},
		{"Mixed alphanumeric", "abc123", "321cba"},
		{"Special characters", "!@#$%^", "^%$#@!"},
		{"Mixed with special", "a!b@c#", "#c@b!a"},

		// Unicode cases
		{"Russian text", "привет", "тевирп"},
		{"Chinese text", "你好世界", "界世好你"},
		{"Arabic text", "مرحبا", "ابحرم"},
		{"Emoji", "😀🌍🚀", "🚀🌍😀"},

		// Edge cases
		{"Mixed languages", "Hello привет 你好", "好你 тевирп olleH"},
		{"With newlines", "Hello\nWorld", "dlroW\nolleH"},
		{"With tabs", "Hello\tWorld", "dlroW\tolleH"},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReverseString(tc.input)
			if result != tc.expected {
				t.Errorf("ReverseString(%q) = %q; expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

// Benchmark the ReverseString function
func BenchmarkReverseString(b *testing.B) {
	// Sample text to reverse (mix of ASCII and Unicode)
	text := "Hello World! Привет мир! 你好世界! 😀🌍🚀"

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		ReverseString(text)
	}
}
