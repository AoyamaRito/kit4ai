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

func writeText(c *canvas.Canvas, x, y int, text string) {
	textRunes := []rune(text)
	for i, r := range textRunes {
		c.ReplaceChar(x+i, y, r)
	}
}

func main() {
	ls := canvas.NewLayerSystem()
	
	// 背景レイヤー群 (Z-Order: 1-5)
	headerID := ls.AddLayerWithName("ヘッダー枠")
	headerLayer, _ := ls.GetLayer(headerID)
	ls.SetZOrder(headerID, 1)
	drawBox(headerLayer.Canvas, 0, 0, 50, 4)
	
	sidebarID := ls.AddLayerWithName("サイドバー枠")
	sidebarLayer, _ := ls.GetLayer(sidebarID)
	ls.SetZOrder(sidebarID, 2)
	drawBox(sidebarLayer.Canvas, 0, 0, 15, 15)
	ls.MoveLayer(sidebarID, 0, 5)
	
	mainID := ls.AddLayerWithName("メイン枠")
	mainLayer, _ := ls.GetLayer(mainID)
	ls.SetZOrder(mainID, 3)
	drawBox(mainLayer.Canvas, 0, 0, 30, 12)
	ls.MoveLayer(mainID, 17, 7)
	
	// 文字レイヤー (Z-Order: 100 - 最前面)
	textID := ls.AddLayerWithName("文字レイヤー")
	textLayer, _ := ls.GetLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// ヘッダーのタイトル
	writeText(textLayer.Canvas, 2, 2, "ヘッダータイトル")
	
	// サイドバーのメニュー項目
	writeText(textLayer.Canvas, 2, 7, "ホーム")
	writeText(textLayer.Canvas, 2, 9, "プロフィール")
	writeText(textLayer.Canvas, 2, 11, "設定")
	writeText(textLayer.Canvas, 2, 13, "ログアウト")
	
	// メインコンテンツ
	writeText(textLayer.Canvas, 19, 9, "メインコンテンツ")
	writeText(textLayer.Canvas, 19, 11, "投稿１：重要なお知らせ")
	writeText(textLayer.Canvas, 19, 13, "投稿２：新機能の紹介")
	writeText(textLayer.Canvas, 19, 15, "投稿３：今後の予定")
	
	fmt.Println("文字レイヤーオーバーレイのデモ")
	fmt.Println("=============================")
	
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n合成結果 (文字が最前面に表示):")
	fmt.Println(result.String())
	
	// 文字レイヤーだけを表示
	fmt.Println("\n文字レイヤーのみ:")
	fmt.Println(textLayer.Canvas.String())
	
	file, err := os.Create("text_overlay_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Text Overlay Demo\n")
	file.WriteString("=================\n\n")
	file.WriteString("文字レイヤー (Z-Order: 100) が最前面でオーバーレイ\n\n")
	file.WriteString(result.String())
	
	fmt.Println("\ntext_overlay_output.txtに出力しました")
}