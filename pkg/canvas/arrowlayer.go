package canvas

import (
	"fmt"
)

// ArrowStyle defines different arrow styles
type ArrowStyle string

const (
	ArrowStyleNormal   ArrowStyle = "---->"   // Normal arrow
	ArrowStyleThick    ArrowStyle = "====>"   // Thick arrow
	ArrowStyleWave     ArrowStyle = "~~~>"    // Wave arrow
	ArrowStyleShort    ArrowStyle = "--->"    // Short arrow
	ArrowStyleDouble   ArrowStyle = "--->>>"  // Double arrow
	ArrowStyleDotted   ArrowStyle = "...>"    // Dotted arrow
)

// Arrow represents a single arrow with start/end points and style
type Arrow struct {
	StartX, StartY int
	EndX, EndY     int
	Style          ArrowStyle
	Label          string
}

// ArrowLayer manages multiple arrows and renders them
type ArrowLayer struct {
	arrows []Arrow
	zOrder int
}

// NewArrowLayer creates a new arrow layer
func NewArrowLayer() *ArrowLayer {
	return &ArrowLayer{
		arrows: make([]Arrow, 0),
		zOrder: 0,
	}
}

// AddArrow adds a new arrow to the layer
func (al *ArrowLayer) AddArrow(startX, startY, endX, endY int, style ArrowStyle) {
	al.arrows = append(al.arrows, Arrow{
		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,
		Style:  style,
	})
}

// AddLabeledArrow adds a new arrow with a label
func (al *ArrowLayer) AddLabeledArrow(startX, startY, endX, endY int, style ArrowStyle, label string) {
	al.arrows = append(al.arrows, Arrow{
		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,
		Style:  style,
		Label:  label,
	})
}

// SetZOrder sets the z-order for layer rendering
func (al *ArrowLayer) SetZOrder(order int) {
	al.zOrder = order
}

// GetZOrder returns the current z-order
func (al *ArrowLayer) GetZOrder() int {
	return al.zOrder
}

// Render draws all arrows onto the canvas
func (al *ArrowLayer) Render(canvas *ByteCanvas) error {
	for _, arrow := range al.arrows {
		err := al.drawArrow(canvas, arrow)
		if err != nil {
			return err
		}
	}
	return nil
}

// drawArrow draws a single arrow on the canvas
func (al *ArrowLayer) drawArrow(canvas *ByteCanvas, arrow Arrow) error {
	// Determine arrow direction and call appropriate drawing function
	if arrow.StartY == arrow.EndY {
		// Horizontal arrow
		return al.drawHorizontalArrow(canvas, arrow)
	} else if arrow.StartX == arrow.EndX {
		// Vertical arrow
		return al.drawVerticalArrow(canvas, arrow)
	} else {
		// Diagonal arrow (L-shaped path)
		return al.drawLShapedArrow(canvas, arrow)
	}
}

// drawHorizontalArrow draws a horizontal arrow
func (al *ArrowLayer) drawHorizontalArrow(canvas *ByteCanvas, arrow Arrow) error {
	startX, endX := arrow.StartX, arrow.EndX
	y := arrow.StartY
	
	// Ensure startX < endX for left-to-right drawing
	if startX > endX {
		startX, endX = endX, startX
	}
	
	// Get line character and arrow head based on style
	lineChar, arrowHead := al.getArrowComponents(arrow.Style)
	
	// Draw the line
	for x := startX; x < endX; x++ {
		if x >= 0 && x < canvas.Width && y >= 0 && y < canvas.Height {
			canvas.Grid[y][x] = lineChar
		}
	}
	
	// Draw arrow head
	if endX >= 0 && endX < canvas.Width && y >= 0 && y < canvas.Height {
		canvas.Grid[y][endX] = arrowHead
	}
	
	// Draw label if provided
	if arrow.Label != "" {
		labelX := (startX + endX) / 2
		labelY := y - 1
		if labelY >= 0 {
			al.drawLabel(canvas, labelX, labelY, arrow.Label)
		}
	}
	
	return nil
}

// drawVerticalArrow draws a vertical arrow
func (al *ArrowLayer) drawVerticalArrow(canvas *ByteCanvas, arrow Arrow) error {
	startY, endY := arrow.StartY, arrow.EndY
	x := arrow.StartX
	
	// Ensure startY < endY for top-to-bottom drawing
	if startY > endY {
		startY, endY = endY, startY
	}
	
	// Get line character for vertical arrows
	var lineChar byte = '|'
	var arrowHead byte = 'v'
	
	if arrow.Style == ArrowStyleThick {
		lineChar = '#'  // Use # for thick vertical lines
		arrowHead = 'V' // Use V for thick arrow head
	}
	
	// Draw the line
	for y := startY; y < endY; y++ {
		if x >= 0 && x < canvas.Width && y >= 0 && y < canvas.Height {
			canvas.Grid[y][x] = lineChar
		}
	}
	
	// Draw arrow head
	if x >= 0 && x < canvas.Width && endY >= 0 && endY < canvas.Height {
		canvas.Grid[endY][x] = arrowHead
	}
	
	// Draw label if provided
	if arrow.Label != "" {
		labelX := x + 2
		labelY := (startY + endY) / 2
		al.drawLabel(canvas, labelX, labelY, arrow.Label)
	}
	
	return nil
}

// drawLShapedArrow draws an L-shaped arrow for diagonal connections
func (al *ArrowLayer) drawLShapedArrow(canvas *ByteCanvas, arrow Arrow) error {
	// Draw horizontal line first, then vertical
	midArrow := Arrow{
		StartX: arrow.StartX,
		StartY: arrow.StartY,
		EndX:   arrow.EndX,
		EndY:   arrow.StartY,
		Style:  arrow.Style,
	}
	
	err := al.drawHorizontalArrow(canvas, midArrow)
	if err != nil {
		return err
	}
	
	// Then vertical line
	midArrow = Arrow{
		StartX: arrow.EndX,
		StartY: arrow.StartY,
		EndX:   arrow.EndX,
		EndY:   arrow.EndY,
		Style:  arrow.Style,
	}
	
	return al.drawVerticalArrow(canvas, midArrow)
}

// getArrowComponents returns the line character and arrow head for a given style
func (al *ArrowLayer) getArrowComponents(style ArrowStyle) (byte, byte) {
	switch style {
	case ArrowStyleNormal:
		return '-', '>'
	case ArrowStyleThick:
		return '=', '>' // Use = for thick horizontal lines
	case ArrowStyleWave:
		return '~', '>'
	case ArrowStyleShort:
		return '-', '>'
	case ArrowStyleDouble:
		return '=', '>' // Use = and > for double arrows
	case ArrowStyleDotted:
		return '.', '>'
	default:
		return '-', '>'
	}
}

// drawLabel draws a text label on the canvas
func (al *ArrowLayer) drawLabel(canvas *ByteCanvas, x, y int, label string) {
	for i, char := range label {
		posX := x + i
		if posX >= 0 && posX < canvas.Width && y >= 0 && y < canvas.Height {
			canvas.Grid[y][posX] = byte(char)
		}
	}
}

// Clear removes all arrows from the layer
func (al *ArrowLayer) Clear() {
	al.arrows = al.arrows[:0]
}

// GetArrowCount returns the number of arrows in the layer
func (al *ArrowLayer) GetArrowCount() int {
	return len(al.arrows)
}

// RemoveArrow removes an arrow at the specified index
func (al *ArrowLayer) RemoveArrow(index int) error {
	if index < 0 || index >= len(al.arrows) {
		return fmt.Errorf("arrow index %d out of range", index)
	}
	
	al.arrows = append(al.arrows[:index], al.arrows[index+1:]...)
	return nil
}