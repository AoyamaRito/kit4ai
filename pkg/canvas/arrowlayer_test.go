package canvas

import (
	"testing"
)

func TestNewArrowLayer(t *testing.T) {
	al := NewArrowLayer()
	
	if al == nil {
		t.Fatal("NewArrowLayer() returned nil")
	}
	
	if len(al.arrows) != 0 {
		t.Errorf("Expected 0 arrows, got %d", len(al.arrows))
	}
	
	if al.GetZOrder() != 0 {
		t.Errorf("Expected z-order 0, got %d", al.GetZOrder())
	}
}

func TestAddArrow(t *testing.T) {
	al := NewArrowLayer()
	
	al.AddArrow(0, 0, 10, 0, ArrowStyleNormal)
	
	if len(al.arrows) != 1 {
		t.Errorf("Expected 1 arrow, got %d", len(al.arrows))
	}
	
	arrow := al.arrows[0]
	if arrow.StartX != 0 || arrow.StartY != 0 || arrow.EndX != 10 || arrow.EndY != 0 {
		t.Errorf("Arrow coordinates incorrect: got (%d,%d) to (%d,%d)", 
			arrow.StartX, arrow.StartY, arrow.EndX, arrow.EndY)
	}
	
	if arrow.Style != ArrowStyleNormal {
		t.Errorf("Expected ArrowStyleNormal, got %v", arrow.Style)
	}
}

func TestAddLabeledArrow(t *testing.T) {
	al := NewArrowLayer()
	
	al.AddLabeledArrow(0, 0, 10, 0, ArrowStyleNormal, "test")
	
	if len(al.arrows) != 1 {
		t.Errorf("Expected 1 arrow, got %d", len(al.arrows))
	}
	
	arrow := al.arrows[0]
	if arrow.Label != "test" {
		t.Errorf("Expected label 'test', got '%s'", arrow.Label)
	}
}

func TestHorizontalArrow(t *testing.T) {
	// Set standard config for consistent testing
	SetConfig(StandardConfig)
	
	canvas := NewByteCanvas()
	al := NewArrowLayer()
	
	// Add horizontal arrow
	al.AddArrow(5, 5, 15, 5, ArrowStyleNormal)
	
	err := al.Render(canvas)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	// Check that line is drawn
	for x := 5; x < 15; x++ {
		if canvas.Grid[5][x] != '-' {
			t.Errorf("Expected '-' at (%d,5), got '%c'", x, canvas.Grid[5][x])
		}
	}
	
	// Check arrow head
	if canvas.Grid[5][15] != '>' {
		t.Errorf("Expected '>' at (15,5), got '%c'", canvas.Grid[5][15])
	}
}

func TestVerticalArrow(t *testing.T) {
	// Set standard config for consistent testing
	SetConfig(StandardConfig)
	
	canvas := NewByteCanvas()
	al := NewArrowLayer()
	
	// Add vertical arrow
	al.AddArrow(10, 2, 10, 8, ArrowStyleNormal)
	
	err := al.Render(canvas)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	// Check that line is drawn
	for y := 2; y < 8; y++ {
		if canvas.Grid[y][10] != '|' {
			t.Errorf("Expected '|' at (10,%d), got '%c'", y, canvas.Grid[y][10])
		}
	}
	
	// Check arrow head
	if canvas.Grid[8][10] != 'v' {
		t.Errorf("Expected 'v' at (10,8), got '%c'", canvas.Grid[8][10])
	}
}

func TestThickArrow(t *testing.T) {
	// Set standard config for consistent testing
	SetConfig(StandardConfig)
	
	canvas := NewByteCanvas()
	al := NewArrowLayer()
	
	// Add thick horizontal arrow
	al.AddArrow(0, 3, 8, 3, ArrowStyleThick)
	
	err := al.Render(canvas)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	// Check that thick line is drawn
	for x := 0; x < 8; x++ {
		if canvas.Grid[3][x] != '=' {
			t.Errorf("Expected '=' at (%d,3), got '%c'", x, canvas.Grid[3][x])
		}
	}
	
	// Check thick arrow head
	if canvas.Grid[3][8] != '>' {
		t.Errorf("Expected '>' at (8,3), got '%c'", canvas.Grid[3][8])
	}
}

func TestArrowWithLabel(t *testing.T) {
	// Set standard config for consistent testing
	SetConfig(StandardConfig)
	
	canvas := NewByteCanvas()
	al := NewArrowLayer()
	
	// Add arrow with label
	al.AddLabeledArrow(5, 5, 15, 5, ArrowStyleNormal, "HTTP")
	
	err := al.Render(canvas)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	
	// Check that label is drawn above the arrow
	labelY := 4
	labelX := 10 // Middle of arrow
	expectedLabel := "HTTP"
	
	for i, char := range expectedLabel {
		if canvas.Grid[labelY][labelX+i] != byte(char) {
			t.Errorf("Expected '%c' at (%d,%d), got '%c'", 
				char, labelX+i, labelY, canvas.Grid[labelY][labelX+i])
		}
	}
}

func TestClearArrows(t *testing.T) {
	al := NewArrowLayer()
	
	al.AddArrow(0, 0, 10, 0, ArrowStyleNormal)
	al.AddArrow(0, 2, 10, 2, ArrowStyleThick)
	
	if len(al.arrows) != 2 {
		t.Errorf("Expected 2 arrows, got %d", len(al.arrows))
	}
	
	al.Clear()
	
	if len(al.arrows) != 0 {
		t.Errorf("Expected 0 arrows after clear, got %d", len(al.arrows))
	}
}

func TestRemoveArrow(t *testing.T) {
	al := NewArrowLayer()
	
	al.AddArrow(0, 0, 10, 0, ArrowStyleNormal)
	al.AddArrow(0, 2, 10, 2, ArrowStyleThick)
	al.AddArrow(0, 4, 10, 4, ArrowStyleWave)
	
	// Remove middle arrow
	err := al.RemoveArrow(1)
	if err != nil {
		t.Fatalf("RemoveArrow failed: %v", err)
	}
	
	if len(al.arrows) != 2 {
		t.Errorf("Expected 2 arrows after removal, got %d", len(al.arrows))
	}
	
	// Check that correct arrow was removed
	if al.arrows[0].Style != ArrowStyleNormal {
		t.Errorf("First arrow should be normal style")
	}
	if al.arrows[1].Style != ArrowStyleWave {
		t.Errorf("Second arrow should be wave style")
	}
}

func TestRemoveArrowOutOfRange(t *testing.T) {
	al := NewArrowLayer()
	
	al.AddArrow(0, 0, 10, 0, ArrowStyleNormal)
	
	// Try to remove arrow that doesn't exist
	err := al.RemoveArrow(5)
	if err == nil {
		t.Error("Expected error when removing arrow out of range")
	}
	
	// Try negative index
	err = al.RemoveArrow(-1)
	if err == nil {
		t.Error("Expected error when removing arrow with negative index")
	}
}

func TestZOrder(t *testing.T) {
	al := NewArrowLayer()
	
	if al.GetZOrder() != 0 {
		t.Errorf("Expected initial z-order 0, got %d", al.GetZOrder())
	}
	
	al.SetZOrder(5)
	
	if al.GetZOrder() != 5 {
		t.Errorf("Expected z-order 5, got %d", al.GetZOrder())
	}
}

func TestArrowStyles(t *testing.T) {
	al := NewArrowLayer()
	
	styles := []ArrowStyle{
		ArrowStyleNormal,
		ArrowStyleThick,
		ArrowStyleWave,
		ArrowStyleShort,
		ArrowStyleDouble,
		ArrowStyleDotted,
	}
	
	for _, style := range styles {
		lineChar, arrowHead := al.getArrowComponents(style)
		
		if lineChar == 0 || arrowHead == 0 {
			t.Errorf("Invalid components for style %v: line='%c', head='%c'", 
				style, lineChar, arrowHead)
		}
	}
}