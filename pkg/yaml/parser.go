package yaml

import (
	"fmt"
	"io"
	"kit4ai/pkg/canvas"

	"gopkg.in/yaml.v3"
)

// Parser handles YAML to ASCII art conversion
type Parser struct {
	spec UISpec
}

// NewParser creates a new YAML parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse reads YAML from an io.Reader and returns a UISpec
func (p *Parser) Parse(r io.Reader) (*UISpec, error) {
	decoder := yaml.NewDecoder(r)
	err := decoder.Decode(&p.spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}
	
	// Validate the spec
	if err := p.validate(); err != nil {
		return nil, err
	}
	
	return &p.spec, nil
}

// validate checks if the UISpec is valid
func (p *Parser) validate() error {
	if p.spec.Canvas.Width <= 0 || p.spec.Canvas.Height <= 0 {
		return fmt.Errorf("canvas dimensions must be positive")
	}
	
	if p.spec.Canvas.Width > 200 || p.spec.Canvas.Height > 200 {
		return fmt.Errorf("canvas dimensions too large (max 200x200)")
	}
	
	return nil
}

// Render converts the UISpec to ASCII art
func (p *Parser) Render(spec *UISpec) (string, error) {
	// Use UTF-8 renderer if Japanese mode is enabled
	if spec.Canvas.JapaneseMode {
		ur := NewUTF8Renderer(spec.Canvas.Width, spec.Canvas.Height)
		if err := ur.RenderElements(spec.Elements); err != nil {
			return "", fmt.Errorf("failed to render UTF-8 elements: %w", err)
		}
		return ur.String(), nil
	}
	
	// Use regular byte canvas for ASCII-only mode
	c := canvas.NewByteCanvasWithSize(spec.Canvas.Width, spec.Canvas.Height)
	
	// Process each element
	for _, elem := range spec.Elements {
		if err := p.renderElement(c, elem); err != nil {
			return "", fmt.Errorf("failed to render element: %w", err)
		}
	}
	
	return c.String(), nil
}

// renderElement renders a single element to the canvas
func (p *Parser) renderElement(c *canvas.ByteCanvas, elem Element) error {
	// Check for direct element types first
	if elem.Box != nil {
		return p.renderBox(c, elem.Box)
	}
	if elem.Text != nil {
		return p.renderText(c, elem.Text)
	}
	if elem.Line != nil {
		return p.renderLine(c, elem.Line)
	}
	if elem.Table != nil {
		return p.renderTable(c, elem.Table)
	}
	
	// Fall back to type-based rendering
	switch elem.Type {
	case "box":
		box := &BoxElement{}
		if err := p.mapToStruct(elem.Properties, box); err != nil {
			return err
		}
		return p.renderBox(c, box)
	case "text":
		text := &TextElement{}
		if err := p.mapToStruct(elem.Properties, text); err != nil {
			return err
		}
		return p.renderText(c, text)
	case "line":
		line := &LineElement{}
		if err := p.mapToStruct(elem.Properties, line); err != nil {
			return err
		}
		return p.renderLine(c, line)
	case "table":
		table := &TableElement{}
		if err := p.mapToStruct(elem.Properties, table); err != nil {
			return err
		}
		return p.renderTable(c, table)
	default:
		return fmt.Errorf("unknown element type: %s", elem.Type)
	}
}

// mapToStruct converts a map to a struct (simplified version)
func (p *Parser) mapToStruct(props map[string]interface{}, target interface{}) error {
	// This is a simplified implementation
	// In production, you'd use a proper mapping library
	return nil
}

// renderBox renders a box element
func (p *Parser) renderBox(c *canvas.ByteCanvas, box *BoxElement) error {
	// Draw the box
	c.DrawBox(box.Position.X, box.Position.Y, 
		box.Position.X+box.Size.Width-1, 
		box.Position.Y+box.Size.Height-1)
	
	// Add title if present
	if box.Title != "" {
		titleX := box.Position.X + 2
		titleY := box.Position.Y
		c.WriteBytesASCII(titleX, titleY, fmt.Sprintf("[ %s ]", box.Title))
	}
	
	return nil
}

// renderText renders a text element
func (p *Parser) renderText(c *canvas.ByteCanvas, text *TextElement) error {
	c.WriteBytesASCII(text.Position.X, text.Position.Y, text.Content)
	return nil
}

// renderLine renders a line element
func (p *Parser) renderLine(c *canvas.ByteCanvas, line *LineElement) error {
	// Determine if horizontal or vertical
	if line.Start.Y == line.End.Y {
		// Horizontal line
		char := byte('-')
		if line.Style == "double" {
			char = '='
		}
		for x := line.Start.X; x <= line.End.X; x++ {
			c.SetByteAt(x, line.Start.Y, char)
		}
	} else if line.Start.X == line.End.X {
		// Vertical line
		char := byte('|')
		for y := line.Start.Y; y <= line.End.Y; y++ {
			c.SetByteAt(line.Start.X, y, char)
		}
	} else {
		return fmt.Errorf("diagonal lines not supported yet")
	}
	
	return nil
}

// renderTable renders a table element
func (p *Parser) renderTable(c *canvas.ByteCanvas, table *TableElement) error {
	x, y := table.Position.X, table.Position.Y
	
	// Calculate column widths
	colWidths := make([]int, len(table.Headers))
	for i, header := range table.Headers {
		colWidths[i] = len(header) + 2 // padding
	}
	
	for _, row := range table.Rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell)+2 > colWidths[i] {
				colWidths[i] = len(cell) + 2
			}
		}
	}
	
	// Draw table border top
	currentX := x
	c.SetByteAt(currentX, y, '+')
	currentX++
	for _, width := range colWidths {
		for i := 0; i < width; i++ {
			c.SetByteAt(currentX, y, '-')
			currentX++
		}
		c.SetByteAt(currentX, y, '+')
		currentX++
	}
	y++
	
	// Draw headers
	currentX = x
	c.SetByteAt(currentX, y, '|')
	currentX++
	for i, header := range table.Headers {
		c.WriteBytesASCII(currentX, y, fmt.Sprintf(" %-*s", colWidths[i]-1, header))
		currentX += colWidths[i]
		c.SetByteAt(currentX, y, '|')
		currentX++
	}
	y++
	
	// Draw separator
	currentX = x
	c.SetByteAt(currentX, y, '+')
	currentX++
	for _, width := range colWidths {
		for i := 0; i < width; i++ {
			c.SetByteAt(currentX, y, '-')
			currentX++
		}
		c.SetByteAt(currentX, y, '+')
		currentX++
	}
	y++
	
	// Draw rows
	for _, row := range table.Rows {
		currentX = x
		c.SetByteAt(currentX, y, '|')
		currentX++
		for i, cell := range row {
			if i < len(colWidths) {
				c.WriteBytesASCII(currentX, y, fmt.Sprintf(" %-*s", colWidths[i]-1, cell))
				currentX += colWidths[i]
				c.SetByteAt(currentX, y, '|')
				currentX++
			}
		}
		y++
	}
	
	// Draw bottom border
	currentX = x
	c.SetByteAt(currentX, y, '+')
	currentX++
	for _, width := range colWidths {
		for i := 0; i < width; i++ {
			c.SetByteAt(currentX, y, '-')
			currentX++
		}
		c.SetByteAt(currentX, y, '+')
		currentX++
	}
	
	return nil
}