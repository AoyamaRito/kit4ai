package main

import (
	"fmt"
	"os"
	"kit4ai/pkg/canvas"
)

func main() {
	// Set ultra-wide configuration for complex UI
	canvas.SetConfig(canvas.UltraWideConfig)
	
	// Create canvas
	ui := canvas.NewByteCanvas()
	
	// Header Bar with System Status
	ui.DrawBox(0, 0, 119, 4)
	ui.WriteBytesASCII(2, 1, "BLOOMBERG TERMINAL 5.0 | TRADING DESK WORKSTATION")
	ui.WriteBytesASCII(60, 1, "USER: J.SMITH | DESK: EQUITY-NY-01")
	ui.WriteBytesASCII(2, 2, "15:42:33 EST | NYSE OPEN | LAST: SPX 4,567.89 +12.34 (+0.27%)")
	ui.WriteBytesASCII(60, 2, "P&L: +$47,892.34 | POS: 15 | ALERTS: 3")
	ui.WriteBytesASCII(95, 2, "CPU:67% RAM:8.2GB")
	
	// Menu Bar
	ui.DrawBox(0, 4, 119, 6)
	ui.WriteBytesASCII(2, 5, "[F1]Markets [F2]News [F3]Charts [F4]Options [F5]Bonds [F6]FX [F7]Crypto [F8]Research [F9]Risk [F10]Settings [ESC]Exit")
	
	// Main Content Area - Multi-panel layout
	
	// Left Panel - Market Data & Watchlist
	ui.DrawBox(0, 6, 39, 35)
	ui.WriteBytesASCII(12, 7, "MARKET OVERVIEW & WATCHLIST")
	
	// Market Indices
	ui.WriteBytesASCII(2, 9, "INDICES          LAST     CHG    %CHG")
	ui.WriteBytesASCII(2, 10, "------------------------------------")
	ui.WriteBytesASCII(2, 11, "S&P 500       4,567.89  +12.34  +0.27")
	ui.WriteBytesASCII(2, 12, "NASDAQ        14,234.56  -23.45  -0.16")
	ui.WriteBytesASCII(2, 13, "DOW JONES     35,123.78  +89.12  +0.25")
	ui.WriteBytesASCII(2, 14, "VIX              18.45   -0.23  -1.23")
	
	// Personal Watchlist
	ui.WriteBytesASCII(2, 16, "MY WATCHLIST     LAST     CHG    VOL")
	ui.WriteBytesASCII(2, 17, "------------------------------------")
	ui.WriteBytesASCII(2, 18, "AAPL          175.23   +2.45   45.2M")
	ui.WriteBytesASCII(2, 19, "MSFT          342.67   -1.23   32.1M")
	ui.WriteBytesASCII(2, 20, "GOOGL       2,789.45   +8.90   12.5M")
	ui.WriteBytesASCII(2, 21, "TSLA          234.56  +12.34   89.7M")
	ui.WriteBytesASCII(2, 22, "NVDA          456.78   -5.67   67.3M")
	
	// Economic Calendar
	ui.WriteBytesASCII(2, 24, "TODAY'S EVENTS   TIME    IMP")
	ui.WriteBytesASCII(2, 25, "------------------------------------")
	ui.WriteBytesASCII(2, 26, "Fed Speech       16:00   HIGH")
	ui.WriteBytesASCII(2, 27, "GDP Q3 Final     16:30   MED")
	ui.WriteBytesASCII(2, 28, "Oil Inventory    17:00   LOW")
	
	// Market Sentiment
	ui.WriteBytesASCII(2, 30, "SENTIMENT    BULL/BEAR   FEAR/GREED")
	ui.WriteBytesASCII(2, 31, "------------------------------------")
	ui.WriteBytesASCII(2, 32, "EQUITY          67/33        45/100")
	ui.WriteBytesASCII(2, 33, "CRYPTO          23/77        78/100")
	
	// Center Panel - Price Chart & Order Book
	ui.DrawBox(39, 6, 79, 35)
	ui.WriteBytesASCII(52, 7, "AAPL - APPLE INC. $175.23")
	
	// Price Chart (ASCII representation)
	ui.WriteBytesASCII(41, 9, "1D Chart    1W    1M    3M    1Y    5Y")
	ui.WriteBytesASCII(41, 10, "=====================================")
	ui.WriteBytesASCII(41, 11, "180|                        ***")
	ui.WriteBytesASCII(41, 12, "175|              *****  ***")
	ui.WriteBytesASCII(41, 13, "170|        ****           *")
	ui.WriteBytesASCII(41, 14, "165|   ***")
	ui.WriteBytesASCII(41, 15, "160|***")
	ui.WriteBytesASCII(41, 16, "   +---+---+---+---+---+---+---+")
	ui.WriteBytesASCII(41, 17, "   9:30    11    13    15   16:00")
	
	// Level II Order Book
	ui.WriteBytesASCII(41, 19, "LEVEL II ORDER BOOK")
	ui.WriteBytesASCII(41, 20, "BID      SIZE    ASK      SIZE")
	ui.WriteBytesASCII(41, 21, "--------------------------------")
	ui.WriteBytesASCII(41, 22, "175.22   1,500   175.23   2,100")
	ui.WriteBytesASCII(41, 23, "175.21   3,200   175.24   1,800")
	ui.WriteBytesASCII(41, 24, "175.20   2,800   175.25   4,500")
	ui.WriteBytesASCII(41, 25, "175.19   1,900   175.26   3,300")
	ui.WriteBytesASCII(41, 26, "175.18   5,100   175.27   2,700")
	
	// Trade History
	ui.WriteBytesASCII(41, 28, "RECENT TRADES    TIME    SIZE")
	ui.WriteBytesASCII(41, 29, "--------------------------------")
	ui.WriteBytesASCII(41, 30, "175.23          15:42:31  500")
	ui.WriteBytesASCII(41, 31, "175.22          15:42:28  1200")
	ui.WriteBytesASCII(41, 32, "175.24          15:42:25  800")
	ui.WriteBytesASCII(41, 33, "175.23          15:42:22  300")
	
	// Right Panel - Trading & Portfolio
	ui.DrawBox(79, 6, 119, 35)
	ui.WriteBytesASCII(89, 7, "TRADING & PORTFOLIO")
	
	// Order Entry
	ui.WriteBytesASCII(81, 9, "ORDER ENTRY")
	ui.WriteBytesASCII(81, 10, "Symbol: [AAPL    ] Qty: [1000  ]")
	ui.WriteBytesASCII(81, 11, "Type: [LIMIT] Price: [175.20  ]")
	ui.WriteBytesASCII(81, 12, "TIF: [DAY ] [BUY ] [SELL] [CANCEL]")
	
	// Open Orders
	ui.WriteBytesASCII(81, 14, "OPEN ORDERS")
	ui.WriteBytesASCII(81, 15, "SYM   SIDE  QTY   PRICE   STATUS")
	ui.WriteBytesASCII(81, 16, "------------------------------------")
	ui.WriteBytesASCII(81, 17, "MSFT  BUY   500  340.00  WORKING")
	ui.WriteBytesASCII(81, 18, "GOOGL SELL  200 2800.00  WORKING")
	ui.WriteBytesASCII(81, 19, "TSLA  BUY  1000  230.00  PARTIAL")
	
	// Portfolio Positions
	ui.WriteBytesASCII(81, 21, "POSITIONS")
	ui.WriteBytesASCII(81, 22, "SYM   QTY   AVG_PX   MKT_VAL  P&L")
	ui.WriteBytesASCII(81, 23, "------------------------------------")
	ui.WriteBytesASCII(81, 24, "AAPL  2500  165.45   438,075 +24,575")
	ui.WriteBytesASCII(81, 25, "MSFT  1800  320.12   617,006 +41,694")
	ui.WriteBytesASCII(81, 26, "GOOGL  300  2650.00  836,835 +41,535")
	ui.WriteBytesASCII(81, 27, "CASH              1,234,567")
	
	// Risk Metrics
	ui.WriteBytesASCII(81, 29, "RISK METRICS")
	ui.WriteBytesASCII(81, 30, "Buying Power:     $2,456,789")
	ui.WriteBytesASCII(81, 31, "Day P&L:          +$47,892")
	ui.WriteBytesASCII(81, 32, "Total P&L:        +$234,567")
	ui.WriteBytesASCII(81, 33, "VaR (1D, 95%):    -$89,234")
	
	// Bottom Panel - News & Messages
	ui.DrawBox(0, 35, 119, 50)
	ui.WriteBytesASCII(45, 36, "REAL-TIME NEWS & MARKET UPDATES")
	
	// News Feed
	ui.WriteBytesASCII(2, 38, "BREAKING NEWS")
	ui.WriteBytesASCII(2, 39, "15:42 | Fed Chair Powell signals potential rate pause in December meeting")
	ui.WriteBytesASCII(2, 40, "15:41 | Apple reports strong iPhone 15 sales, beats Q4 estimates by 12%")
	ui.WriteBytesASCII(2, 41, "15:39 | Tesla announces new Gigafactory in Southeast Asia, stock up 5%")
	ui.WriteBytesASCII(2, 42, "15:38 | Oil prices surge on Middle East tensions, WTI up 3.2%")
	
	// System Messages
	ui.WriteBytesASCII(2, 44, "SYSTEM MESSAGES")
	ui.WriteBytesASCII(2, 45, "[ALERT] Large block order detected: NVDA 50K shares at $456.50")
	ui.WriteBytesASCII(2, 46, "[INFO] Market volatility increased: VIX up 8% in last 30 minutes")
	ui.WriteBytesASCII(2, 47, "[WARN] Position limit approaching: AAPL position at 95% of limit")
	ui.WriteBytesASCII(2, 48, "[TRADE] Order filled: MSFT 200 shares @ $342.45")
	
	// Status Bar
	ui.DrawBox(0, 50, 119, 52)
	ui.WriteBytesASCII(2, 51, "Market Data: REAL-TIME | Orders: CONNECTED | Risk: MONITORING | Latency: 0.8ms | Last Update: 15:42:33.456")
	
	// Shortcuts Panel (Right side of status)
	ui.WriteBytesASCII(85, 51, "F11:FullScreen F12:Layout")
	
	// Output
	fmt.Println("💼 BLOOMBERG TERMINAL COMPLEX UI")
	fmt.Println("===============================")
	fmt.Println("Professional trading workstation interface")
	fmt.Println()
	fmt.Println(ui.String())
	
	// Save to file
	content := fmt.Sprintf("BLOOMBERG TERMINAL COMPLEX UI\n=============================\n\nProfessional Trading Workstation Interface\nGenerated by Kit4AI Ultra-Wide UI System\n\nFeatures:\n- Real-time market data\n- Multi-panel layout\n- Order management\n- Portfolio tracking\n- Risk monitoring\n- News integration\n\n%s", ui.String())
	
	file, err := os.Create("complex_trading_ui.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
	} else {
		defer file.Close()
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		} else {
			fmt.Printf("\n💼 Complex trading UI saved: complex_trading_ui.txt\n")
		}
	}
}