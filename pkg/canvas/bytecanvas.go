package canvas

import (
	"errors"
	"strings"
	"unicode"
)

// ByteCanvas は8bit単位で処理するキャンバス
type ByteCanvas struct {
	Width  int
	Height int
	Grid   [][]byte
}

// NewByteCanvas creates a new byte-based canvas with current configuration
func NewByteCanvas() *ByteCanvas {
	width := GetCurrentWidth()
	height := GetCurrentHeight()
	
	bc := &ByteCanvas{
		Width:  width,
		Height: height,
		Grid:   make([][]byte, height),
	}
	
	for i := range bc.Grid {
		bc.Grid[i] = make([]byte, width)
		for j := range bc.Grid[i] {
			bc.Grid[i][j] = ' '
		}
	}
	
	return bc
}

// NewByteCanvasWithSize creates a new byte-based canvas with specific dimensions
func NewByteCanvasWithSize(width, height int) *ByteCanvas {
	bc := &ByteCanvas{
		Width:  width,
		Height: height,
		Grid:   make([][]byte, height),
	}
	
	for i := range bc.Grid {
		bc.Grid[i] = make([]byte, width)
		for j := range bc.Grid[i] {
			bc.Grid[i][j] = ' '
		}
	}
	
	return bc
}

// ReplaceByte replaces a byte at the specified position
func (bc *ByteCanvas) ReplaceByte(x, y int, b byte) error {
	if x < 0 || x >= bc.Width || y < 0 || y >= bc.Height {
		return errors.New("coordinates out of bounds")
	}
	
	bc.Grid[y][x] = b
	return nil
}

// GetByte gets a byte at the specified position
func (bc *ByteCanvas) GetByte(x, y int) (byte, error) {
	if x < 0 || x >= bc.Width || y < 0 || y >= bc.Height {
		return 0, errors.New("coordinates out of bounds")
	}
	
	return bc.Grid[y][x], nil
}

// Clear clears the canvas with spaces
func (bc *ByteCanvas) Clear() {
	for i := range bc.Grid {
		for j := range bc.Grid[i] {
			bc.Grid[i][j] = ' '
		}
	}
}

// IsFullWidth 全角文字かどうかを判定
func (bc *ByteCanvas) IsFullWidth(r rune) bool {
	return unicode.Is(unicode.Han, r) ||
		unicode.Is(unicode.Hiragana, r) ||
		unicode.Is(unicode.Katakana, r) ||
		(r >= 0xFF01 && r <= 0xFF5E) ||
		(r >= 0x3000 && r <= 0x303F)
}

// FilterASCII 全角文字を除去してASCII文字のみを残す
func (bc *ByteCanvas) FilterASCII(data string) string {
	runes := []rune(data)
	filtered := make([]rune, 0, len(runes))
	
	for _, r := range runes {
		if !bc.IsFullWidth(r) && r < 128 {
			filtered = append(filtered, r)
		}
	}
	
	return string(filtered)
}

// WriteBytes writes a byte string to the canvas
func (bc *ByteCanvas) WriteBytes(x, y int, data string) error {
	bytes := []byte(data)
	
	for i, b := range bytes {
		if x+i >= bc.Width || y < 0 || y >= bc.Height {
			break
		}
		bc.Grid[y][x+i] = b
	}
	
	return nil
}

// WriteBytesASCII writes ASCII-only text to the canvas (filters full-width chars)
func (bc *ByteCanvas) WriteBytesASCII(x, y int, data string) error {
	filteredData := bc.FilterASCII(data)
	return bc.WriteBytes(x, y, filteredData)
}

// String converts the canvas to a string
func (bc *ByteCanvas) String() string {
	lines := make([]string, 0, bc.Height)
	
	lastNonEmptyLine := -1
	for i := 0; i < bc.Height; i++ {
		line := string(bc.Grid[i])
		if strings.TrimSpace(line) != "" {
			lastNonEmptyLine = i
		}
	}
	
	for i := 0; i <= lastNonEmptyLine; i++ {
		lines = append(lines, string(bc.Grid[i]))
	}
	
	if len(lines) == 0 {
		return ""
	}
	
	return strings.Join(lines, "\n")
}

// DrawHorizontalLine draws a horizontal line
func (bc *ByteCanvas) DrawHorizontalLine(x1, x2, y int, char byte) {
	for x := x1; x <= x2; x++ {
		if x >= 0 && x < bc.Width && y >= 0 && y < bc.Height {
			bc.Grid[y][x] = char
		}
	}
}

// DrawVerticalLine draws a vertical line
func (bc *ByteCanvas) DrawVerticalLine(x, y1, y2 int, char byte) {
	for y := y1; y <= y2; y++ {
		if x >= 0 && x < bc.Width && y >= 0 && y < bc.Height {
			bc.Grid[y][x] = char
		}
	}
}

// DrawBox draws a box using ASCII characters
func (bc *ByteCanvas) DrawBox(x1, y1, x2, y2 int) {
	// 水平線
	bc.DrawHorizontalLine(x1, x2, y1, '-')
	bc.DrawHorizontalLine(x1, x2, y2, '-')
	
	// 垂直線
	bc.DrawVerticalLine(x1, y1, y2, '|')
	bc.DrawVerticalLine(x2, y1, y2, '|')
	
	// 角
	bc.ReplaceByte(x1, y1, '+')
	bc.ReplaceByte(x2, y1, '+')
	bc.ReplaceByte(x1, y2, '+')
	bc.ReplaceByte(x2, y2, '+')
}