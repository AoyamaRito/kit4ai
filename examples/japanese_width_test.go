package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"unicode"
)

func writeTextWithWidthCheck(c *canvas.Canvas, x, y int, text string) int {
	currentX := x
	textRunes := []rune(text)
	
	for _, r := range textRunes {
		// 全角文字かどうかチェック
		isFullWidth := unicode.Is(unicode.Han, r) || 
			unicode.Is(unicode.Hiragana, r) || 
			unicode.Is(unicode.Katakana, r) ||
			(r >= 0xFF01 && r <= 0xFF5E) // 全角英数字記号
		
		if isFullWidth {
			// 全角文字は2文字分のスペースを使用
			c.ReplaceChar(currentX, y, r)
			if currentX+1 < 80 {
				c.ReplaceChar(currentX+1, y, ' ') // 次の位置は空白で埋める
			}
			currentX += 2
		} else {
			// 半角文字は1文字分
			c.ReplaceChar(currentX, y, r)
			currentX += 1
		}
	}
	
	return currentX - x // 使用した文字数を返す
}

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
	fmt.Println("日本語文字幅テスト")
	fmt.Println("================")
	
	ls := canvas.NewLayerSystem()
	
	// 枠レイヤー
	boxID := ls.AddLayerWithName("枠")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 1)
	drawBox(boxLayer.Canvas, 0, 0, 40, 10)
	
	// テキストレイヤー
	textID := ls.AddLayerWithName("テキスト")
	textLayer, _ := ls.GetLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// テスト用テキスト
	testTexts := []string{
		"Hello World",           // 半角のみ
		"こんにちは",                // ひらがなのみ
		"コンニチハ",                // カタカナのみ  
		"日本語テスト",               // 漢字混じり
		"Hello日本語World",      // 混在
		"１２３４５",                // 全角数字
		"ＡＢＣ",                  // 全角英字
	}
	
	fmt.Println("テキスト幅計算:")
	for i, text := range testTexts {
		width := writeTextWithWidthCheck(textLayer.Canvas, 2, 2+i, text)
		fmt.Printf("'%s' -> %d文字幅\n", text, width)
	}
	
	result := ls.Composite()
	
	fmt.Println("\n結果 (枠内に正しく収まっているかチェック):")
	fmt.Println(result.String())
	
	// 参考: rune単位でのテスト
	fmt.Println("\n各文字の詳細分析:")
	testString := "Hello日本語123"
	fmt.Printf("テスト文字列: %s\n", testString)
	
	for i, r := range []rune(testString) {
		isFullWidth := unicode.Is(unicode.Han, r) || 
			unicode.Is(unicode.Hiragana, r) || 
			unicode.Is(unicode.Katakana, r) ||
			(r >= 0xFF01 && r <= 0xFF5E)
		
		width := 1
		if isFullWidth {
			width = 2
		}
		
		fmt.Printf("位置%d: '%c' (U+%04X) -> %d文字幅\n", i, r, r, width)
	}
}