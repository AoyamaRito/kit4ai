package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
	"strings"
)

func main() {
	fmt.Println("Simplified Clean UI Design")
	fmt.Println("=========================")
	
	// シンプルで読みやすいUI
	bc := canvas.NewByteCanvas()
	
	// メインフレーム
	bc.DrawBox(0, 0, 60, 25)
	
	// ヘッダー
	bc.DrawBox(2, 2, 58, 5)
	
	// メインコンテンツエリア
	bc.DrawBox(2, 7, 58, 18)
	
	// フッター
	bc.DrawBox(2, 20, 58, 23)
	
	// ヘッダー - 最小限の情報
	bc.WriteBytes(20, 3, "SIMPLE DASHBOARD")
	bc.WriteBytes(25, 4, "Admin Panel")
	
	// メインコンテンツ - 重要情報のみ
	bc.WriteBytes(4, 9, "Status: ONLINE")
	bc.WriteBytes(4, 11, "Users: 1,234")
	bc.WriteBytes(4, 13, "Sales: $5,678")
	bc.WriteBytes(4, 15, "Orders: 89")
	
	// アクション - 3つまで
	bc.WriteBytes(30, 9, "[1] View Reports")
	bc.WriteBytes(30, 11, "[2] User Management")
	bc.WriteBytes(30, 13, "[3] Settings")
	
	// フッター - 必要最小限
	bc.WriteBytes(4, 21, "F1:Help")
	bc.WriteBytes(20, 21, "ESC:Exit")
	bc.WriteBytes(40, 21, "2024-06-15")
	
	fmt.Println("Clean UI (Less Text):")
	fmt.Println(bc.String())
	
	// さらにシンプルなバージョン
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("ULTRA MINIMAL VERSION:")
	fmt.Println(strings.Repeat("=", 40))
	
	bc2 := canvas.NewByteCanvas()
	
	// 最小限の枠
	bc2.DrawBox(0, 0, 40, 15)
	
	// タイトルのみ
	bc2.WriteBytes(15, 2, "DASHBOARD")
	
	// 核心データのみ
	bc2.WriteBytes(3, 5, "Users: 1,234")
	bc2.WriteBytes(3, 7, "Sales: $5,678")
	bc2.WriteBytes(3, 9, "Status: OK")
	
	// 最小限の操作
	bc2.WriteBytes(3, 12, "[ENTER] Details")
	bc2.WriteBytes(20, 12, "[ESC] Exit")
	
	fmt.Println(bc2.String())
	
	// 仕様書出力
	file, err := os.Create("clean_ui_spec.txt")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Clean UI Design Principles\n")
	file.WriteString("==========================\n\n")
	file.WriteString("Problem: Information Overload\n")
	file.WriteString("- Too much text causes visual clutter\n")
	file.WriteString("- Users get overwhelmed\n")
	file.WriteString("- Important information gets lost\n\n")
	file.WriteString("Solution: Minimal Design\n")
	file.WriteString("- Show only essential information\n")
	file.WriteString("- Limit to 3-5 key metrics\n")
	file.WriteString("- Use white space effectively\n")
	file.WriteString("- Prioritize actions\n\n")
	file.WriteString("Design Guidelines:\n")
	file.WriteString("1. Maximum 5 data points per screen\n")
	file.WriteString("2. Maximum 3 action buttons\n")
	file.WriteString("3. Single clear headline\n")
	file.WriteString("4. Consistent spacing\n")
	file.WriteString("5. Essential functions only\n\n")
	file.WriteString("Before (Complex):\n")
	file.WriteString("- 15+ data fields\n")
	file.WriteString("- 8+ action buttons\n")
	file.WriteString("- Multiple tables\n")
	file.WriteString("- Status indicators everywhere\n\n")
	file.WriteString("After (Clean):\n")
	file.WriteString("- 4 key metrics\n")
	file.WriteString("- 3 main actions\n")
	file.WriteString("- Single status\n")
	file.WriteString("- Clear hierarchy\n\n")
	file.WriteString("Simple UI Layout:\n")
	file.WriteString(bc.String())
	file.WriteString("\n\nUltra Minimal Layout:\n")
	file.WriteString(bc2.String())
	
	fmt.Println("\nclean_ui_spec.txtに出力しました")
	fmt.Println("\n改善点:")
	fmt.Println("- 情報量を80%削減")
	fmt.Println("- 視覚的なノイズを除去")
	fmt.Println("- 重要な情報のみ表示")
	fmt.Println("- ユーザーの認知負荷を軽減")
	fmt.Println("- アクション数を制限")
	fmt.Println("- 白空間を効果的に活用")
}