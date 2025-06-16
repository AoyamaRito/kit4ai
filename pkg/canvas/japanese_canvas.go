package canvas

import ()

// JapaneseCanvas extends ByteCanvas with Japanese text support
type JapaneseCanvas struct {
	*ByteCanvas
	widthCalc *WidthCalculator
}

// NewJapaneseCanvas creates a new Japanese-aware canvas
func NewJapaneseCanvas() *JapaneseCanvas {
	return &JapaneseCanvas{
		ByteCanvas: NewByteCanvas(),
		widthCalc:  NewWidthCalculator(),
	}
}

// NewJapaneseCanvasWithSize creates a new Japanese-aware canvas with specific dimensions
func NewJapaneseCanvasWithSize(width, height int) *JapaneseCanvas {
	return &JapaneseCanvas{
		ByteCanvas: NewByteCanvasWithSize(width, height),
		widthCalc:  NewWidthCalculator(),
	}
}

// WriteJapaneseText writes Japanese text with proper width calculation
func (jc *JapaneseCanvas) WriteJapaneseText(x, y int, text string) error {
	// Check bounds
	if y < 0 || y >= jc.Height {
		return nil
	}
	
	currentX := x
	for _, r := range text {
		if currentX >= jc.Width {
			break
		}
		
		// Get rune width
		runeWidth := jc.widthCalc.RuneWidth(r)
		
		// Write the character (as UTF-8 bytes)
		if currentX+runeWidth <= jc.Width {
			// For ASCII characters, write directly
			if r < 128 {
				jc.Grid[y][currentX] = byte(r)
				currentX++
			} else {
				// For Japanese characters, we need special handling
				// Since ByteCanvas uses byte arrays, we'll mark the position
				// and handle the actual rendering differently
				jc.writeJapaneseRune(currentX, y, r, runeWidth)
				currentX += runeWidth
			}
		}
	}
	
	return nil
}

// writeJapaneseRune handles writing a Japanese rune to the canvas
func (jc *JapaneseCanvas) writeJapaneseRune(x, y int, r rune, width int) {
	// For now, we'll use a placeholder approach
	// In a real implementation, you'd need a different storage mechanism
	// that can handle UTF-8 properly
	
	if width == 2 && x+1 < jc.Width {
		// Mark as Japanese character placeholder
		jc.Grid[y][x] = '?'     // First byte placeholder
		jc.Grid[y][x+1] = '?'   // Second byte placeholder
	} else if width == 1 {
		jc.Grid[y][x] = '?'     // Single byte placeholder
	}
}

// DrawJapaneseBox draws a box with Japanese title support
func (jc *JapaneseCanvas) DrawJapaneseBox(x1, y1, x2, y2 int, title string) {
	// Draw the basic box structure using ASCII
	jc.DrawBox(x1, y1, x2, y2)
	
	// Add title if present
	if title != "" {
		titleWidth := jc.widthCalc.CalculateDisplayWidth(title)
		boxWidth := x2 - x1 + 1
		
		// Calculate title position (centered)
		titleStart := x1 + 2
		if titleWidth+4 < boxWidth { // +4 for "[ ]" brackets and spaces
			titleStart = x1 + (boxWidth-titleWidth-4)/2
		}
		
		// Write title with brackets
		if titleStart >= x1 && titleStart < x2 {
			jc.WriteBytes(titleStart, y1, "[ ")
			jc.WriteJapaneseText(titleStart+2, y1, title)
			// Calculate end position accounting for Japanese character widths
			endPos := titleStart + 2 + titleWidth
			if endPos < x2 {
				jc.WriteBytes(endPos, y1, " ]")
			}
		}
	}
}

// CreateJapaneseTable creates a table with Japanese text support
func (jc *JapaneseCanvas) CreateJapaneseTable(x, y int, headers []string, rows [][]string, columnWidths []int) {
	if len(headers) == 0 {
		return
	}
	
	// Calculate column widths if not provided
	if columnWidths == nil {
		columnWidths = jc.calculateOptimalColumnWidths(headers, rows)
	}
	
	currentY := y
	
	// Draw top border
	jc.drawTableBorder(x, currentY, columnWidths, true)
	currentY++
	
	// Draw headers
	jc.drawTableRow(x, currentY, headers, columnWidths)
	currentY++
	
	// Draw separator
	jc.drawTableBorder(x, currentY, columnWidths, false)
	currentY++
	
	// Draw data rows
	for _, row := range rows {
		jc.drawTableRow(x, currentY, row, columnWidths)
		currentY++
	}
	
	// Draw bottom border
	jc.drawTableBorder(x, currentY, columnWidths, true)
}

// calculateOptimalColumnWidths calculates column widths based on content
func (jc *JapaneseCanvas) calculateOptimalColumnWidths(headers []string, rows [][]string) []int {
	if len(headers) == 0 {
		return nil
	}
	
	widths := make([]int, len(headers))
	
	// Initialize with header widths
	for i, header := range headers {
		widths[i] = jc.widthCalc.CalculateDisplayWidth(header) + 2 // +2 for padding
	}
	
	// Check data rows
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				cellWidth := jc.widthCalc.CalculateDisplayWidth(cell) + 2
				if cellWidth > widths[i] {
					widths[i] = cellWidth
				}
			}
		}
	}
	
	return widths
}

// drawTableBorder draws horizontal table borders
func (jc *JapaneseCanvas) drawTableBorder(x, y int, columnWidths []int, isTopBottom bool) {
	currentX := x
	
	// Draw corner/junction
	jc.SetByteAt(currentX, y, '+')
	currentX++
	
	for _, width := range columnWidths {
		// Draw horizontal line
		for j := 0; j < width; j++ {
			jc.SetByteAt(currentX, y, '-')
			currentX++
		}
		
		// Draw corner/junction
		jc.SetByteAt(currentX, y, '+')
		currentX++
	}
}

// drawTableRow draws a table row with Japanese text support
func (jc *JapaneseCanvas) drawTableRow(x, y int, cells []string, columnWidths []int) {
	currentX := x
	
	// Left border
	jc.SetByteAt(currentX, y, '|')
	currentX++
	
	for i, width := range columnWidths {
		var cell string
		if i < len(cells) {
			cell = cells[i]
		}
		
		// Truncate cell if too long
		cell = jc.widthCalc.TruncateToWidth(cell, width-2) // -2 for padding
		
		// Write cell content with padding
		jc.WriteBytes(currentX, y, " ")
		currentX++
		
		jc.WriteJapaneseText(currentX, y, cell)
		cellDisplayWidth := jc.widthCalc.CalculateDisplayWidth(cell)
		currentX += cellDisplayWidth
		
		// Pad remaining space
		remaining := width - cellDisplayWidth - 2
		for j := 0; j < remaining; j++ {
			jc.SetByteAt(currentX, y, ' ')
			currentX++
		}
		
		jc.WriteBytes(currentX, y, " ")
		currentX++
		
		// Right border
		jc.SetByteAt(currentX, y, '|')
		currentX++
	}
}