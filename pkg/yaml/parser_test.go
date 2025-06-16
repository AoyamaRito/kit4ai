package yaml

import (
	"strings"
	"testing"
)

func TestParseSimpleYAML(t *testing.T) {
	yamlContent := `
canvas:
  width: 40
  height: 10

elements:
  - box:
      position:
        x: 0
        y: 0
      size:
        width: 20
        height: 5
      title: "Test Box"
  - text:
      position:
        x: 2
        y: 2
      content: "Hello World"
`

	parser := NewParser()
	spec, err := parser.Parse(strings.NewReader(yamlContent))
	if err != nil {
		t.Fatalf("Failed to parse YAML: %v", err)
	}

	if spec.Canvas.Width != 40 || spec.Canvas.Height != 10 {
		t.Errorf("Canvas dimensions incorrect: got %dx%d, want 40x10",
			spec.Canvas.Width, spec.Canvas.Height)
	}

	if len(spec.Elements) != 2 {
		t.Errorf("Expected 2 elements, got %d", len(spec.Elements))
	}
}

func TestRenderBox(t *testing.T) {
	yamlContent := `
canvas:
  width: 30
  height: 10

elements:
  - box:
      position: {x: 0, y: 0}
      size: {width: 15, height: 5}
      title: "Test"
`

	parser := NewParser()
	spec, err := parser.Parse(strings.NewReader(yamlContent))
	if err != nil {
		t.Fatalf("Failed to parse YAML: %v", err)
	}

	result, err := parser.Render(spec)
	if err != nil {
		t.Fatalf("Failed to render: %v", err)
	}

	// Check that result contains box characters
	if !strings.Contains(result, "+") || !strings.Contains(result, "-") {
		t.Errorf("Rendered output doesn't contain box characters")
	}
}

func TestRenderTable(t *testing.T) {
	yamlContent := `
canvas:
  width: 50
  height: 20

elements:
  - table:
      position: {x: 0, y: 0}
      headers: ["ID", "Name", "Status"]
      rows:
        - ["001", "Service A", "Active"]
        - ["002", "Service B", "Inactive"]
`

	parser := NewParser()
	spec, err := parser.Parse(strings.NewReader(yamlContent))
	if err != nil {
		t.Fatalf("Failed to parse YAML: %v", err)
	}

	result, err := parser.Render(spec)
	if err != nil {
		t.Fatalf("Failed to render: %v", err)
	}

	// Check that result contains table headers
	if !strings.Contains(result, "ID") || !strings.Contains(result, "Name") {
		t.Errorf("Rendered output doesn't contain table headers")
	}
}