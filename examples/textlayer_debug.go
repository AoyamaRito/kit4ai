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
	fmt.Println("テキストレイヤーのずれ問題デバッグ")
	fmt.Println("===============================")
	
	ls := canvas.NewLayerSystem()
	
	// 基準ボックス
	boxID := ls.AddLayerWithName("基準ボックス")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 1)
	drawBox(boxLayer.Canvas, 0, 0, 30, 10)
	
	// 位置マーカー
	markerID := ls.AddLayerWithName("位置マーカー")
	markerLayer, _ := ls.GetLayer(markerID)
	ls.SetZOrder(markerID, 2)
	
	// 0-9の位置マーカーを配置
	for i := 0; i < 30; i++ {
		marker := fmt.Sprintf("%d", i%10)
		for j, r := range []rune(marker) {
			markerLayer.Canvas.ReplaceChar(i+j, 11, r)
		}
	}
	
	// テキストレイヤー
	textID := ls.AddTextLayer("テキストレイヤー")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// テスト1: 従来方式（WriteTextSimple）
	textLayer.WriteTextSimple(2, 2, "Simple: Hello")
	textLayer.WriteTextSimple(2, 3, "Simple: こんにちは")
	
	// テスト2: 新方式（WriteTextWithWidth）
	textLayer.WriteTextWithWidth(2, 5, "Width: Hello")
	textLayer.WriteTextWithWidth(2, 6, "Width: こんにちは")
	
	// テスト3: 中央寄せ
	textLayer.WriteTextCentered(2, 8, 26, "Center: テスト")
	
	result := ls.Composite()
	
	fmt.Println("結果:")
	fmt.Println(result.String())
	
	// 詳細分析
	fmt.Println("\n詳細分析:")
	
	testTexts := []string{"Hello", "こんにちは", "テスト"}
	for _, text := range testTexts {
		width := textLayer.GetTextWidth(text)
		fmt.Printf("'%s' -> 計算幅: %d文字\n", text, width)
		
		// 実際の文字数
		runeCount := len([]rune(text))
		fmt.Printf("  -> rune数: %d\n", runeCount)
	}
	
	// WriteTextWithWidthの動作確認
	fmt.Println("\nWriteTextWithWidth の動作:")
	text := "こんにちは"
	fmt.Printf("テキスト: %s\n", text)
	
	currentX := 2
	for i, r := range []rune(text) {
		isFullWidth := textLayer.IsFullWidth(r)
		charWidth := textLayer.GetCharWidth(r)
		
		fmt.Printf("位置%d: '%c' -> 全角:%v, 幅:%d, 配置位置:%d\n", 
			i, r, isFullWidth, charWidth, currentX)
		
		if isFullWidth {
			currentX += 2
		} else {
			currentX += 1
		}
	}
}