package yaml

// UISpec represents the root structure of a YAML UI specification
type UISpec struct {
	Canvas   CanvasSpec   `yaml:"canvas"`
	Elements []Element    `yaml:"elements"`
}

// CanvasSpec defines the canvas dimensions
type CanvasSpec struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

// Element represents a UI element with its type and properties
type Element struct {
	Type       string                 `yaml:"type"`
	Properties map[string]interface{} `yaml:"properties"`
	// Direct element type shortcuts
	Box   *BoxElement   `yaml:"box,omitempty"`
	Text  *TextElement  `yaml:"text,omitempty"`
	Line  *LineElement  `yaml:"line,omitempty"`
	Table *TableElement `yaml:"table,omitempty"`
}

// Position represents x,y coordinates
type Position struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}

// Size represents width and height
type Size struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

// BoxElement represents a box with border
type BoxElement struct {
	Position Position `yaml:"position"`
	Size     Size     `yaml:"size"`
	Title    string   `yaml:"title,omitempty"`
	Border   string   `yaml:"border,omitempty"` // single, double, none
}

// TextElement represents a text element
type TextElement struct {
	Position Position `yaml:"position"`
	Content  string   `yaml:"content"`
}

// LineElement represents a line
type LineElement struct {
	Start      Position `yaml:"start"`
	End        Position `yaml:"end"`
	Style      string   `yaml:"style,omitempty"` // horizontal, vertical, dashed
}

// TableElement represents a table
type TableElement struct {
	Position Position     `yaml:"position"`
	Headers  []string     `yaml:"headers"`
	Rows     [][]string   `yaml:"rows"`
	Width    int          `yaml:"width,omitempty"`
}