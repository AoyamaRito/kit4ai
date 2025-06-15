package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
)

func drawBox(c *canvas.Canvas, x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		c.ReplaceChar(x, y1, '-')
		c.ReplaceChar(x, y2, '-')
	}
	
	for y := y1; y <= y2; y++ {
		c.ReplaceChar(x1, y, '|')
		c.ReplaceChar(x2, y, '|')
	}
	
	c.ReplaceChar(x1, y1, '+')
	c.ReplaceChar(x2, y1, '+')
	c.ReplaceChar(x1, y2, '+')
	c.ReplaceChar(x2, y2, '+')
}

func main() {
	fmt.Println("文字列処理によるずれ問題の調査")
	fmt.Println("=============================")
	
	ls := canvas.NewLayerSystem()
	
	// 基準ボックス
	boxID := ls.AddLayerWithName("基準ボックス")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 1)
	drawBox(boxLayer.Canvas, 10, 5, 30, 15)
	
	// テキストレイヤー
	textID := ls.AddTextLayer("テキスト")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// 問題の調査: 複数行文字列の処理
	problemString := "┌──────────────────────────────┐"
	fmt.Printf("問題の文字列: %s\n", problemString)
	fmt.Printf("文字列長: %d\n", len(problemString))
	fmt.Printf("rune数: %d\n", len([]rune(problemString)))
	
	// 各文字の詳細分析
	fmt.Println("\n各文字の分析:")
	for i, r := range []rune(problemString) {
		fmt.Printf("位置%d: '%c' (U+%04X)\n", i, r, r)
	}
	
	// WriteTextWithWidthでの配置
	textLayer.WriteTextWithWidth(12, 7, "単純テキスト")
	textLayer.WriteTextWithWidth(12, 8, problemString)
	textLayer.WriteTextWithWidth(12, 9, "│ 内容テキスト │")
	textLayer.WriteTextWithWidth(12, 10, "└──────────────────────────────┘")
	
	// 1文字ずつ手動配置で比較
	manualID := ls.AddLayerWithName("手動配置")
	manualLayer, _ := ls.GetLayer(manualID)
	ls.SetZOrder(manualID, 50)
	
	// 手動で1文字ずつ配置
	chars := []rune("┌──────┐")
	for i, char := range chars {
		manualLayer.Canvas.ReplaceChar(12+i, 12, char)
	}
	
	chars2 := []rune("│テスト│")
	for i, char := range chars2 {
		manualLayer.Canvas.ReplaceChar(12+i, 13, char)
	}
	
	chars3 := []rune("└──────┘")
	for i, char := range chars3 {
		manualLayer.Canvas.ReplaceChar(12+i, 14, char)
	}
	
	result := ls.Composite()
	
	fmt.Println("\n結果比較:")
	fmt.Println("（WriteTextWithWidth vs 手動配置）")
	fmt.Println(result.String())
	
	// 具体的な問題の確認
	fmt.Println("\n問題の詳細分析:")
	fmt.Println("1. Unicode罫線文字が使われている")
	fmt.Println("2. 各罫線文字の幅が統一されていない可能性")
	fmt.Println("3. WriteTextWithWidthでの連続配置に問題")
	
	// 安全な文字での再テスト
	fmt.Println("\n安全な文字でのテスト:")
	textLayer.WriteTextWithWidth(12, 3, "+----------+")
	textLayer.WriteTextWithWidth(12, 4, "| 安全文字 |")
	textLayer.WriteTextWithWidth(12, 5, "+----------+")
	
	// 文字幅計算の確認
	fmt.Println("\n文字幅計算:")
	testStrings := []string{
		"+----------+",
		"┌──────────┐",
		"│テストです│",
		"| 安全文字 |",
	}
	
	for _, str := range testStrings {
		width := textLayer.GetTextWidth(str)
		runeCount := len([]rune(str))
		fmt.Printf("'%s' -> 計算幅:%d, rune数:%d\n", str, width, runeCount)
	}
}