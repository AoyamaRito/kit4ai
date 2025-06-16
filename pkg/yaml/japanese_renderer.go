package yaml

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"strings"
	"unicode/utf8"
)

// JapaneseRenderer handles rendering with Japanese text support
type JapaneseRenderer struct {
	canvas *canvas.ByteCanvas
}

// NewJapaneseRenderer creates a new Japanese renderer
func NewJapaneseRenderer(width, height int) *JapaneseRenderer {
	return &JapaneseRenderer{
		canvas: canvas.NewByteCanvasWithSize(width, height),
	}
}

// RenderElements renders elements with Japanese support
func (jr *JapaneseRenderer) RenderElements(elements []Element) error {
	for _, elem := range elements {
		if err := jr.renderElement(elem); err != nil {
			return err
		}
	}
	return nil
}

// renderElement renders a single element with Japanese support
func (jr *JapaneseRenderer) renderElement(elem Element) error {
	if elem.Box != nil {
		return jr.renderJapaneseBox(elem.Box)
	}
	if elem.Text != nil {
		return jr.renderJapaneseText(elem.Text)
	}
	if elem.Line != nil {
		return jr.renderLine(elem.Line)
	}
	if elem.Table != nil {
		return jr.renderJapaneseTable(elem.Table)
	}
	return nil
}

// renderJapaneseBox renders a box with Japanese title support
func (jr *JapaneseRenderer) renderJapaneseBox(box *BoxElement) error {
	// Draw the ASCII box structure
	jr.canvas.DrawBox(box.Position.X, box.Position.Y,
		box.Position.X+box.Size.Width-1,
		box.Position.Y+box.Size.Height-1)
	
	// Add Japanese title if present
	if box.Title != "" {
		// For now, use ASCII-safe approach: encode Japanese as ASCII-safe representation
		titleSafe := jr.makeASCIISafe(box.Title)
		titleX := box.Position.X + 2
		titleY := box.Position.Y
		jr.canvas.WriteBytesASCII(titleX, titleY, fmt.Sprintf("[ %s ]", titleSafe))
	}
	
	return nil
}

// renderJapaneseText renders Japanese text
func (jr *JapaneseRenderer) renderJapaneseText(text *TextElement) error {
	// Convert Japanese to ASCII-safe representation
	textSafe := jr.makeASCIISafe(text.Content)
	jr.canvas.WriteBytesASCII(text.Position.X, text.Position.Y, textSafe)
	return nil
}

// renderLine renders a line (same as regular)
func (jr *JapaneseRenderer) renderLine(line *LineElement) error {
	if line.Start.Y == line.End.Y {
		// Horizontal line
		char := byte('-')
		for x := line.Start.X; x <= line.End.X; x++ {
			jr.canvas.SetByteAt(x, line.Start.Y, char)
		}
	} else if line.Start.X == line.End.X {
		// Vertical line
		char := byte('|')
		for y := line.Start.Y; y <= line.End.Y; y++ {
			jr.canvas.SetByteAt(line.Start.X, y, char)
		}
	}
	return nil
}

// renderJapaneseTable renders a table with Japanese content
func (jr *JapaneseRenderer) renderJapaneseTable(table *TableElement) error {
	x, y := table.Position.X, table.Position.Y
	
	// Convert headers and data to ASCII-safe
	safeHeaders := make([]string, len(table.Headers))
	for i, header := range table.Headers {
		safeHeaders[i] = jr.makeASCIISafe(header)
	}
	
	safeRows := make([][]string, len(table.Rows))
	for i, row := range table.Rows {
		safeRows[i] = make([]string, len(row))
		for j, cell := range row {
			safeRows[i][j] = jr.makeASCIISafe(cell)
		}
	}
	
	// Calculate column widths
	colWidths := jr.calculateColumnWidths(safeHeaders, safeRows)
	
	// Draw table using ASCII-safe content
	jr.drawTable(x, y, safeHeaders, safeRows, colWidths)
	
	return nil
}

// makeASCIISafe converts Japanese text to ASCII-safe representation
// This is a temporary solution - in production you'd want proper Unicode handling
func (jr *JapaneseRenderer) makeASCIISafe(text string) string {
	if !utf8.ValidString(text) {
		return text // Return as-is if not valid UTF-8
	}
	
	// For now, use romanization/translation hints in comments
	// In a real implementation, you might:
	// 1. Use romanization libraries
	// 2. Use width-aware truncation
	// 3. Use special Unicode handling
	
	result := strings.Builder{}
	for _, r := range text {
		if r < 128 {
			// ASCII character - keep as-is
			result.WriteRune(r)
		} else {
			// Non-ASCII - for demo, replace with placeholder that indicates Japanese
			// In real implementation, you'd do proper character width calculation
			result.WriteString("*") // Placeholder for Japanese characters
		}
	}
	
	return result.String()
}

// calculateColumnWidths calculates optimal column widths
func (jr *JapaneseRenderer) calculateColumnWidths(headers []string, rows [][]string) []int {
	if len(headers) == 0 {
		return nil
	}
	
	widths := make([]int, len(headers))
	
	// Initialize with header widths
	for i, header := range headers {
		widths[i] = len(header) + 2 // +2 for padding
	}
	
	// Check data rows
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) {
				cellWidth := len(cell) + 2
				if cellWidth > widths[i] {
					widths[i] = cellWidth
				}
			}
		}
	}
	
	return widths
}

// drawTable draws the table structure
func (jr *JapaneseRenderer) drawTable(x, y int, headers []string, rows [][]string, colWidths []int) {
	currentY := y
	
	// Top border
	jr.drawTableBorder(x, currentY, colWidths)
	currentY++
	
	// Headers
	jr.drawTableRow(x, currentY, headers, colWidths)
	currentY++
	
	// Separator
	jr.drawTableBorder(x, currentY, colWidths)
	currentY++
	
	// Data rows
	for _, row := range rows {
		jr.drawTableRow(x, currentY, row, colWidths)
		currentY++
	}
	
	// Bottom border
	jr.drawTableBorder(x, currentY, colWidths)
}

// drawTableBorder draws horizontal table border
func (jr *JapaneseRenderer) drawTableBorder(x, y int, colWidths []int) {
	currentX := x
	jr.canvas.SetByteAt(currentX, y, '+')
	currentX++
	
	for _, width := range colWidths {
		for i := 0; i < width; i++ {
			jr.canvas.SetByteAt(currentX, y, '-')
			currentX++
		}
		jr.canvas.SetByteAt(currentX, y, '+')
		currentX++
	}
}

// drawTableRow draws a table row
func (jr *JapaneseRenderer) drawTableRow(x, y int, cells []string, colWidths []int) {
	currentX := x
	jr.canvas.SetByteAt(currentX, y, '|')
	currentX++
	
	for i, width := range colWidths {
		var cell string
		if i < len(cells) {
			cell = cells[i]
		}
		
		jr.canvas.WriteBytesASCII(currentX, y, fmt.Sprintf(" %-*s", width-1, cell))
		currentX += width
		jr.canvas.SetByteAt(currentX, y, '|')
		currentX++
	}
}

// String returns the rendered canvas as string
func (jr *JapaneseRenderer) String() string {
	return jr.canvas.String()
}