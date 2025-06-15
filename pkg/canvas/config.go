package canvas

import "fmt"

// CanvasConfig holds configuration for canvas dimensions
type CanvasConfig struct {
	Width  int
	Height int
}

// Predefined configurations
var (
	// Standard 80-column configuration (legacy compatible)
	StandardConfig = CanvasConfig{
		Width:  80,
		Height: 100,
	}
	
	// Wide 100-column configuration (modern displays)
	WideConfig = CanvasConfig{
		Width:  100,
		Height: 100,
	}
	
	// Ultra-wide 120-column configuration (large monitors)
	UltraWideConfig = CanvasConfig{
		Width:  120,
		Height: 100,
	}
	
	// Compact 60-column configuration (mobile/narrow displays)
	CompactConfig = CanvasConfig{
		Width:  60,
		Height: 80,
	}
	
	// Print-friendly A4 configuration
	PrintConfig = CanvasConfig{
		Width:  72,  // Fits A4 paper with margins
		Height: 90,
	}
)

// CurrentConfig holds the active configuration
var CurrentConfig = StandardConfig

// SetConfig changes the global canvas configuration
func SetConfig(config CanvasConfig) {
	CurrentConfig = config
}

// SetStandardWidth sets 80-character width (legacy compatible)
func SetStandardWidth() {
	SetConfig(StandardConfig)
}

// SetWideWidth sets 100-character width (modern displays)
func SetWideWidth() {
	SetConfig(WideConfig)
}

// SetUltraWideWidth sets 120-character width (large monitors)
func SetUltraWideWidth() {
	SetConfig(UltraWideConfig)
}

// SetCompactWidth sets 60-character width (mobile/narrow)
func SetCompactWidth() {
	SetConfig(CompactConfig)
}

// SetPrintWidth sets 72-character width (print-friendly)
func SetPrintWidth() {
	SetConfig(PrintConfig)
}

// SetCustomConfig allows setting custom dimensions
func SetCustomConfig(width, height int) {
	SetConfig(CanvasConfig{
		Width:  width,
		Height: height,
	})
}

// GetCurrentWidth returns the current configured width
func GetCurrentWidth() int {
	return CurrentConfig.Width
}

// GetCurrentHeight returns the current configured height
func GetCurrentHeight() int {
	return CurrentConfig.Height
}

// GetConfigName returns a description of the current configuration
func GetConfigName() string {
	switch CurrentConfig {
	case StandardConfig:
		return "Standard (80x100) - Legacy Compatible"
	case WideConfig:
		return "Wide (100x100) - Modern Display"
	case UltraWideConfig:
		return "Ultra-Wide (120x100) - Large Monitor"
	case CompactConfig:
		return "Compact (60x80) - Mobile/Narrow"
	case PrintConfig:
		return "Print (72x90) - A4 Paper Friendly"
	default:
		return fmt.Sprintf("Custom (%dx%d)", CurrentConfig.Width, CurrentConfig.Height)
	}
}