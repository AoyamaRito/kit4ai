# Kit4AI - ASCII Art UI Specification Tool

**[🇯🇵 日本語版 README](README.ja.md) | 🇺🇸 English**

A powerful Go-based tool for creating perfectly aligned ASCII art UI specifications that AI can use to generate web interfaces.

## Demo - Live Output Examples

### Enterprise Dashboard UI

```
+------------------------------------------------------------------------------+
| ENTERPRISE CONTROL PANEL v2.4.1                           2024-06-15 14:32:17|
+------------------------------------------------------------------------------+
| [F1]File [F2]Edit [F3]View [F4]Tools [F5]Reports [F6]Admin [ESC]Exit         |
+------------------------------------------------------------------------------+
+------------------++----------------------------------------------------------+
|NAVIGATION        ||+----------------+ +-----------------+ +-----------------+|
| [1] Dashboard    |||REVENUE METRICS | |PERFORMANCE      | |SECURITY         ||
|>[2] Analytics    |||Daily: $47,892  | |Avg Resp: 245ms  | |Threats: 0       ||
| [3] Users        |||Weekly: $312,456| |Uptime: 99.97%   | |Blocked: 127     ||
| [4] Settings     |||Monthly: $1.2M  | |Errors: 0.03%    | |Firewall: ON     ||
|                  |||Growth: +12.5%  | |Requests: 847K   | |SSL: Valid       ||
|QUICK ACTIONS     |||Target: 87%     | |Cache Hit: 94%   | |Backup: OK       ||
| [R] Refresh      ||+----------------+ +-----------------+ +-----------------+|
| [B] Backup       ||+------------------------------------+ +-----------------+|
| [M] Maintenance  |||ANALYTICS CHART - Last 7 Days       | |LIVE ACTIVITY    ||
|                  |||Revenue |#######*****:::::....      | |14:32 Login: admin|
|SYSTEM STATUS     |||Traffic |****#######****::....      | |14:31 Order #4891||
| CPU: 67%         |||Users   |:::*****########**..       | |14:30 User signup||
| RAM: 4.2/8GB     |||Errors  |.....::::*****......       | |14:29 Payment OK ||
| Online: 1,247    ||+        Mon Tue Wed Thu Fri Sat Sun-+ |14:28 Backup done||
+------------------++----------------------------------------------------------+
| Connected: DB-MAIN | Cache: REDIS-01 | Queue: 247 | Alerts: 0  F10:Settings  |
+------------------------------------------------------------------------------+
```

### Smartphone UI

```
+----------------------------------------------------------+
| 12:34 PM               5G                     [====] 87% |
+----------------------------------------------------------+
| <-                   MESSAGES                       [+]  |
+----------------------------------------------------------+
| [1] New Message               Active: 3                  |
| [2] Contacts                  Total: 127                 |
| [3] Recent Chats              Unread: 5                  |
| [4] Settings                  Status: Online             |
+----------------------------------------------------------+
| John Doe                               2:30 PM           |
| Hey, are you free for lunch?                             |
| Jane Smith                             1:45 PM           |
| Meeting at 3 PM confirmed                                |
+----------------------------------------------------------+
|    [HOME]     [CHAT]     [CALL]     [MORE]     [USER]    |
+----------------------------------------------------------+
```

### ASCII Filter Demo

Input: `"Hello 世界 World! メッセージ Test"`  
Output: `"Hello  World!  Test"`

✅ **Perfect alignment guaranteed** - All full-width characters automatically removed

## Table of Contents

- [Demo - Live Output Examples](#demo---live-output-examples)
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [Examples](#examples)
- [Best Practices](#best-practices)
- [Contributing](#contributing)
- [Language / 言語](#language--言語)

## Overview

Kit4AI solves the layout misalignment problem that occurs when AI directly creates ASCII art. By providing a structured canvas system with automatic character filtering, developers can create consistent, professional UI specifications that render perfectly in Markdown documents.

## Key Features

- **Perfect Alignment**: ByteCanvas system ensures no layout drift
- **Multi-Width Support**: Configurable canvas widths (60, 72, 80, 100, 120 characters)
- **ASCII Filter**: Automatically removes full-width characters to prevent misalignment
- **Layer System**: Z-ordered layers for complex UI composition
- **Markdown Ready**: Output designed for embedding in documentation

## Architecture

### Core Components

- **Canvas**: Basic rune-based grid system
- **ByteCanvas**: 8-bit processing for stable ASCII art
- **TextLayer**: Full-width character support (deprecated for alignment)
- **LayerSystem**: Multi-layer composition with Z-ordering
- **Config System**: Flexible width configurations

### Canvas Configurations

```go
StandardConfig    = 80x100   // Legacy compatible
WideConfig        = 100x100  // Modern displays  
UltraWideConfig   = 120x100  // Large monitors
CompactConfig     = 60x80    // Mobile/narrow
PrintConfig       = 72x90    // A4 paper friendly
```

## Installation

### Prerequisites

- Go 1.19 or later
- Git

### Step 1: Clone the Repository

```bash
git clone https://github.com/your-username/kit4ai.git
cd kit4ai
```

### Step 2: Initialize Go Module

```bash
go mod init kit4ai
go mod tidy
```

### Step 3: Verify Installation

```bash
go run main.go
```

This should create a complex enterprise UI and output:
```
Complex Enterprise Dashboard UI created: complex_enterprise_ui.txt
```

### Step 4: Run Tests (Optional)

```bash
go test ./pkg/canvas/...
```

### Alternative: Direct Download

If you prefer not to use Git:

1. Download the source code as ZIP
2. Extract to your desired directory
3. Follow steps 2-4 above

### Project Structure After Installation

```
kit4ai/
├── go.mod              # Go module file
├── main.go             # Example implementations
├── pkg/
│   └── canvas/
│       ├── canvas.go
│       ├── bytecanvas.go
│       ├── textlayer.go
│       ├── layer.go
│       └── config.go
├── README.md
└── generated files:
    ├── complex_enterprise_ui.txt
    ├── ascii_filter_demo.txt
    └── other example outputs
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "kit4ai/pkg/canvas"
)

func main() {
    // Set configuration
    canvas.SetConfig(canvas.StandardConfig)
    
    // Create canvas
    ui := canvas.NewByteCanvas()
    
    // Draw frames
    ui.DrawBox(0, 0, 79, 10)
    
    // Add text (automatically filters full-width characters)
    ui.WriteBytesASCII(2, 2, "Hello World!")
    
    // Output
    fmt.Println(ui.String())
}
```

### Complex UI Example

```go
// Create enterprise dashboard
canvas.SetConfig(canvas.StandardConfig)
dashboard := canvas.NewByteCanvas()

// Title bar
dashboard.DrawBox(0, 0, 79, 2)
dashboard.WriteBytesASCII(2, 1, "ENTERPRISE DASHBOARD v2.1")

// Multi-panel layout
dashboard.DrawBox(0, 3, 25, 15)  // Sidebar
dashboard.DrawBox(26, 3, 79, 15) // Main content

// Add content with automatic ASCII filtering
dashboard.WriteBytesASCII(2, 5, "Navigation Menu")
dashboard.WriteBytesASCII(28, 5, "Analytics Data")
```

## API Reference

### ByteCanvas Methods

- `NewByteCanvas()` - Create new canvas with current config
- `DrawBox(x1, y1, x2, y2)` - Draw rectangular frame
- `WriteBytes(x, y, text)` - Write raw text
- `WriteBytesASCII(x, y, text)` - Write with full-width character filtering
- `FilterASCII(text)` - Remove full-width characters from string
- `String()` - Convert to string with trailing line removal

### Configuration Methods

- `SetConfig(config)` - Set canvas dimensions
- `SetStandardWidth()` - 80 characters (legacy)
- `SetWideWidth()` - 100 characters (modern)
- `SetCompactWidth()` - 60 characters (mobile)
- `GetCurrentWidth()` - Get active width
- `GetConfigName()` - Get configuration description

## Problem Solved

### Before Kit4AI
```
+------------------+
| 不整列なUI    |  <- Misaligned due to full-width chars
| レイアウト      |
+------------------+
```

### After Kit4AI
```
+------------------+
| Perfect Layout   |  <- Perfect alignment with ASCII filter
| Clean Design     |
+------------------+
```

## Examples

The repository includes several example UIs:

- **Enterprise Dashboard**: Complex multi-panel admin interface
- **Smartphone UI**: Mobile app layout with navigation
- **Banking App**: Financial interface with security features
- **POS Terminal**: Retail point-of-sale system
- **Hospital System**: Medical management interface

## Generated UI Specifications

All examples generate markdown-compatible text files:

```
Configuration: Standard (80x100) - Legacy Compatible
Features: Multi-panel layout, real-time data, charts, logs
ASCII Filter: Enabled (all full-width characters removed)

Layout:
+------------------------------------------------------------------------------+
| ENTERPRISE CONTROL PANEL v2.4.1                           2024-06-15 14:32:17|
+------------------------------------------------------------------------------+
```

## Technical Decisions

### Why ByteCanvas?
- Eliminates Unicode alignment issues
- Consistent 8-bit character processing  
- Stable positioning across all environments

### Why ASCII-Only?
- Universal compatibility
- Prevents layout drift in Markdown
- Consistent rendering in all text editors
- Professional appearance

### Why Layer System?
- Complex UI composition
- Z-order management
- Modular development
- Easy testing of individual components

## File Structure

```
kit4ai/
├── pkg/canvas/
│   ├── canvas.go      # Basic rune canvas
│   ├── bytecanvas.go  # ASCII-optimized canvas  
│   ├── textlayer.go   # Full-width text support
│   ├── layer.go       # Layer composition system
│   └── config.go      # Configuration management
├── main.go            # Example implementations
├── *.txt             # Generated UI specifications
└── README.md         # This file
```

## Best Practices

1. **Always use WriteBytesASCII()** for text content
2. **Set configuration before creating canvas**
3. **Use appropriate width for target display**
4. **Test with different configurations**
5. **Keep UI elements within canvas bounds**

## ASCII Filter Details

The automatic ASCII filter removes:
- Japanese characters (Hiragana, Katakana, Kanji)
- Full-width Unicode characters (0xFF01-0xFF5E)
- Unicode punctuation (0x3000-0x303F)
- Any character > 127 (non-ASCII)

Preserves:
- Standard ASCII (0-127)
- Numbers, letters, symbols
- Box drawing characters (for frames)

## Contributing

1. Fork the repository
2. Create feature branch
3. Add tests for new functionality
4. Ensure ASCII filter compatibility
5. Update documentation
6. Submit pull request

## License

MIT License - see LICENSE file for details

## Use Cases

- **AI UI Generation**: Provide structured templates for AI systems
- **Documentation**: Embed UI mockups in technical docs
- **Prototyping**: Rapid ASCII-based interface design
- **Cross-platform**: Universal text-based UI specifications
- **Legacy Systems**: Terminal-based interface design

---

**Language / 言語:**  
🇺🇸 **English** | [🇯🇵 日本語版 README](README.ja.md)

*Kit4AI enables AI systems to create perfectly aligned ASCII art UI specifications for web development projects.*