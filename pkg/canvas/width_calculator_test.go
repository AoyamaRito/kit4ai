package canvas

import (
	"testing"
)

func TestCalculateDisplayWidth(t *testing.T) {
	wc := NewWidthCalculator()
	
	tests := []struct {
		input    string
		expected int
	}{
		{"Hello", 5},                    // ASCII only
		{"こんにちは", 10},                  // 5 hiragana characters = 10 width
		{"Hello世界", 9},                 // 5 ASCII + 2 kanji = 5 + 4 = 9
		{"Test123", 7},                  // ASCII mixed
		{"テスト", 6},                     // 3 katakana = 6 width
		{"", 0},                         // Empty string
		{"Ａｂｃ", 6},                     // 3 full-width ASCII = 6 width
		{"Hello こんにちは World", 21},     // Mixed with spaces
	}
	
	for _, test := range tests {
		result := wc.CalculateDisplayWidth(test.input)
		if result != test.expected {
			t.Errorf("CalculateDisplayWidth(%q) = %d, expected %d", 
				test.input, result, test.expected)
		}
	}
}

func TestRuneWidth(t *testing.T) {
	wc := NewWidthCalculator()
	
	tests := []struct {
		input    rune
		expected int
	}{
		{'A', 1},      // ASCII letter
		{'1', 1},      // ASCII digit
		{' ', 1},      // ASCII space
		{'あ', 2},      // Hiragana
		{'カ', 2},      // Katakana
		{'漢', 2},      // Kanji
		{'Ａ', 2},      // Full-width ASCII
		{'１', 2},      // Full-width digit
		{'？', 2},      // Full-width punctuation
	}
	
	for _, test := range tests {
		result := wc.RuneWidth(test.input)
		if result != test.expected {
			t.Errorf("RuneWidth(%q) = %d, expected %d", 
				test.input, result, test.expected)
		}
	}
}

func TestTruncateToWidth(t *testing.T) {
	wc := NewWidthCalculator()
	
	tests := []struct {
		input     string
		maxWidth  int
		expected  string
	}{
		{"Hello", 3, "Hel"},
		{"こんにちは", 6, "こんに"},        // 3 characters = 6 width
		{"Hello世界", 7, "Hello世"},       // 5 + 2 = 7 width exactly
		{"Test", 10, "Test"},             // No truncation needed
		{"", 5, ""},                      // Empty string
		{"Ａｂｃ", 4, "Ａｂ"},             // 2 full-width chars = 4 width
	}
	
	for _, test := range tests {
		result := wc.TruncateToWidth(test.input, test.maxWidth)
		if result != test.expected {
			t.Errorf("TruncateToWidth(%q, %d) = %q, expected %q", 
				test.input, test.maxWidth, result, test.expected)
		}
	}
}

func TestPadToWidth(t *testing.T) {
	wc := NewWidthCalculator()
	
	tests := []struct {
		input      string
		width      int
		padLeft    bool
		expected   string
	}{
		{"Hello", 10, false, "Hello     "},  // Right padding
		{"Hello", 10, true, "     Hello"},   // Left padding
		{"こんにちは", 12, false, "こんにちは "},   // Japanese right pad
		{"Test", 3, false, "Test"},          // No padding needed
		{"", 5, false, "     "},             // Empty string padding
	}
	
	for _, test := range tests {
		result := wc.PadToWidth(test.input, test.width, test.padLeft)
		if result != test.expected {
			t.Errorf("PadToWidth(%q, %d, %t) = %q, expected %q", 
				test.input, test.width, test.padLeft, result, test.expected)
		}
		
		// Verify the result has correct display width
		actualWidth := wc.CalculateDisplayWidth(result)
		expectedWidth := test.width
		if len(test.input) > 0 && wc.CalculateDisplayWidth(test.input) >= test.width {
			expectedWidth = wc.CalculateDisplayWidth(test.input)
		}
		
		if actualWidth != expectedWidth {
			t.Errorf("PadToWidth result has wrong width: got %d, expected %d", 
				actualWidth, expectedWidth)
		}
	}
}