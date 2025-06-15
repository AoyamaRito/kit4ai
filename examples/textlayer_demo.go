package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
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
	fmt.Println("テキストレイヤー専用全角対応デモ")
	fmt.Println("==============================")
	
	ls := canvas.NewLayerSystem()
	
	// 通常レイヤー（枠組み用）
	boxID := ls.AddLayerWithName("メインボックス")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 1)
	drawBox(boxLayer.Canvas, 0, 0, 60, 20)
	
	// カード1
	card1ID := ls.AddLayerWithName("カード1")
	card1Layer, _ := ls.GetLayer(card1ID)
	ls.SetZOrder(card1ID, 2)
	drawBox(card1Layer.Canvas, 2, 2, 28, 8)
	
	// カード2
	card2ID := ls.AddLayerWithName("カード2")
	card2Layer, _ := ls.GetLayer(card2ID)
	ls.SetZOrder(card2ID, 3)
	drawBox(card2Layer.Canvas, 32, 2, 58, 8)
	
	// グラフエリア
	graphID := ls.AddLayerWithName("グラフエリア")
	graphLayer, _ := ls.GetLayer(graphID)
	ls.SetZOrder(graphID, 4)
	drawBox(graphLayer.Canvas, 2, 10, 58, 18)
	
	// テキストレイヤー（全角対応）
	textID := ls.AddTextLayer("テキストレイヤー")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// ヘッダータイトル（幅を考慮した配置）
	textLayer.WriteTextCentered(2, 3, 26, "売上統計")
	textLayer.WriteTextCentered(32, 3, 26, "ユーザー分析")
	
	// 数値データ（右寄せ）
	textLayer.WriteTextCentered(2, 5, 26, "¥2,500,000")
	textLayer.WriteTextCentered(32, 5, 26, "1,234人")
	
	// 変化率（中央寄せ）
	textLayer.WriteTextCentered(2, 6, 26, "前月比 +15.2%")
	textLayer.WriteTextCentered(32, 6, 26, "前月比 +8.7%")
	
	// グラフタイトル
	textLayer.WriteTextCentered(2, 11, 56, "月別売上推移グラフ")
	
	// 簡単なグラフデータ
	months := []string{"1月", "2月", "3月", "4月", "5月", "6月"}
	values := []string{"70%", "85%", "92%", "88%", "95%", "100%"}
	bars := []string{"[======----]", "[========--]", "[==========-]", "[========---]", "[==========]", "[==========]"}
	
	for i, month := range months {
		y := 13 + i
		textLayer.WriteTextSimple(4, y, month)
		textLayer.WriteTextSimple(8, y, bars[i])
		textLayer.WriteTextSimple(22, y, values[i])
	}
	
	// 比較テスト用の説明
	fmt.Println("文字幅計算テスト:")
	testTexts := []string{
		"売上統計",
		"ユーザー分析", 
		"月別売上推移グラフ",
		"Hello World",
	}
	
	for _, text := range testTexts {
		width := textLayer.GetTextWidth(text)
		fmt.Printf("'%s' -> %d文字幅\n", text, width)
	}
	
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n完成したダッシュボード（テキストレイヤー全角対応版）:")
	fmt.Println(result.String())
	
	// 従来版との比較
	fmt.Println("\n=== 比較テスト ===")
	
	// 従来版のテキストレイヤー
	normalTextID := ls.AddLayerWithName("通常テキスト")
	normalTextLayer, _ := ls.GetLayer(normalTextID)
	ls.SetZOrder(normalTextID, 50)
	
	// 同じテキストを従来方式で配置
	writeTextOld := func(c *canvas.Canvas, x, y int, text string) {
		for i, r := range []rune(text) {
			c.ReplaceChar(x+i, y, r)
		}
	}
	
	writeTextOld(normalTextLayer.Canvas, 2, 1, "従来版: 売上統計")
	
	// 新しいテキストレイヤーでも配置
	textLayer.WriteTextSimple(2, 0, "新版: 売上統計")
	
	// ファイル出力
	file, err := os.Create("textlayer_demo_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("TextLayer Demo - Full-width Character Support\n")
	file.WriteString("=============================================\n\n")
	file.WriteString("Features:\n")
	file.WriteString("- Dedicated text layer with full-width character support\n")
	file.WriteString("- Frame layers remain simple (1 char = 1 grid)\n")
	file.WriteString("- Text layer handles 2-char width calculation\n")
	file.WriteString("- Centered and right-aligned text functions\n\n")
	
	for _, text := range testTexts {
		width := textLayer.GetTextWidth(text)
		file.WriteString(fmt.Sprintf("'%s' -> %d chars width\n", text, width))
	}
	
	file.WriteString("\nASCII Art Layout:\n")
	file.WriteString(result.String())
	
	fmt.Println("\ntextlayer_demo_output.txtに出力しました")
}