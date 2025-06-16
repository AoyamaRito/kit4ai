package main

import (
	"fmt"
	"os"
	"kit4ai/pkg/canvas"
)

func main() {
	// Set ultra-wide configuration for complex layout
	canvas.SetConfig(canvas.UltraWideConfig)
	
	// Create canvas and arrow layer
	ui := canvas.NewByteCanvas()
	arrowLayer := canvas.NewArrowLayer()
	
	// Main title
	ui.DrawBox(5, 1, 115, 5)
	ui.WriteBytesASCII(25, 2, "🏢 ENTERPRISE MICROSERVICES ARCHITECTURE")
	ui.WriteBytesASCII(35, 3, "Real-time Financial Trading Platform")
	ui.WriteBytesASCII(50, 4, "2024 Q4 Design")
	
	// Client Layer
	ui.WriteBytesASCII(5, 8, "CLIENT LAYER:")
	drawServiceBox(ui, 5, 10, "WEB APP", []string{"React", "TypeScript", "WebSocket"})
	drawServiceBox(ui, 25, 10, "MOBILE APP", []string{"React Native", "Push Notif", "Biometric"})
	drawServiceBox(ui, 45, 10, "DESKTOP", []string{"Electron", "Native API", "Hardware"})
	drawServiceBox(ui, 65, 10, "TRADING BOT", []string{"Python", "ML Models", "APIs"})
	
	// Load Balancer
	drawServiceBox(ui, 35, 18, "LOAD BALANCER", []string{"NGINX", "SSL Term", "Rate Limit"})
	
	// Client to Load Balancer arrows
	arrowLayer.AddLabeledArrow(12, 16, 40, 18, canvas.ArrowStyleThick, "HTTPS")
	arrowLayer.AddLabeledArrow(32, 16, 40, 18, canvas.ArrowStyleThick, "HTTPS")
	arrowLayer.AddLabeledArrow(52, 16, 45, 18, canvas.ArrowStyleThick, "WSS")
	arrowLayer.AddLabeledArrow(72, 16, 45, 18, canvas.ArrowStyleNormal, "API")
	
	// API Gateway Layer
	ui.WriteBytesASCII(5, 26, "API GATEWAY LAYER:")
	drawServiceBox(ui, 15, 28, "AUTH GATEWAY", []string{"JWT", "OAuth2", "MFA"})
	drawServiceBox(ui, 45, 28, "API GATEWAY", []string{"GraphQL", "REST", "gRPC"})
	drawServiceBox(ui, 75, 28, "RATE LIMITER", []string{"Redis", "Circuit Break", "Throttle"})
	
	// Load Balancer to API Gateway
	arrowLayer.AddLabeledArrow(40, 24, 22, 28, canvas.ArrowStyleThick, "")
	arrowLayer.AddLabeledArrow(40, 24, 52, 28, canvas.ArrowStyleThick, "")
	arrowLayer.AddLabeledArrow(45, 24, 82, 28, canvas.ArrowStyleNormal, "")
	
	// Core Services Layer
	ui.WriteBytesASCII(5, 36, "CORE SERVICES:")
	drawServiceBox(ui, 5, 38, "USER SERVICE", []string{"Profile", "KYC", "Settings"})
	drawServiceBox(ui, 25, 38, "AUTH SERVICE", []string{"Login", "2FA", "Session"})
	drawServiceBox(ui, 45, 38, "ACCOUNT SVC", []string{"Balance", "Wallet", "History"})
	drawServiceBox(ui, 65, 38, "TRADING SVC", []string{"Orders", "Execution", "Risk"})
	drawServiceBox(ui, 85, 38, "MARKET DATA", []string{"Real-time", "History", "Analytics"})
	drawServiceBox(ui, 105, 38, "NOTIFICATION", []string{"Email", "SMS", "Push"})
	
	// API Gateway to Services
	arrowLayer.AddLabeledArrow(22, 34, 12, 38, canvas.ArrowStyleNormal, "")
	arrowLayer.AddLabeledArrow(25, 34, 32, 38, canvas.ArrowStyleNormal, "")
	arrowLayer.AddLabeledArrow(52, 34, 52, 38, canvas.ArrowStyleThick, "")
	arrowLayer.AddLabeledArrow(55, 34, 72, 38, canvas.ArrowStyleThick, "")
	arrowLayer.AddLabeledArrow(58, 34, 92, 38, canvas.ArrowStyleWave, "")
	arrowLayer.AddLabeledArrow(60, 34, 112, 38, canvas.ArrowStyleDotted, "")
	
	// Data Layer
	ui.WriteBytesASCII(5, 46, "DATA & INFRASTRUCTURE:")
	drawServiceBox(ui, 5, 48, "USER DB", []string{"PostgreSQL", "Encrypted", "Backup"})
	drawServiceBox(ui, 25, 48, "TRADING DB", []string{"MongoDB", "Sharded", "Real-time"})
	drawServiceBox(ui, 45, 48, "CACHE", []string{"Redis", "Cluster", "Failover"})
	drawServiceBox(ui, 65, 48, "MESSAGE Q", []string{"RabbitMQ", "Kafka", "Events"})
	drawServiceBox(ui, 85, 48, "BLOCKCHAIN", []string{"Ethereum", "Bitcoin", "DeFi"})
	drawServiceBox(ui, 105, 48, "FILE STORE", []string{"S3", "CDN", "Images"})
	
	// Services to Data arrows
	arrowLayer.AddLabeledArrow(12, 44, 12, 48, canvas.ArrowStyleNormal, "")
	arrowLayer.AddLabeledArrow(32, 44, 32, 48, canvas.ArrowStyleNormal, "")
	arrowLayer.AddLabeledArrow(72, 44, 52, 48, canvas.ArrowStyleThick, "")
	arrowLayer.AddLabeledArrow(72, 44, 72, 48, canvas.ArrowStyleDouble, "")
	arrowLayer.AddLabeledArrow(92, 44, 92, 48, canvas.ArrowStyleWave, "")
	arrowLayer.AddLabeledArrow(112, 44, 112, 48, canvas.ArrowStyleDotted, "")
	
	// External Services
	ui.WriteBytesASCII(5, 56, "EXTERNAL INTEGRATIONS:")
	drawServiceBox(ui, 5, 58, "BANK API", []string{"ACH", "Wire", "Swift"})
	drawServiceBox(ui, 25, 58, "KYC/AML", []string{"Jumio", "Identity", "Compliance"})
	drawServiceBox(ui, 45, 58, "EXCHANGE", []string{"Binance", "Coinbase", "FTX"})
	drawServiceBox(ui, 65, 58, "PRICE FEED", []string{"CoinGecko", "Real-time", "WebSocket"})
	drawServiceBox(ui, 85, 58, "MONITORING", []string{"DataDog", "Grafana", "Alerts"})
	
	// Cross-service connections
	arrowLayer.AddLabeledArrow(52, 54, 12, 58, canvas.ArrowStyleDotted, "ACH")
	arrowLayer.AddLabeledArrow(12, 54, 32, 58, canvas.ArrowStyleDotted, "KYC")
	arrowLayer.AddLabeledArrow(72, 54, 52, 58, canvas.ArrowStyleThick, "Trade")
	arrowLayer.AddLabeledArrow(92, 54, 72, 58, canvas.ArrowStyleWave, "Price")
	
	// Security Layer (Overlay)
	ui.DrawBox(95, 8, 115, 25)
	ui.WriteBytesASCII(97, 9, "🔒 SECURITY")
	ui.WriteBytesASCII(97, 11, "• WAF")
	ui.WriteBytesASCII(97, 12, "• DDoS Protect")
	ui.WriteBytesASCII(97, 13, "• Encryption")
	ui.WriteBytesASCII(97, 14, "• Audit Logs")
	ui.WriteBytesASCII(97, 15, "• Compliance")
	ui.WriteBytesASCII(97, 16, "• Penetration")
	ui.WriteBytesASCII(97, 17, "• SOC 2")
	ui.WriteBytesASCII(97, 18, "• PCI DSS")
	ui.WriteBytesASCII(97, 19, "• ISO 27001")
	ui.WriteBytesASCII(97, 21, "🚨 24/7 SOC")
	ui.WriteBytesASCII(97, 22, "🔍 SIEM")
	ui.WriteBytesASCII(97, 23, "🛡️ Zero Trust")
	
	// Performance Metrics
	ui.DrawBox(5, 66, 115, 78)
	ui.WriteBytesASCII(45, 67, "📊 SYSTEM PERFORMANCE METRICS")
	
	ui.WriteBytesASCII(8, 69, "💰 TRADING VOLUME:")
	ui.WriteBytesASCII(8, 70, "├─ Daily: $2.4B")
	ui.WriteBytesASCII(8, 71, "├─ Peak TPS: 50,000")
	ui.WriteBytesASCII(8, 72, "├─ Avg Latency: <2ms")
	ui.WriteBytesASCII(8, 73, "└─ Orders/sec: 10,000")
	
	ui.WriteBytesASCII(40, 69, "🏗️ INFRASTRUCTURE:")
	ui.WriteBytesASCII(40, 70, "├─ Kubernetes Cluster")
	ui.WriteBytesASCII(40, 71, "├─ 500+ Microservices")
	ui.WriteBytesASCII(40, 72, "├─ Multi-region AWS")
	ui.WriteBytesASCII(40, 73, "└─ 99.99% Uptime SLA")
	
	ui.WriteBytesASCII(75, 69, "🔥 REAL-TIME FEATURES:")
	ui.WriteBytesASCII(75, 70, "├─ Live Price Streaming")
	ui.WriteBytesASCII(75, 71, "├─ Order Book Updates")
	ui.WriteBytesASCII(75, 72, "├─ Trade Notifications")
	ui.WriteBytesASCII(75, 73, "└─ Risk Monitoring")
	
	ui.WriteBytesASCII(8, 75, "🌍 GLOBAL SCALE: 50M+ Users | 200+ Countries | 24/7/365 Operations")
	ui.WriteBytesASCII(8, 76, "🔐 COMPLIANCE: SEC, CFTC, MiFID II, GDPR, SOX, Basel III")
	
	// Render all arrows
	err := arrowLayer.Render(ui)
	if err != nil {
		fmt.Printf("Error rendering arrows: %v\n", err)
		return
	}
	
	// Footer
	ui.WriteBytesASCII(5, 80, "Generated by Kit4AI | Enterprise Architecture Visualization | Ultra-Wide Layout")
	
	// Output
	fmt.Println("🏢 ENTERPRISE MICROSERVICES ARCHITECTURE")
	fmt.Println("========================================")
	fmt.Println("Complex real-time financial trading platform")
	fmt.Println()
	fmt.Println(ui.String())
	
	// Save to file
	content := fmt.Sprintf("ENTERPRISE MICROSERVICES ARCHITECTURE\n=====================================\n\nReal-time Financial Trading Platform\nGenerated by Kit4AI Ultra-Wide System\n\n%s", ui.String())
	
	file, err := os.Create("complex_enterprise_architecture.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	} else {
		defer file.Close()
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		} else {
			fmt.Printf("\n🏢 Complex enterprise architecture saved: complex_enterprise_architecture.txt\n")
		}
	}
}

func drawServiceBox(ui *canvas.ByteCanvas, x, y int, title string, details []string) {
	// Draw service box with title and details
	ui.DrawBox(x, y, x+18, y+6)
	
	// Title (centered)
	titleX := x + (18-len(title))/2
	ui.WriteBytesASCII(titleX, y+1, title)
	
	// Details
	for i, detail := range details {
		if i < 3 {
			detailX := x + (18-len(detail))/2
			ui.WriteBytesASCII(detailX, y+2+i, detail)
		}
	}
}