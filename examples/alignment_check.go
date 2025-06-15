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
	fmt.Println("位置ずれ確認テスト")
	fmt.Println("==================")
	
	ls := canvas.NewLayerSystem()
	
	// 基準グリッド
	gridID := ls.AddLayerWithName("グリッド")
	gridLayer, _ := ls.GetLayer(gridID)
	ls.SetZOrder(gridID, 1)
	
	// 10文字ごとの縦線
	for x := 0; x < 80; x += 10 {
		for y := 0; y < 30; y++ {
			gridLayer.Canvas.ReplaceChar(x, y, '|')
		}
	}
	
	// 5行ごとの横線
	for y := 0; y < 30; y += 5 {
		for x := 0; x < 80; x++ {
			gridLayer.Canvas.ReplaceChar(x, y, '-')
		}
	}
	
	// 位置番号
	for x := 0; x < 80; x += 10 {
		for y := 0; y < 30; y += 5 {
			num := fmt.Sprintf("%d", x/10)
			gridLayer.Canvas.ReplaceChar(x, y, rune(num[0]))
		}
	}
	
	// テストボックス
	boxID := ls.AddLayerWithName("テストボックス")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 2)
	drawBox(boxLayer.Canvas, 10, 10, 40, 20)
	
	// テキストレイヤー
	textID := ls.AddTextLayer("テキスト")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// 位置確認テキスト
	textLayer.WriteTextWithWidth(12, 12, "Hello")        // 半角のみ
	textLayer.WriteTextWithWidth(12, 13, "こんにちは")        // 全角のみ
	textLayer.WriteTextWithWidth(12, 14, "Hello世界")     // 混在
	
	// 中央寄せテスト
	textLayer.WriteTextCentered(10, 16, 30, "中央寄せ")    // ボックス内中央
	textLayer.WriteTextCentered(10, 17, 30, "Center")   // 半角中央寄せ
	
	// 右寄せテスト
	textLayer.WriteTextRight(10, 18, 30, "右寄せ")       // ボックス内右寄せ
	textLayer.WriteTextRight(10, 19, 30, "Right")      // 半角右寄せ
	
	result := ls.Composite()
	
	fmt.Println("位置ずれ確認結果:")
	fmt.Println("（縦線は10文字間隔、横線は5行間隔）")
	fmt.Println(result.String())
	
	// 詳細分析
	fmt.Println("\n詳細分析:")
	
	tests := []struct{
		text string
		func_name string
	}{
		{"Hello", "半角のみ"},
		{"こんにちは", "全角のみ"},
		{"Hello世界", "混在"},
		{"中央寄せ", "全角中央寄せ"},
		{"右寄せ", "全角右寄せ"},
	}
	
	for _, test := range tests {
		width := textLayer.GetTextWidth(test.text)
		runeCount := len([]rune(test.text))
		fmt.Printf("'%s' (%s) -> 計算幅:%d, rune数:%d\n", 
			test.text, test.func_name, width, runeCount)
	}
	
	// 中央寄せの計算確認
	fmt.Println("\n中央寄せ計算確認:")
	text := "中央寄せ"
	boxWidth := 30
	textWidth := textLayer.GetTextWidth(text)
	offset := (boxWidth - textWidth) / 2
	startPos := 10 + offset
	
	fmt.Printf("ボックス幅: %d\n", boxWidth)
	fmt.Printf("テキスト幅: %d\n", textWidth)
	fmt.Printf("オフセット: %d\n", offset)
	fmt.Printf("開始位置: %d\n", startPos)
	fmt.Printf("期待される中央位置: %d\n", 10 + boxWidth/2)
}