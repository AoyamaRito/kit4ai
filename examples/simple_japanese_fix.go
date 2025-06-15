package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
	"unicode"
)

func isFullWidth(r rune) bool {
	return unicode.Is(unicode.Han, r) ||
		unicode.Is(unicode.Hiragana, r) ||
		unicode.Is(unicode.Katakana, r) ||
		(r >= 0xFF01 && r <= 0xFF5E)
}

func calculateTextWidth(text string) int {
	width := 0
	for _, r := range []rune(text) {
		if isFullWidth(r) {
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

func writeTextSimple(c *canvas.Canvas, x, y int, text string) {
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
	fmt.Println("シンプル解決案テスト")
	fmt.Println("==================")
	
	ls := canvas.NewLayerSystem()
	
	// メインボックス
	boxID := ls.AddLayerWithName("メインボックス")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 1)
	drawBox(boxLayer.Canvas, 0, 0, 50, 15)
	
	// テキストレイヤー
	textID := ls.AddLayerWithName("テキスト")
	textLayer, _ := ls.GetLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// レイアウト計算してテキスト配置
	texts := []string{
		"ダッシュボード",
		"売上管理",
		"ユーザー統計", 
		"Hello World",
		"混在テストHello世界",
	}
	
	for i, text := range texts {
		width := calculateTextWidth(text)
		writeTextSimple(textLayer.Canvas, 2, 2+i*2, text)
		
		// 計算された幅を右側に表示
		widthText := fmt.Sprintf("(%d文字幅)", width)
		writeTextSimple(textLayer.Canvas, 25, 2+i*2, widthText)
	}
	
	// 罫線テスト
	writeTextSimple(textLayer.Canvas, 2, 13, "罫線テスト: +--+--+")
	
	result := ls.Composite()
	
	fmt.Println("結果:")
	fmt.Println(result.String())
	
	fmt.Println("\n文字幅計算テスト:")
	for _, text := range texts {
		fmt.Printf("'%s' -> %d文字幅\n", text, calculateTextWidth(text))
	}
	
	file, err := os.Create("simple_fix_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Simple Japanese Fix Test\n")
	file.WriteString("========================\n\n")
	file.WriteString("この方式では:\n")
	file.WriteString("- 全角文字は1runeとして配置（空白なし）\n")
	file.WriteString("- レイアウト計算時のみ2文字幅として扱う\n")
	file.WriteString("- 実際の表示では連続配置\n\n")
	file.WriteString(result.String())
	
	fmt.Println("\nsimple_fix_output.txtに出力しました")
	fmt.Println("この方式でレイアウト崩れが解決できそうですか？")
}