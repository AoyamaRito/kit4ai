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
	fmt.Println("修正版テキストレイヤーテスト")
	fmt.Println("===========================")
	
	ls := canvas.NewLayerSystem()
	
	// 基準ボックス
	boxID := ls.AddLayerWithName("基準ボックス")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 1)
	drawBox(boxLayer.Canvas, 0, 0, 40, 15)
	
	// 位置マーカー
	markerID := ls.AddLayerWithName("位置マーカー")
	markerLayer, _ := ls.GetLayer(markerID)
	ls.SetZOrder(markerID, 2)
	
	for i := 0; i < 40; i++ {
		marker := fmt.Sprintf("%d", i%10)
		for j, r := range []rune(marker) {
			markerLayer.Canvas.ReplaceChar(i+j, 16, r)
		}
	}
	
	// テキストレイヤー
	textID := ls.AddTextLayer("テキストレイヤー")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// テスト1: 修正後のWriteTextWithWidth
	textLayer.WriteTextWithWidth(2, 2, "Hello")
	textLayer.WriteTextWithWidth(2, 3, "こんにちは")
	textLayer.WriteTextWithWidth(2, 4, "Hello世界")
	
	// テスト2: 中央寄せ（幅計算は維持）
	textLayer.WriteTextCentered(2, 6, 36, "中央寄せテスト")
	textLayer.WriteTextCentered(2, 7, 36, "Center Test")
	
	// テスト3: 右寄せ
	textLayer.WriteTextRight(2, 9, 36, "右寄せテスト")
	textLayer.WriteTextRight(2, 10, 36, "Right Test")
	
	// テスト4: 幅計算確認
	textLayer.WriteTextWithWidth(2, 12, "幅計算確認用")
	textLayer.WriteTextWithWidth(2, 13, "Width Check")
	
	result := ls.Composite()
	
	fmt.Println("修正結果:")
	fmt.Println(result.String())
	
	// 幅計算テスト
	fmt.Println("\n幅計算テスト:")
	testTexts := []string{
		"Hello",
		"こんにちは", 
		"Hello世界",
		"中央寄せテスト",
		"幅計算確認用",
	}
	
	for _, text := range testTexts {
		width := textLayer.GetTextWidth(text)
		runeCount := len([]rune(text))
		fmt.Printf("'%s' -> 表示幅:%d, rune数:%d\n", text, width, runeCount)
	}
	
	fmt.Println("\n修正内容:")
	fmt.Println("- WriteTextWithWidth: 連続配置（空白なし）")
	fmt.Println("- GetTextWidth: 幅計算は維持（レイアウト用）")
	fmt.Println("- 中央寄せ・右寄せ: 幅計算を使用して配置位置決定")
}