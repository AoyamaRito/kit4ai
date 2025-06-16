package canvas

import (
	"unicode"
	"unicode/utf8"
)

// WidthCalculator handles text width calculations for mixed ASCII/Japanese content
type WidthCalculator struct{}

// NewWidthCalculator creates a new width calculator
func NewWidthCalculator() *WidthCalculator {
	return &WidthCalculator{}
}

// CalculateDisplayWidth calculates the display width of a string
// considering full-width characters as 2 units and half-width as 1 unit
func (wc *WidthCalculator) CalculateDisplayWidth(text string) int {
	if !utf8.ValidString(text) {
		return len(text) // fallback for invalid UTF-8
	}
	
	width := 0
	for _, r := range text {
		width += wc.RuneWidth(r)
	}
	return width
}

// RuneWidth returns the display width of a single rune
func (wc *WidthCalculator) RuneWidth(r rune) int {
	// ASCII characters (0-127) are always width 1
	if r < 128 {
		return 1
	}
	
	// Full-width characters are width 2
	if wc.IsFullWidth(r) {
		return 2
	}
	
	// Other Unicode characters default to width 1
	return 1
}

// IsFullWidth determines if a rune is full-width
func (wc *WidthCalculator) IsFullWidth(r rune) bool {
	return unicode.Is(unicode.Han, r) ||           // 漢字 (CJK Unified Ideographs)
		unicode.Is(unicode.Hiragana, r) ||         // ひらがな
		unicode.Is(unicode.Katakana, r) ||         // カタカナ
		(r >= 0xFF01 && r <= 0xFF5E) ||           // 全角英数字・記号
		(r >= 0x3000 && r <= 0x303F) ||           // CJK記号・句読点
		(r >= 0x2E80 && r <= 0x2EFF) ||           // CJK部首補助
		(r >= 0x2F00 && r <= 0x2FDF) ||           // 康熙部首
		(r >= 0x31C0 && r <= 0x31EF) ||           // CJK筆画
		(r >= 0x3200 && r <= 0x32FF) ||           // 囲み文字
		(r >= 0x3300 && r <= 0x33FF) ||           // CJK互換文字
		(r >= 0xF900 && r <= 0xFAFF) ||           // CJK互換漢字
		(r >= 0xFE30 && r <= 0xFE4F)              // CJK互換形式
}

// TruncateToWidth truncates text to fit within specified display width
func (wc *WidthCalculator) TruncateToWidth(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}
	
	currentWidth := 0
	result := make([]rune, 0, len(text))
	
	for _, r := range text {
		runeWidth := wc.RuneWidth(r)
		if currentWidth+runeWidth > maxWidth {
			break
		}
		result = append(result, r)
		currentWidth += runeWidth
	}
	
	return string(result)
}

// PadToWidth pads text with spaces to reach the specified display width
func (wc *WidthCalculator) PadToWidth(text string, targetWidth int, padLeft bool) string {
	currentWidth := wc.CalculateDisplayWidth(text)
	if currentWidth >= targetWidth {
		return text
	}
	
	padding := targetWidth - currentWidth
	spaces := make([]rune, padding)
	for i := range spaces {
		spaces[i] = ' '
	}
	
	if padLeft {
		return string(spaces) + text
	}
	return text + string(spaces)
}

// SplitToFitWidth splits text into lines that fit within the specified width
func (wc *WidthCalculator) SplitToFitWidth(text string, maxWidth int) []string {
	if maxWidth <= 0 {
		return []string{}
	}
	
	words := []rune(text)
	if len(words) == 0 {
		return []string{}
	}
	
	var lines []string
	currentLine := make([]rune, 0)
	currentWidth := 0
	
	for _, r := range words {
		runeWidth := wc.RuneWidth(r)
		
		// If adding this character would exceed width, start new line
		if currentWidth+runeWidth > maxWidth && len(currentLine) > 0 {
			lines = append(lines, string(currentLine))
			currentLine = make([]rune, 0)
			currentWidth = 0
		}
		
		// Add character to current line
		currentLine = append(currentLine, r)
		currentWidth += runeWidth
	}
	
	// Add the last line if it has content
	if len(currentLine) > 0 {
		lines = append(lines, string(currentLine))
	}
	
	return lines
}