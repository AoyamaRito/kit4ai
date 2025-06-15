package main

import (
	"fmt"
	"os"
	"kit4ai/pkg/canvas"
)

func main() {
	// Set standard 80-character width configuration
	canvas.SetConfig(canvas.StandardConfig)
	
	// Create complex enterprise dashboard UI
	mainCanvas := canvas.NewByteCanvas()
	
	// Title bar
	mainCanvas.DrawBox(0, 0, 79, 2)
	
	// Menu bar
	mainCanvas.DrawBox(0, 3, 79, 5)
	
	// Left sidebar (navigation)
	mainCanvas.DrawBox(0, 6, 19, 35)
	
	// Main content area
	mainCanvas.DrawBox(20, 6, 79, 25)
	
	// Statistics panels (3 columns)
	mainCanvas.DrawBox(21, 7, 38, 15)
	mainCanvas.DrawBox(40, 7, 58, 15)
	mainCanvas.DrawBox(60, 7, 78, 15)
	
	// Chart area
	mainCanvas.DrawBox(21, 16, 58, 24)
	
	// Live feed
	mainCanvas.DrawBox(60, 16, 78, 24)
	
	// Log/activity area
	mainCanvas.DrawBox(20, 26, 79, 35)
	
	// Status bar
	mainCanvas.DrawBox(0, 36, 79, 38)
	
	// Title bar content
	mainCanvas.WriteBytesASCII(2, 1, "ENTERPRISE CONTROL PANEL v2.4.1")
	mainCanvas.WriteBytesASCII(60, 1, "2024-06-15 14:32:17")
	
	// Menu items
	mainCanvas.WriteBytesASCII(2, 4, "[F1]File [F2]Edit [F3]View [F4]Tools [F5]Reports [F6]Admin [ESC]Exit")
	
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
	
	// Statistics panels
	mainCanvas.WriteBytesASCII(22, 8, "REVENUE METRICS")
	mainCanvas.WriteBytesASCII(22, 10, "Daily: $47,892")
	mainCanvas.WriteBytesASCII(22, 11, "Weekly: $312,456")
	mainCanvas.WriteBytesASCII(22, 12, "Monthly: $1.2M")
	mainCanvas.WriteBytesASCII(22, 13, "Growth: +12.5%")
	mainCanvas.WriteBytesASCII(22, 14, "Target: 87%")
	
	mainCanvas.WriteBytesASCII(41, 8, "PERFORMANCE")
	mainCanvas.WriteBytesASCII(41, 10, "Avg Resp: 245ms")
	mainCanvas.WriteBytesASCII(41, 11, "Uptime: 99.97%")
	mainCanvas.WriteBytesASCII(41, 12, "Errors: 0.03%")
	mainCanvas.WriteBytesASCII(41, 13, "Requests: 847K")
	mainCanvas.WriteBytesASCII(41, 14, "Cache Hit: 94%")
	
	mainCanvas.WriteBytesASCII(61, 8, "SECURITY")
	mainCanvas.WriteBytesASCII(61, 10, "Threats: 0")
	mainCanvas.WriteBytesASCII(61, 11, "Blocked: 127")
	mainCanvas.WriteBytesASCII(61, 12, "Firewall: ON")
	mainCanvas.WriteBytesASCII(61, 13, "SSL: Valid")
	mainCanvas.WriteBytesASCII(61, 14, "Backup: OK")
	
	// Chart area
	mainCanvas.WriteBytesASCII(22, 17, "ANALYTICS CHART - Last 7 Days")
	mainCanvas.WriteBytesASCII(22, 19, "Revenue |#######*****:::::....")
	mainCanvas.WriteBytesASCII(22, 20, "Traffic |****#######****::....")
	mainCanvas.WriteBytesASCII(22, 21, "Users   |:::*****########**..")
	mainCanvas.WriteBytesASCII(22, 22, "Errors  |.....::::*****......")
	mainCanvas.WriteBytesASCII(22, 23, "        +------------------------")
	mainCanvas.WriteBytesASCII(22, 24, "        Mon Tue Wed Thu Fri Sat Sun")
	
	// Live feed
	mainCanvas.WriteBytesASCII(61, 17, "LIVE ACTIVITY")
	mainCanvas.WriteBytesASCII(61, 18, "14:32 Login: admin")
	mainCanvas.WriteBytesASCII(61, 19, "14:31 Order #4891")
	mainCanvas.WriteBytesASCII(61, 20, "14:30 User signup")
	mainCanvas.WriteBytesASCII(61, 21, "14:29 Payment OK")
	mainCanvas.WriteBytesASCII(61, 22, "14:28 Backup done")
	mainCanvas.WriteBytesASCII(61, 23, "14:27 Alert clear")
	
	// Activity log
	mainCanvas.WriteBytesASCII(21, 27, "SYSTEM LOG - Recent Activities")
	mainCanvas.WriteBytesASCII(21, 28, "[INFO] 14:32:17 Database connection pool expanded to 50 connections")
	mainCanvas.WriteBytesASCII(21, 29, "[WARN] 14:31:45 High memory usage detected on server node-03")
	mainCanvas.WriteBytesASCII(21, 30, "[INFO] 14:30:22 Scheduled backup completed successfully")
	mainCanvas.WriteBytesASCII(21, 31, "[INFO] 14:29:33 User authentication rate: 1,247 logins/hour")
	mainCanvas.WriteBytesASCII(21, 32, "[DEBUG] 14:28:11 Cache refresh cycle completed in 2.3 seconds")
	mainCanvas.WriteBytesASCII(21, 33, "[INFO] 14:27:56 SSL certificate validation passed")
	mainCanvas.WriteBytesASCII(21, 34, "[WARN] 14:26:42 API rate limit reached for client 192.168.1.100")
	
	// Status bar
	mainCanvas.WriteBytesASCII(2, 37, "Connected: DB-MAIN | Cache: REDIS-01 | Queue: 247 | Alerts: 0")
	mainCanvas.WriteBytesASCII(65, 37, "F10:Settings")
	
	// Output
	output := fmt.Sprintf("Complex Enterprise Dashboard UI\n")
	output += fmt.Sprintf("===============================\n\n")
	output += fmt.Sprintf("Configuration: Standard (80x100) - Enterprise Display\n")
	output += fmt.Sprintf("Features: Multi-panel layout, real-time data, charts, logs\n")
	output += fmt.Sprintf("ASCII Filter: Enabled (all full-width characters removed)\n\n")
	output += fmt.Sprintf("Layout:\n")
	output += mainCanvas.String()
	
	err := os.WriteFile("complex_enterprise_ui.txt", []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}
	
	fmt.Println("Complex Enterprise Dashboard UI created: complex_enterprise_ui.txt")
}