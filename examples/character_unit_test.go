package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"unicode"
)

func writeText(c *canvas.Canvas, x, y int, text string) {
	for i, r := range []rune(text) {
		c.ReplaceChar(x+i, y, r)
	}
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
	fmt.Println("文字単位テスト")
	fmt.Println("=============")
	
	c := canvas.NewCanvas()
	
	// テスト1: 半角文字
	writeText(c, 0, 0, "Hello")
	
	// テスト2: 全角文字
	writeText(c, 0, 2, "こんにちは")
	
	// テスト3: 混在
	writeText(c, 0, 4, "Hello世界")
	
	// テスト4: 位置確認用のマーカー
	for i := 0; i < 20; i++ {
		marker := fmt.Sprintf("%d", i%10)
		writeText(c, i, 6, marker)
	}
	
	// テスト5: ボックステスト
	drawBox(c, 0, 8, 10, 12)
	writeText(c, 1, 9, "日本語")
	writeText(c, 1, 10, "English")
	
	result := c.String()
	fmt.Println("結果:")
	fmt.Println(result)
	
	fmt.Println("\n分析:")
	fmt.Println("- 現在の実装は1文字=1グリッド位置")
	fmt.Println("- 全角文字も半角文字も同じ1つのグリッドに配置")
	fmt.Println("- 実際の表示では全角文字は2文字分の幅を占める")
	
	// 詳細分析
	fmt.Println("\n詳細分析:")
	testString := "Hello世界"
	fmt.Printf("文字列: %s\n", testString)
	
	for i, r := range []rune(testString) {
		isFullWidth := unicode.Is(unicode.Han, r) ||
			unicode.Is(unicode.Hiragana, r) ||
			unicode.Is(unicode.Katakana, r)
		
		displayWidth := 1
		if isFullWidth {
			displayWidth = 2
		}
		
		fmt.Printf("位置%d: '%c' -> グリッド1つ, 表示幅%d\n", i, r, displayWidth)
	}
	
	fmt.Println("\n課題:")
	fmt.Println("全角文字が隣接すると表示が重なる可能性があります")
}