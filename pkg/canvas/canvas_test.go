package canvas

import (
	"strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas()
	
	if c.Width != DefaultWidth {
		t.Errorf("Expected width %d, got %d", DefaultWidth, c.Width)
	}
	
	if c.Height != DefaultHeight {
		t.Errorf("Expected height %d, got %d", DefaultHeight, c.Height)
	}
	
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			if c.Grid[y][x] != ' ' {
				t.Errorf("Expected space at (%d,%d), got %c", x, y, c.Grid[y][x])
			}
		}
	}
}

func TestReplaceChar(t *testing.T) {
	c := NewCanvas()
	
	err := c.ReplaceChar(10, 10, 'X')
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	char, err := c.GetChar(10, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if char != 'X' {
		t.Errorf("Expected 'X', got %c", char)
	}
}

func TestReplaceCharOutOfBounds(t *testing.T) {
	c := NewCanvas()
	
	err := c.ReplaceChar(-1, 0, 'X')
	if err == nil {
		t.Error("Expected error for negative x coordinate")
	}
	
	err = c.ReplaceChar(0, -1, 'X')
	if err == nil {
		t.Error("Expected error for negative y coordinate")
	}
	
	err = c.ReplaceChar(DefaultWidth, 0, 'X')
	if err == nil {
		t.Error("Expected error for x coordinate out of bounds")
	}
	
	err = c.ReplaceChar(0, DefaultHeight, 'X')
	if err == nil {
		t.Error("Expected error for y coordinate out of bounds")
	}
}

func TestGetChar(t *testing.T) {
	c := NewCanvas()
	c.ReplaceChar(5, 5, 'Y')
	
	char, err := c.GetChar(5, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if char != 'Y' {
		t.Errorf("Expected 'Y', got %c", char)
	}
}

func TestClear(t *testing.T) {
	c := NewCanvas()
	c.ReplaceChar(10, 10, 'Z')
	c.ReplaceChar(20, 20, 'A')
	
	c.Clear()
	
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			if c.Grid[y][x] != ' ' {
				t.Errorf("Expected space at (%d,%d) after clear, got %c", x, y, c.Grid[y][x])
			}
		}
	}
}

func TestString(t *testing.T) {
	c := NewCanvas()
	
	emptyResult := c.String()
	if emptyResult != "" {
		t.Errorf("Expected empty string for empty canvas, got: %q", emptyResult)
	}
	
	c.ReplaceChar(0, 0, 'A')
	c.ReplaceChar(1, 0, 'B')
	c.ReplaceChar(0, 1, 'C')
	
	result := c.String()
	lines := strings.Split(result, "\n")
	
	if len(lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(lines))
	}
	
	if !strings.HasPrefix(lines[0], "AB") {
		t.Errorf("Expected first line to start with 'AB', got: %q", lines[0])
	}
	
	if !strings.HasPrefix(lines[1], "C") {
		t.Errorf("Expected second line to start with 'C', got: %q", lines[1])
	}
}