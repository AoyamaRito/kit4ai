package canvas

import (
	"strings"
)

// UTF8Canvas is a canvas that can handle UTF-8 characters properly
type UTF8Canvas struct {
	Width  int
	Height int
	Grid   [][]string // Store strings instead of bytes to handle UTF-8
	widthCalc *WidthCalculator
}

// NewUTF8Canvas creates a new UTF-8 capable canvas
func NewUTF8Canvas() *UTF8Canvas {
	width := GetCurrentWidth()
	height := GetCurrentHeight()
	return NewUTF8CanvasWithSize(width, height)
}

// NewUTF8CanvasWithSize creates a new UTF-8 canvas with specific dimensions
func NewUTF8CanvasWithSize(width, height int) *UTF8Canvas {
	canvas := &UTF8Canvas{
		Width:     width,
		Height:    height,
		Grid:      make([][]string, height),
		widthCalc: NewWidthCalculator(),
	}
	
	// Initialize grid with spaces
	for i := range canvas.Grid {
		canvas.Grid[i] = make([]string, width)
		for j := range canvas.Grid[i] {
			canvas.Grid[i][j] = " "
		}
	}
	
	return canvas
}

// SetCharAt sets a character at the specified position, accounting for width
func (c *UTF8Canvas) SetCharAt(x, y int, char string) {
	if x < 0 || y < 0 || y >= c.Height || x >= c.Width {
		return
	}
	
	// Calculate width of the character
	width := c.widthCalc.CalculateDisplayWidth(char)
	
	if x+width > c.Width {
		return // Character doesn't fit
	}
	
	// Clear any existing full-width character that might be here
	if x > 0 && c.Grid[y][x-1] != "" {
		// Check if previous character is full-width
		prevWidth := c.widthCalc.CalculateDisplayWidth(c.Grid[y][x-1])
		if prevWidth == 2 {
			// We're overwriting the second half of a full-width char
			c.Grid[y][x-1] = " "
		}
	}
	
	// Set the character
	c.Grid[y][x] = char
	
	// If it's a full-width character, mark the next position as occupied
	if width == 2 && x+1 < c.Width {
		c.Grid[y][x+1] = "" // Empty string indicates "continuation"
	}
}

// WriteText writes text starting at the specified position
func (c *UTF8Canvas) WriteText(x, y int, text string) {
	if y < 0 || y >= c.Height {
		return
	}
	
	currentX := x
	for _, r := range text {
		if currentX >= c.Width {
			break
		}
		
		char := string(r)
		width := c.widthCalc.RuneWidth(r)
		
		if currentX+width <= c.Width {
			c.SetCharAt(currentX, y, char)
			currentX += width
		} else {
			break
		}
	}
}

// DrawBox draws a box using ASCII characters
func (c *UTF8Canvas) DrawBox(x1, y1, x2, y2 int) {
	// Top and bottom borders
	for x := x1; x <= x2; x++ {
		c.SetCharAt(x, y1, "-")
		c.SetCharAt(x, y2, "-")
	}
	
	// Left and right borders  
	for y := y1; y <= y2; y++ {
		c.SetCharAt(x1, y, "|")
		c.SetCharAt(x2, y, "|")
	}
	
	// Corners
	c.SetCharAt(x1, y1, "+")
	c.SetCharAt(x2, y1, "+")
	c.SetCharAt(x1, y2, "+")
	c.SetCharAt(x2, y2, "+")
}

// DrawLine draws a line
func (c *UTF8Canvas) DrawLine(x1, y1, x2, y2 int) {
	if y1 == y2 {
		// Horizontal line
		start, end := x1, x2
		if start > end {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			c.SetCharAt(x, y1, "-")
		}
	} else if x1 == x2 {
		// Vertical line
		start, end := y1, y2
		if start > end {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			c.SetCharAt(x1, y, "|")
		}
	}
}

// String converts the canvas to a string representation
func (c *UTF8Canvas) String() string {
	lines := make([]string, 0, c.Height)
	
	// Find the last non-empty line
	lastNonEmptyLine := -1
	for i := 0; i < c.Height; i++ {
		lineHasContent := false
		for j := 0; j < c.Width; j++ {
			if c.Grid[i][j] != " " && c.Grid[i][j] != "" {
				lineHasContent = true
				break
			}
		}
		if lineHasContent {
			lastNonEmptyLine = i
		}
	}
	
	// Build output up to the last non-empty line
	for i := 0; i <= lastNonEmptyLine; i++ {
		line := strings.Builder{}
		for j := 0; j < c.Width; j++ {
			char := c.Grid[i][j]
			if char == "" {
				// This is a continuation position for a full-width character
				// Skip it completely
				continue
			}
			line.WriteString(char)
		}
		// Remove trailing spaces but preserve structure
		lineStr := strings.TrimRight(line.String(), " ")
		lines = append(lines, lineStr)
	}
	
	if len(lines) == 0 {
		return ""
	}
	
	return strings.Join(lines, "\n")
}

// CreateTable creates a table with proper Japanese text handling
func (c *UTF8Canvas) CreateTable(x, y int, headers []string, rows [][]string) {
	if len(headers) == 0 {
		return
	}
	
	// Calculate column widths based on actual character widths
	colWidths := make([]int, len(headers))
	
	// Initialize with header widths (display width + padding)
	for i, header := range headers {
		displayWidth := 0
		for _, r := range header {
			displayWidth += c.widthCalc.RuneWidth(r)
		}
		colWidths[i] = displayWidth + 2 // +2 for padding
	}
	
	// Check data rows for maximum width per column
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) {
				displayWidth := 0
				for _, r := range cell {
					displayWidth += c.widthCalc.RuneWidth(r)
				}
				cellWidth := displayWidth + 2 // +2 for padding
				if cellWidth > colWidths[i] {
					colWidths[i] = cellWidth
				}
			}
		}
	}
	
	currentY := y
	
	// Draw top border
	c.drawTableBorder(x, currentY, colWidths)
	currentY++
	
	// Draw headers
	c.drawTableRow(x, currentY, headers, colWidths)
	currentY++
	
	// Draw separator
	c.drawTableBorder(x, currentY, colWidths)
	currentY++
	
	// Draw data rows
	for _, row := range rows {
		c.drawTableRow(x, currentY, row, colWidths)
		currentY++
	}
	
	// Draw bottom border
	c.drawTableBorder(x, currentY, colWidths)
}

// drawTableBorder draws horizontal table border
func (c *UTF8Canvas) drawTableBorder(x, y int, colWidths []int) {
	currentX := x
	c.SetCharAt(currentX, y, "+")
	currentX++
	
	for _, width := range colWidths {
		for i := 0; i < width; i++ {
			c.SetCharAt(currentX, y, "-")
			currentX++
		}
		c.SetCharAt(currentX, y, "+")
		currentX++
	}
}

// drawTableRow draws a table row with proper text alignment
func (c *UTF8Canvas) drawTableRow(x, y int, cells []string, colWidths []int) {
	currentX := x
	c.SetCharAt(currentX, y, "|")
	currentX++
	
	for i, width := range colWidths {
		var cell string
		if i < len(cells) {
			cell = cells[i]
		}
		
		// Truncate cell if too long
		maxCellWidth := width - 2 // -2 for padding
		if c.widthCalc.CalculateDisplayWidth(cell) > maxCellWidth {
			cell = c.widthCalc.TruncateToWidth(cell, maxCellWidth)
		}
		
		// Write left padding
		c.SetCharAt(currentX, y, " ")
		currentX++
		
		// Write cell content
		cellStart := currentX
		c.WriteText(currentX, y, cell)
		
		// Calculate actual content width and pad to column width
		actualContentWidth := 0
		for _, r := range cell {
			actualContentWidth += c.widthCalc.RuneWidth(r)
		}
		
		// Fill remaining space with spaces to maintain column alignment
		remainingSpace := (width - 2) - actualContentWidth
		for j := 0; j < remainingSpace; j++ {
			c.SetCharAt(currentX + actualContentWidth + j, y, " ")
		}
		
		// Move currentX to end of column
		currentX = cellStart + (width - 2)
		
		// Write right padding  
		c.SetCharAt(currentX, y, " ")
		currentX++
		
		// Write separator
		c.SetCharAt(currentX, y, "|")
		currentX++
	}
}