package canvas

import (
	"strings"
	"testing"
)

func TestUTF8Canvas_BasicOperations(t *testing.T) {
	canvas := NewUTF8CanvasWithSize(20, 10)
	
	// Test ASCII character
	canvas.SetCharAt(0, 0, "A")
	if canvas.Grid[0][0] != "A" {
		t.Errorf("ASCII character not set correctly")
	}
	
	// Test full-width character
	canvas.SetCharAt(2, 0, "あ")
	if canvas.Grid[0][2] != "あ" {
		t.Errorf("Full-width character not set correctly")
	}
	if canvas.Grid[0][3] != "" {
		t.Errorf("Full-width character continuation not set correctly")
	}
}

func TestUTF8Canvas_WriteText(t *testing.T) {
	canvas := NewUTF8CanvasWithSize(20, 10)
	
	// Test mixed text
	canvas.WriteText(0, 0, "Hello世界")
	
	// Check that characters are placed correctly
	if canvas.Grid[0][0] != "H" {
		t.Errorf("Expected 'H' at position 0, got '%s'", canvas.Grid[0][0])
	}
	if canvas.Grid[0][5] != "世" {
		t.Errorf("Expected '世' at position 5, got '%s'", canvas.Grid[0][5])
	}
	if canvas.Grid[0][6] != "" {
		t.Errorf("Expected empty continuation at position 6, got '%s'", canvas.Grid[0][6])
	}
	if canvas.Grid[0][7] != "界" {
		t.Errorf("Expected '界' at position 7, got '%s'", canvas.Grid[0][7])
	}
}

func TestUTF8Canvas_DrawBox(t *testing.T) {
	canvas := NewUTF8CanvasWithSize(10, 5)
	
	canvas.DrawBox(0, 0, 9, 4)
	
	// Check corners
	if canvas.Grid[0][0] != "+" {
		t.Errorf("Top-left corner not correct")
	}
	if canvas.Grid[0][9] != "+" {
		t.Errorf("Top-right corner not correct")
	}
	if canvas.Grid[4][0] != "+" {
		t.Errorf("Bottom-left corner not correct")
	}
	if canvas.Grid[4][9] != "+" {
		t.Errorf("Bottom-right corner not correct")
	}
	
	// Check borders
	if canvas.Grid[0][1] != "-" {
		t.Errorf("Top border not correct")
	}
	if canvas.Grid[1][0] != "|" {
		t.Errorf("Left border not correct")
	}
}

func TestUTF8Canvas_String(t *testing.T) {
	canvas := NewUTF8CanvasWithSize(10, 3)
	
	canvas.WriteText(0, 0, "Hello")
	canvas.WriteText(0, 1, "世界")
	
	result := canvas.String()
	lines := strings.Split(result, "\n")
	
	if len(lines) != 2 {
		t.Errorf("Expected 2 lines, got %d", len(lines))
	}
	
	if lines[0] != "Hello" {
		t.Errorf("First line incorrect: '%s'", lines[0])
	}
	
	if lines[1] != "世界" {
		t.Errorf("Second line incorrect: '%s'", lines[1])
	}
}