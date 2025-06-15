package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
	"strings"
)

func createDashboard(configName string) string {
	bc := canvas.NewByteCanvas()
	
	// メインフレーム（現在の設定幅に合わせて）
	bc.DrawBox(0, 0, bc.Width-1, 15)
	
	// ヘッダー
	bc.DrawBox(2, 2, bc.Width-3, 4)
	
	// コンテンツエリア
	bc.DrawBox(2, 6, bc.Width-3, 12)
	
	// ヘッダーテキスト
	headerText := fmt.Sprintf("DASHBOARD - %s", configName)
	bc.WriteBytes(4, 3, headerText)
	
	// データ表示（幅に応じて調整）
	bc.WriteBytes(4, 8, "Users: 1,234")
	bc.WriteBytes(20, 8, "Sales: $5,678")
	
	if bc.Width >= 80 {
		bc.WriteBytes(40, 8, "Orders: 89")
		bc.WriteBytes(55, 8, "Status: OK")
	}
	
	if bc.Width >= 100 {
		bc.WriteBytes(70, 8, "More: Data")
		bc.WriteBytes(85, 8, "Extra: Info")
	}
	
	if bc.Width >= 120 {
		bc.WriteBytes(100, 8, "Ultra: Wide")
		bc.WriteBytes(115, 8, "Max: Space")
	}
	
	// 幅情報表示
	widthInfo := fmt.Sprintf("Canvas Width: %d chars", bc.Width)
	bc.WriteBytes(4, 10, widthInfo)
	
	return bc.String()
}

func main() {
	fmt.Println("設定可能な横幅システムデモ")
	fmt.Println("==========================")
	
	// 設定一覧を表示
	fmt.Println("利用可能な設定:")
	fmt.Println("1. Standard (80x100) - Legacy Compatible")
	fmt.Println("2. Wide (100x100) - Modern Display") 
	fmt.Println("3. Ultra-Wide (120x100) - Large Monitor")
	fmt.Println("4. Compact (60x80) - Mobile/Narrow")
	fmt.Println("5. Print (72x90) - A4 Paper Friendly")
	fmt.Println("6. Custom - User Defined")
	
	// 各設定でダッシュボードを作成
	configs := []struct {
		name     string
		setFunc  func()
		fileName string
	}{
		{"Standard", canvas.SetStandardWidth, "dashboard_standard.txt"},
		{"Wide", canvas.SetWideWidth, "dashboard_wide.txt"},
		{"Ultra-Wide", canvas.SetUltraWideWidth, "dashboard_ultrawide.txt"},
		{"Compact", canvas.SetCompactWidth, "dashboard_compact.txt"},
		{"Print", canvas.SetPrintWidth, "dashboard_print.txt"},
	}
	
	for _, config := range configs {
		fmt.Printf("\n%s\n", strings.Repeat("=", 50))
		fmt.Printf("%s 設定\n", config.name)
		fmt.Printf("%s\n", strings.Repeat("=", 50))
		
		// 設定を変更
		config.setFunc()
		
		// 現在の設定情報を表示
		fmt.Printf("設定: %s\n", canvas.GetConfigName())
		fmt.Printf("サイズ: %dx%d\n", canvas.GetCurrentWidth(), canvas.GetCurrentHeight())
		
		// ダッシュボードを作成
		dashboard := createDashboard(config.name)
		fmt.Println(dashboard)
		
		// ファイルに出力
		file, err := os.Create(config.fileName)
		if err != nil {
			fmt.Printf("ファイル作成エラー: %v\n", err)
			continue
		}
		
		file.WriteString(fmt.Sprintf("%s Configuration Dashboard\n", config.name))
		file.WriteString(strings.Repeat("=", 40) + "\n\n")
		file.WriteString(fmt.Sprintf("Configuration: %s\n", canvas.GetConfigName()))
		file.WriteString(fmt.Sprintf("Dimensions: %dx%d\n\n", canvas.GetCurrentWidth(), canvas.GetCurrentHeight()))
		file.WriteString("Layout:\n")
		file.WriteString(dashboard)
		
		file.Close()
		fmt.Printf("-> %s に出力しました\n", config.fileName)
	}
	
	// カスタム設定のデモ
	fmt.Printf("\n%s\n", strings.Repeat("=", 50))
	fmt.Println("カスタム設定デモ")
	fmt.Printf("%s\n", strings.Repeat("=", 50))
	
	// カスタムサイズ設定
	canvas.SetCustomConfig(90, 50)
	fmt.Printf("設定: %s\n", canvas.GetConfigName())
	
	customDashboard := createDashboard("Custom 90x50")
	fmt.Println(customDashboard)
	
	// カスタム設定ファイル出力
	file, err := os.Create("dashboard_custom.txt")
	if err == nil {
		file.WriteString("Custom Configuration Dashboard\n")
		file.WriteString(strings.Repeat("=", 40) + "\n\n")
		file.WriteString(fmt.Sprintf("Configuration: %s\n", canvas.GetConfigName()))
		file.WriteString("Layout:\n")
		file.WriteString(customDashboard)
		file.Close()
		fmt.Println("-> dashboard_custom.txt に出力しました")
	}
	
	fmt.Println("\n機能:")
	fmt.Println("- 5つのプリセット設定")
	fmt.Println("- カスタムサイズ対応")
	fmt.Println("- 動的な幅調整")
	fmt.Println("- 用途別最適化")
	fmt.Println("- 設定変更の簡単切り替え")
}