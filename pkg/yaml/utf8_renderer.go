package yaml

import (
	"fmt"
	"kit4ai/pkg/canvas"
)

// UTF8Renderer handles rendering with full UTF-8 support
type UTF8Renderer struct {
	canvas *canvas.UTF8Canvas
}

// NewUTF8Renderer creates a new UTF-8 renderer
func NewUTF8Renderer(width, height int) *UTF8Renderer {
	return &UTF8Renderer{
		canvas: canvas.NewUTF8CanvasWithSize(width, height),
	}
}

// RenderElements renders elements with full UTF-8 support
func (ur *UTF8Renderer) RenderElements(elements []Element) error {
	for _, elem := range elements {
		if err := ur.renderElement(elem); err != nil {
			return err
		}
	}
	return nil
}

// renderElement renders a single element
func (ur *UTF8Renderer) renderElement(elem Element) error {
	if elem.Box != nil {
		return ur.renderBox(elem.Box)
	}
	if elem.Text != nil {
		return ur.renderText(elem.Text)
	}
	if elem.Line != nil {
		return ur.renderLine(elem.Line)
	}
	if elem.Table != nil {
		return ur.renderTable(elem.Table)
	}
	return nil
}

// renderBox renders a box with UTF-8 title support
func (ur *UTF8Renderer) renderBox(box *BoxElement) error {
	// Draw the box structure
	ur.canvas.DrawBox(box.Position.X, box.Position.Y,
		box.Position.X+box.Size.Width-1,
		box.Position.Y+box.Size.Height-1)
	
	// Add title if present
	if box.Title != "" {
		title := fmt.Sprintf("[ %s ]", box.Title)
		titleX := box.Position.X + 2
		titleY := box.Position.Y
		ur.canvas.WriteText(titleX, titleY, title)
	}
	
	return nil
}

// renderText renders UTF-8 text
func (ur *UTF8Renderer) renderText(text *TextElement) error {
	ur.canvas.WriteText(text.Position.X, text.Position.Y, text.Content)
	return nil
}

// renderLine renders a line
func (ur *UTF8Renderer) renderLine(line *LineElement) error {
	ur.canvas.DrawLine(line.Start.X, line.Start.Y, line.End.X, line.End.Y)
	return nil
}

// renderTable renders a table with UTF-8 content
func (ur *UTF8Renderer) renderTable(table *TableElement) error {
	ur.canvas.CreateTable(table.Position.X, table.Position.Y, table.Headers, table.Rows)
	return nil
}

// String returns the rendered canvas as string
func (ur *UTF8Renderer) String() string {
	return ur.canvas.String()
}