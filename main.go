package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"kit4ai/pkg/canvas"
	"kit4ai/pkg/yaml"
)

func main() {
	// Define command-line flags
	var (
		template = flag.String("template", "enterprise", "UI template to generate (enterprise, mobile, simple)")
		width    = flag.Int("width", 80, "Canvas width (60, 72, 80, 100, 120)")
		output   = flag.String("output", "", "Output file name (default: auto-generated)")
		insert   = flag.String("insert", "", "Insert UI into existing file at specified line (format: file:line)")
		backup   = flag.Bool("backup", false, "Create backup when inserting into existing file")
		help     = flag.Bool("help", false, "Show help information")
		version  = flag.Bool("version", false, "Show version information")
		yamlFile = flag.String("yaml", "", "YAML file to parse (use '-' for stdin)")
		jp       = flag.Bool("jp", false, "Enable Japanese mode for YAML processing")
	)
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Kit4AI - ASCII Art UI Specification Tool\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nTemplates:\n")
		fmt.Fprintf(os.Stderr, "  enterprise  Complex dashboard UI (default)\n")
		fmt.Fprintf(os.Stderr, "  mobile      Smartphone interface\n")
		fmt.Fprintf(os.Stderr, "  simple      Basic box layout\n")
		fmt.Fprintf(os.Stderr, "\nWidth Options:\n")
		fmt.Fprintf(os.Stderr, "  60   Compact (mobile/narrow)\n")
		fmt.Fprintf(os.Stderr, "  72   Print-friendly (A4)\n")
		fmt.Fprintf(os.Stderr, "  80   Standard (legacy compatible)\n")
		fmt.Fprintf(os.Stderr, "  100  Wide (modern displays)\n")
		fmt.Fprintf(os.Stderr, "  120  Ultra-wide (large monitors)\n")
		fmt.Fprintf(os.Stderr, "\nInsertion:\n")
		fmt.Fprintf(os.Stderr, "  --insert file:line   Insert UI into existing file at specified line number\n")
		fmt.Fprintf(os.Stderr, "  --backup             Create backup (.bak) before inserting\n")
		fmt.Fprintf(os.Stderr, "\nJapanese Support:\n")
		fmt.Fprintf(os.Stderr, "  --jp                 Enable Japanese mode for YAML processing\n")
		fmt.Fprintf(os.Stderr, "                       (Use with --yaml for Japanese text support)\n")
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s --template=mobile --width=60\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --template=enterprise --width=100 --output=dashboard.txt\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --template=simple --insert=document.txt:10 --backup\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --yaml=file.yaml --jp --output=japanese_ui.txt\n", os.Args[0])
	}
	
	flag.Parse()
	
	if *help {
		flag.Usage()
		return
	}
	
	if *version {
		fmt.Println("Kit4AI v1.0.0 - ASCII Art UI Specification Tool")
		return
	}
	
	// Check if YAML mode
	if *yamlFile != "" {
		if err := processYAML(*yamlFile, *output, *insert, *backup, *jp); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}
	
	// Set canvas configuration based on width
	switch *width {
	case 60:
		canvas.SetConfig(canvas.CompactConfig)
	case 72:
		canvas.SetConfig(canvas.PrintConfig)
	case 80:
		canvas.SetConfig(canvas.StandardConfig)
	case 100:
		canvas.SetConfig(canvas.WideConfig)
	case 120:
		canvas.SetConfig(canvas.UltraWideConfig)
	default:
		fmt.Fprintf(os.Stderr, "Error: Unsupported width %d. Use 60, 72, 80, 100, or 120.\n", *width)
		os.Exit(1)
	}
	
	// Check for insert mode
	var insertFile string
	var insertLine int
	if *insert != "" {
		parts := strings.Split(*insert, ":")
		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "Error: Insert format must be 'file:line' (e.g., document.txt:10)\n")
			os.Exit(1)
		}
		insertFile = parts[0]
		var err error
		insertLine, err = strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid line number '%s'\n", parts[1])
			os.Exit(1)
		}
		if insertLine < 1 {
			fmt.Fprintf(os.Stderr, "Error: Line number must be >= 1\n")
			os.Exit(1)
		}
	}
	
	// Generate UI based on template
	var uiContent string
	switch *template {
	case "enterprise":
		uiContent = generateEnterpriseUI(*output, insertFile != "")
	case "mobile":
		uiContent = generateMobileUI(*output, insertFile != "")
	case "simple":
		uiContent = generateSimpleUI(*output, insertFile != "")
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown template '%s'. Use enterprise, mobile, or simple.\n", *template)
		os.Exit(1)
	}
	
	// Handle insertion if requested
	if insertFile != "" {
		err := insertIntoFile(insertFile, insertLine, uiContent, *backup)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error inserting into file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("UI inserted into %s at line %d\n", insertFile, insertLine)
	}
}

func insertIntoFile(filename string, lineNum int, content string, createBackup bool) error {
	// Read existing file
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	
	// Create backup if requested
	if createBackup {
		backupName := filename + ".bak"
		err := copyFile(filename, backupName)
		if err != nil {
			return fmt.Errorf("failed to create backup: %v", err)
		}
		fmt.Printf("Backup created: %s\n", backupName)
	}
	
	// Insert content at specified line
	insertPos := lineNum - 1 // Convert to 0-based index
	if insertPos > len(lines) {
		insertPos = len(lines) // Append at end if line number is beyond file
	}
	
	// Split content into lines
	contentLines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	
	// Create new line slice with inserted content
	newLines := make([]string, 0, len(lines)+len(contentLines))
	newLines = append(newLines, lines[:insertPos]...)
	newLines = append(newLines, contentLines...)
	newLines = append(newLines, lines[insertPos:]...)
	
	// Write back to file
	err = os.WriteFile(filename, []byte(strings.Join(newLines, "\n")+"\n"), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	
	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

func generateEnterpriseUI(outputFile string, insertMode bool) string {
	// Create complex enterprise dashboard UI
	mainCanvas := canvas.NewByteCanvas()
	
	currentWidth := canvas.GetCurrentWidth()
	maxX := currentWidth - 1
	
	// Title bar
	mainCanvas.DrawBox(0, 0, maxX, 2)
	
	// Menu bar
	mainCanvas.DrawBox(0, 3, maxX, 5)
	
	// Left sidebar (navigation) - proportional width
	sidebarWidth := currentWidth / 4
	mainCanvas.DrawBox(0, 6, sidebarWidth, 35)
	
	// Main content area
	mainCanvas.DrawBox(sidebarWidth+1, 6, maxX, 25)
	
	// Statistics panels (3 columns) - responsive layout
	panelWidth := (maxX - sidebarWidth - 4) / 3
	panel1Start := sidebarWidth + 2
	panel2Start := panel1Start + panelWidth + 2
	panel3Start := panel2Start + panelWidth + 2
	
	mainCanvas.DrawBox(panel1Start, 7, panel1Start+panelWidth, 15)
	if panel2Start+panelWidth <= maxX {
		mainCanvas.DrawBox(panel2Start, 7, panel2Start+panelWidth, 15)
	}
	if panel3Start+panelWidth <= maxX {
		mainCanvas.DrawBox(panel3Start, 7, panel3Start+panelWidth, 15)
	}
	
	// Chart area
	chartEnd := maxX - 20
	if chartEnd > panel1Start+panelWidth*2 {
		mainCanvas.DrawBox(panel1Start, 16, chartEnd, 24)
		
		// Live feed
		mainCanvas.DrawBox(chartEnd+1, 16, maxX-1, 24)
	}
	
	// Log/activity area
	mainCanvas.DrawBox(sidebarWidth+1, 26, maxX, 35)
	
	// Status bar
	mainCanvas.DrawBox(0, 36, maxX, 38)
	
	// Title bar content
	mainCanvas.WriteBytesASCII(2, 1, "ENTERPRISE CONTROL PANEL v2.4.1")
	if currentWidth >= 80 {
		mainCanvas.WriteBytesASCII(currentWidth-18, 1, "2024-06-15 14:32:17")
	}
	
	// Menu items - adjust for width
	if currentWidth >= 100 {
		mainCanvas.WriteBytesASCII(2, 4, "[F1]File [F2]Edit [F3]View [F4]Tools [F5]Reports [F6]Admin [ESC]Exit")
	} else if currentWidth >= 80 {
		mainCanvas.WriteBytesASCII(2, 4, "[F1]File [F2]Edit [F3]View [F4]Tools [ESC]Exit")
	} else {
		mainCanvas.WriteBytesASCII(2, 4, "[F1]File [F2]Edit [ESC]Exit")
	}
	
	// Left sidebar navigation
	mainCanvas.WriteBytesASCII(1, 7, "NAVIGATION")
	mainCanvas.WriteBytesASCII(1, 9, " [1] Dashboard")
	mainCanvas.WriteBytesASCII(1, 10, ">[2] Analytics")
	mainCanvas.WriteBytesASCII(1, 11, " [3] Users")
	mainCanvas.WriteBytesASCII(1, 12, " [4] Settings")
	mainCanvas.WriteBytesASCII(1, 13, " [5] Reports")
	mainCanvas.WriteBytesASCII(1, 14, " [6] Logs")
	mainCanvas.WriteBytesASCII(1, 15, " [7] System")
	mainCanvas.WriteBytesASCII(1, 16, " [8] Security")
	
	mainCanvas.WriteBytesASCII(1, 18, "QUICK ACTIONS")
	mainCanvas.WriteBytesASCII(1, 19, " [R] Refresh")
	mainCanvas.WriteBytesASCII(1, 20, " [B] Backup")
	mainCanvas.WriteBytesASCII(1, 21, " [M] Maintenance")
	mainCanvas.WriteBytesASCII(1, 22, " [A] Alerts")
	
	mainCanvas.WriteBytesASCII(1, 24, "SYSTEM STATUS")
	mainCanvas.WriteBytesASCII(1, 25, " CPU: 67%")
	mainCanvas.WriteBytesASCII(1, 26, " RAM: 4.2/8GB")
	mainCanvas.WriteBytesASCII(1, 27, " Disk: 234/500GB")
	mainCanvas.WriteBytesASCII(1, 28, " Net: 89 Mbps")
	
	mainCanvas.WriteBytesASCII(1, 30, "ACTIVE USERS")
	mainCanvas.WriteBytesASCII(1, 31, " Online: 1,247")
	mainCanvas.WriteBytesASCII(1, 32, " Peak: 1,892")
	mainCanvas.WriteBytesASCII(1, 33, " Sessions: 3,456")
	
	// Statistics panels - responsive positioning
	mainCanvas.WriteBytesASCII(panel1Start+1, 8, "REVENUE METRICS")
	mainCanvas.WriteBytesASCII(panel1Start+1, 10, "Daily: $47,892")
	mainCanvas.WriteBytesASCII(panel1Start+1, 11, "Weekly: $312,456")
	mainCanvas.WriteBytesASCII(panel1Start+1, 12, "Monthly: $1.2M")
	mainCanvas.WriteBytesASCII(panel1Start+1, 13, "Growth: +12.5%")
	mainCanvas.WriteBytesASCII(panel1Start+1, 14, "Target: 87%")
	
	if panel2Start+panelWidth <= maxX {
		mainCanvas.WriteBytesASCII(panel2Start+1, 8, "PERFORMANCE")
		mainCanvas.WriteBytesASCII(panel2Start+1, 10, "Avg Resp: 245ms")
		mainCanvas.WriteBytesASCII(panel2Start+1, 11, "Uptime: 99.97%")
		mainCanvas.WriteBytesASCII(panel2Start+1, 12, "Errors: 0.03%")
		mainCanvas.WriteBytesASCII(panel2Start+1, 13, "Requests: 847K")
		mainCanvas.WriteBytesASCII(panel2Start+1, 14, "Cache Hit: 94%")
	}
	
	if panel3Start+panelWidth <= maxX {
		mainCanvas.WriteBytesASCII(panel3Start+1, 8, "SECURITY")
		mainCanvas.WriteBytesASCII(panel3Start+1, 10, "Threats: 0")
		mainCanvas.WriteBytesASCII(panel3Start+1, 11, "Blocked: 127")
		mainCanvas.WriteBytesASCII(panel3Start+1, 12, "Firewall: ON")
		mainCanvas.WriteBytesASCII(panel3Start+1, 13, "SSL: Valid")
		mainCanvas.WriteBytesASCII(panel3Start+1, 14, "Backup: OK")
	}
	
	// Chart area - responsive positioning
	if chartEnd > panel1Start+panelWidth*2 {
		mainCanvas.WriteBytesASCII(panel1Start+1, 17, "ANALYTICS CHART - Last 7 Days")
		mainCanvas.WriteBytesASCII(panel1Start+1, 19, "Revenue |#######*****:::::....")
		mainCanvas.WriteBytesASCII(panel1Start+1, 20, "Traffic |****#######****::....")
		mainCanvas.WriteBytesASCII(panel1Start+1, 21, "Users   |:::*****########**..")
		mainCanvas.WriteBytesASCII(panel1Start+1, 22, "Errors  |.....::::*****......")
		mainCanvas.WriteBytesASCII(panel1Start+1, 23, "        +------------------------")
		mainCanvas.WriteBytesASCII(panel1Start+1, 24, "        Mon Tue Wed Thu Fri Sat Sun")
		
		// Live feed
		mainCanvas.WriteBytesASCII(chartEnd+2, 17, "LIVE ACTIVITY")
		mainCanvas.WriteBytesASCII(chartEnd+2, 18, "14:32 Login: admin")
		mainCanvas.WriteBytesASCII(chartEnd+2, 19, "14:31 Order #4891")
		mainCanvas.WriteBytesASCII(chartEnd+2, 20, "14:30 User signup")
		mainCanvas.WriteBytesASCII(chartEnd+2, 21, "14:29 Payment OK")
		mainCanvas.WriteBytesASCII(chartEnd+2, 22, "14:28 Backup done")
		mainCanvas.WriteBytesASCII(chartEnd+2, 23, "14:27 Alert clear")
	}
	
	// Activity log
	mainCanvas.WriteBytesASCII(sidebarWidth+2, 27, "SYSTEM LOG - Recent Activities")
	if currentWidth >= 100 {
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 28, "[INFO] 14:32:17 Database connection pool expanded to 50 connections")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 29, "[WARN] 14:31:45 High memory usage detected on server node-03")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 30, "[INFO] 14:30:22 Scheduled backup completed successfully")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 31, "[INFO] 14:29:33 User authentication rate: 1,247 logins/hour")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 32, "[DEBUG] 14:28:11 Cache refresh cycle completed in 2.3 seconds")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 33, "[INFO] 14:27:56 SSL certificate validation passed")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 34, "[WARN] 14:26:42 API rate limit reached for client 192.168.1.100")
	} else {
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 28, "[INFO] 14:32:17 Database pool expanded")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 29, "[WARN] 14:31:45 High memory usage")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 30, "[INFO] 14:30:22 Backup completed")
		mainCanvas.WriteBytesASCII(sidebarWidth+2, 31, "[INFO] 14:29:33 Auth rate: 1,247/hour")
	}
	
	// Status bar
	mainCanvas.WriteBytesASCII(2, 37, "Connected: DB-MAIN | Cache: REDIS-01 | Queue: 247 | Alerts: 0")
	if currentWidth >= 80 {
		mainCanvas.WriteBytesASCII(currentWidth-12, 37, "F10:Settings")
	}
	
	// Generate output
	configName := canvas.GetConfigName()
	output := fmt.Sprintf("Enterprise Dashboard UI\n")
	output += fmt.Sprintf("========================\n\n")
	output += fmt.Sprintf("Configuration: %s\n", configName)
	output += fmt.Sprintf("Features: Multi-panel layout, real-time data, charts, logs\n")
	output += fmt.Sprintf("ASCII Filter: Enabled (all full-width characters removed)\n\n")
	output += fmt.Sprintf("Layout:\n")
	output += mainCanvas.String()
	
	// If not in insert mode, write to file
	if !insertMode {
		// Determine output filename
		filename := outputFile
		if filename == "" {
			filename = fmt.Sprintf("enterprise_ui_%dx%d.txt", currentWidth, canvas.GetCurrentHeight())
		}
		
		err := os.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			return output
		}
		
		fmt.Printf("Enterprise Dashboard UI created: %s\n", filename)
	}
	
	return output
}

func generateMobileUI(outputFile string, insertMode bool) string {
	// Create mobile smartphone UI
	mobileCanvas := canvas.NewByteCanvas()
	
	currentWidth := canvas.GetCurrentWidth()
	maxX := currentWidth - 1
	
	// Status bar
	mobileCanvas.DrawBox(0, 0, maxX, 2)
	mobileCanvas.WriteBytesASCII(1, 1, "12:34 PM")
	if currentWidth >= 60 {
		mobileCanvas.WriteBytesASCII(currentWidth/2-2, 1, "5G")
		mobileCanvas.WriteBytesASCII(currentWidth-12, 1, "[====] 87%")
	}
	
	// Header
	mobileCanvas.DrawBox(0, 3, maxX, 5)
	mobileCanvas.WriteBytesASCII(1, 4, "<-")
	mobileCanvas.WriteBytesASCII(currentWidth/2-4, 4, "MESSAGES")
	if currentWidth >= 50 {
		mobileCanvas.WriteBytesASCII(maxX-3, 4, "[+]")
	}
	
	// Menu/Stats
	mobileCanvas.DrawBox(0, 6, maxX, 10)
	mobileCanvas.WriteBytesASCII(2, 7, "[1] New Message")
	mobileCanvas.WriteBytesASCII(2, 8, "[2] Contacts")
	mobileCanvas.WriteBytesASCII(2, 9, "[3] Recent Chats")
	
	if currentWidth >= 50 {
		mobileCanvas.WriteBytesASCII(currentWidth-15, 7, "Active: 3")
		mobileCanvas.WriteBytesASCII(currentWidth-15, 8, "Total: 127")
		mobileCanvas.WriteBytesASCII(currentWidth-15, 9, "Unread: 5")
	}
	
	// Message list
	mobileCanvas.DrawBox(0, 11, maxX, 20)
	mobileCanvas.WriteBytesASCII(1, 12, "John Doe")
	if currentWidth >= 50 {
		mobileCanvas.WriteBytesASCII(currentWidth-10, 12, "2:30 PM")
	}
	mobileCanvas.WriteBytesASCII(1, 13, "Hey, are you free for lunch?")
	
	mobileCanvas.WriteBytesASCII(1, 15, "Jane Smith")
	if currentWidth >= 50 {
		mobileCanvas.WriteBytesASCII(currentWidth-10, 15, "1:45 PM")
	}
	mobileCanvas.WriteBytesASCII(1, 16, "Meeting at 3 PM confirmed")
	
	if currentWidth >= 70 {
		mobileCanvas.WriteBytesASCII(1, 18, "Mike Johnson")
		mobileCanvas.WriteBytesASCII(currentWidth-10, 18, "12:15 PM")
		mobileCanvas.WriteBytesASCII(1, 19, "Thanks for the help!")
	}
	
	// Bottom navigation
	mobileCanvas.DrawBox(0, 21, maxX, 23)
	if currentWidth >= 60 {
		navWidth := currentWidth / 5
		mobileCanvas.WriteBytesASCII(navWidth*0+2, 22, "[HOME]")
		mobileCanvas.WriteBytesASCII(navWidth*1+2, 22, "[CHAT]")
		mobileCanvas.WriteBytesASCII(navWidth*2+2, 22, "[CALL]")
		mobileCanvas.WriteBytesASCII(navWidth*3+2, 22, "[MORE]")
		mobileCanvas.WriteBytesASCII(navWidth*4+2, 22, "[USER]")
	} else {
		mobileCanvas.WriteBytesASCII(2, 22, "[HOME] [CHAT] [CALL] [USER]")
	}
	
	// Generate output
	configName := canvas.GetConfigName()
	output := fmt.Sprintf("Mobile Smartphone UI\n")
	output += fmt.Sprintf("====================\n\n")
	output += fmt.Sprintf("Configuration: %s\n", configName)
	output += fmt.Sprintf("Features: Mobile interface, messaging, navigation\n")
	output += fmt.Sprintf("ASCII Filter: Enabled (all full-width characters removed)\n\n")
	output += fmt.Sprintf("Layout:\n")
	output += mobileCanvas.String()
	
	// If not in insert mode, write to file
	if !insertMode {
		// Determine output filename
		filename := outputFile
		if filename == "" {
			filename = fmt.Sprintf("mobile_ui_%dx%d.txt", currentWidth, canvas.GetCurrentHeight())
		}
		
		err := os.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			return output
		}
		
		fmt.Printf("Mobile UI created: %s\n", filename)
	}
	
	return output
}

func generateSimpleUI(outputFile string, insertMode bool) string {
	// Create simple box layout UI
	simpleCanvas := canvas.NewByteCanvas()
	
	currentWidth := canvas.GetCurrentWidth()
	maxX := currentWidth - 1
	
	// Main container
	simpleCanvas.DrawBox(0, 0, maxX, 15)
	
	// Header
	simpleCanvas.DrawBox(2, 2, maxX-2, 4)
	simpleCanvas.WriteBytesASCII(4, 3, "Simple UI Layout")
	
	// Content areas
	halfWidth := currentWidth / 2
	simpleCanvas.DrawBox(2, 6, halfWidth-2, 12)
	simpleCanvas.DrawBox(halfWidth, 6, maxX-2, 12)
	
	// Content
	simpleCanvas.WriteBytesASCII(4, 7, "Left Panel")
	simpleCanvas.WriteBytesASCII(4, 9, "Content Area 1")
	simpleCanvas.WriteBytesASCII(4, 10, "- Item 1")
	simpleCanvas.WriteBytesASCII(4, 11, "- Item 2")
	
	simpleCanvas.WriteBytesASCII(halfWidth+2, 7, "Right Panel")
	simpleCanvas.WriteBytesASCII(halfWidth+2, 9, "Content Area 2")
	simpleCanvas.WriteBytesASCII(halfWidth+2, 10, "- Option A")
	simpleCanvas.WriteBytesASCII(halfWidth+2, 11, "- Option B")
	
	// Footer
	simpleCanvas.WriteBytesASCII(4, 14, "Status: Ready")
	if currentWidth >= 60 {
		simpleCanvas.WriteBytesASCII(currentWidth-15, 14, "Press ESC to exit")
	}
	
	// Generate output
	configName := canvas.GetConfigName()
	output := fmt.Sprintf("Simple UI Layout\n")
	output += fmt.Sprintf("================\n\n")
	output += fmt.Sprintf("Configuration: %s\n", configName)
	output += fmt.Sprintf("Features: Basic box layout, two-panel design\n")
	output += fmt.Sprintf("ASCII Filter: Enabled (all full-width characters removed)\n\n")
	output += fmt.Sprintf("Layout:\n")
	output += simpleCanvas.String()
	
	// If not in insert mode, write to file
	if !insertMode {
		// Determine output filename
		filename := outputFile
		if filename == "" {
			filename = fmt.Sprintf("simple_ui_%dx%d.txt", currentWidth, canvas.GetCurrentHeight())
		}
		
		err := os.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			return output
		}
		
		fmt.Printf("Simple UI created: %s\n", filename)
	}
	
	return output
}

// processYAML handles YAML input processing
func processYAML(yamlFile, outputFile, insertSpec string, backup bool, japaneseMode bool) error {
	var reader *os.File
	var err error
	
	// Open YAML file or stdin
	if yamlFile == "-" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(yamlFile)
		if err != nil {
			return fmt.Errorf("failed to open YAML file: %w", err)
		}
		defer reader.Close()
	}
	
	// Parse YAML
	parser := yaml.NewParser()
	spec, err := parser.Parse(reader)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}
	
	// Override Japanese mode if --jp flag is set
	if japaneseMode {
		spec.Canvas.JapaneseMode = true
	}
	
	// Render to ASCII art
	result, err := parser.Render(spec)
	if err != nil {
		return fmt.Errorf("failed to render: %w", err)
	}
	
	// Handle output
	if insertSpec != "" {
		// Insert mode
		parts := strings.Split(insertSpec, ":")
		if len(parts) != 2 {
			return fmt.Errorf("insert format must be 'file:line'")
		}
		
		insertFile := parts[0]
		insertLine, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("invalid line number: %w", err)
		}
		
		return insertIntoFile(insertFile, insertLine, result, backup)
	} else if outputFile != "" {
		// Write to specified file
		return os.WriteFile(outputFile, []byte(result), 0644)
	} else {
		// Write to stdout
		fmt.Print(result)
	}
	
	return nil
}

