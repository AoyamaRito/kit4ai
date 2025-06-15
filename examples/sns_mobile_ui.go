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

func drawRoundedBox(c *canvas.Canvas, x1, y1, x2, y2 int) {
	// 角を丸くしたボックス風
	for x := x1+1; x < x2; x++ {
		c.ReplaceChar(x, y1, '-')
		c.ReplaceChar(x, y2, '-')
	}
	
	for y := y1+1; y < y2; y++ {
		c.ReplaceChar(x1, y, '|')
		c.ReplaceChar(x2, y, '|')
	}
	
	// 丸い角の表現
	c.ReplaceChar(x1, y1, '.')
	c.ReplaceChar(x2, y1, '.')
	c.ReplaceChar(x1, y2, '`')
	c.ReplaceChar(x2, y2, '\'')
}

func main() {
	fmt.Println("SNSモバイルアプリUI仕様書")
	fmt.Println("========================")
	
	ls := canvas.NewLayerSystem()
	
	// スマホ外枠
	phoneID := ls.AddLayerWithName("スマホ外枠")
	phoneLayer, _ := ls.GetLayer(phoneID)
	ls.SetZOrder(phoneID, 1)
	drawBox(phoneLayer.Canvas, 0, 0, 35, 50)
	
	// ステータスバー
	statusID := ls.AddLayerWithName("ステータスバー")
	statusLayer, _ := ls.GetLayer(statusID)
	ls.SetZOrder(statusID, 2)
	drawBox(statusLayer.Canvas, 1, 1, 34, 3)
	
	// ヘッダー
	headerID := ls.AddLayerWithName("ヘッダー")
	headerLayer, _ := ls.GetLayer(headerID)
	ls.SetZOrder(headerID, 3)
	drawBox(headerLayer.Canvas, 1, 4, 34, 7)
	
	// 投稿フィード
	feed1ID := ls.AddLayerWithName("投稿1")
	feed1Layer, _ := ls.GetLayer(feed1ID)
	ls.SetZOrder(feed1ID, 4)
	drawRoundedBox(feed1Layer.Canvas, 2, 8, 33, 16)
	
	feed2ID := ls.AddLayerWithName("投稿2")
	feed2Layer, _ := ls.GetLayer(feed2ID)
	ls.SetZOrder(feed2ID, 5)
	drawRoundedBox(feed2Layer.Canvas, 2, 17, 33, 25)
	
	feed3ID := ls.AddLayerWithName("投稿3")
	feed3Layer, _ := ls.GetLayer(feed3ID)
	ls.SetZOrder(feed3ID, 6)
	drawRoundedBox(feed3Layer.Canvas, 2, 26, 33, 34)
	
	// ストーリーズエリア
	storiesID := ls.AddLayerWithName("ストーリーズ")
	storiesLayer, _ := ls.GetLayer(storiesID)
	ls.SetZOrder(storiesID, 7)
	
	// 小さな円でストーリー表現
	for i := 0; i < 5; i++ {
		x := 4 + i*6
		storiesLayer.Canvas.ReplaceChar(x, 35, '(')
		storiesLayer.Canvas.ReplaceChar(x+1, 35, 'o')
		storiesLayer.Canvas.ReplaceChar(x+2, 35, ')')
	}
	
	// ボトムナビ
	navID := ls.AddLayerWithName("ボトムナビ")
	navLayer, _ := ls.GetLayer(navID)
	ls.SetZOrder(navID, 8)
	drawBox(navLayer.Canvas, 1, 46, 34, 49)
	
	// テキストレイヤー
	textID := ls.AddLayerWithName("テキスト")
	textLayer, _ := ls.GetLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// ステータスバーテキスト
	writeText(textLayer.Canvas, 3, 2, "9:41")
	writeText(textLayer.Canvas, 28, 2, "100%")
	
	// ヘッダーテキスト
	writeText(textLayer.Canvas, 3, 5, "Timeline")
	writeText(textLayer.Canvas, 28, 5, "🔍 ⚙")
	
	// 投稿内容
	writeText(textLayer.Canvas, 4, 9, "👤 田中太郎")
	writeText(textLayer.Canvas, 4, 10, "今日の昼食は美味しいラーメンでした！")
	writeText(textLayer.Canvas, 4, 11, "#ラーメン #美味しい")
	writeText(textLayer.Canvas, 4, 13, "❤ 24  💬 5  🔄 2")
	writeText(textLayer.Canvas, 4, 14, "2時間前")
	
	writeText(textLayer.Canvas, 4, 18, "👤 佐藤花子")
	writeText(textLayer.Canvas, 4, 19, "新しいカフェを発見！")
	writeText(textLayer.Canvas, 4, 20, "とても素敵な雰囲気でした✨")
	writeText(textLayer.Canvas, 4, 22, "❤ 42  💬 8  🔄 3")
	writeText(textLayer.Canvas, 4, 23, "5時間前")
	
	writeText(textLayer.Canvas, 4, 27, "👤 鈴木一郎")
	writeText(textLayer.Canvas, 4, 28, "週末の予定どうしようかな")
	writeText(textLayer.Canvas, 4, 29, "映画でも見に行こうかと思案中")
	writeText(textLayer.Canvas, 4, 31, "❤ 8   💬 12 🔄 1")
	writeText(textLayer.Canvas, 4, 32, "1日前")
	
	// ストーリーズラベル
	writeText(textLayer.Canvas, 4, 36, "Stories")
	
	// ボトムナビゲーション
	writeText(textLayer.Canvas, 4, 47, "🏠")
	writeText(textLayer.Canvas, 10, 47, "🔍")
	writeText(textLayer.Canvas, 16, 47, "➕")
	writeText(textLayer.Canvas, 22, 47, "❤")
	writeText(textLayer.Canvas, 28, 47, "👤")
	
	writeText(textLayer.Canvas, 3, 48, "Home")
	writeText(textLayer.Canvas, 8, 48, "Search")
	writeText(textLayer.Canvas, 15, 48, "Post")
	writeText(textLayer.Canvas, 21, 48, "Like")
	writeText(textLayer.Canvas, 26, 48, "Profile")
	
	fmt.Println("レイヤー構成:")
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\nSNSモバイルアプリUI:")
	fmt.Println(result.String())
	
	// 仕様書として出力
	file, err := os.Create("sns_mobile_ui_spec.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("SNS Mobile App UI Specification\n")
	file.WriteString("===============================\n\n")
	file.WriteString("Screen Size: 36x51 characters (mobile portrait)\n\n")
	file.WriteString("Components:\n")
	file.WriteString("1. Status Bar (time, battery)\n")
	file.WriteString("2. Header (title, search, settings)\n")
	file.WriteString("3. Post Feed (user posts with like/comment/share)\n")
	file.WriteString("4. Stories Section\n")
	file.WriteString("5. Bottom Navigation (5 tabs)\n\n")
	file.WriteString("Layer Structure:\n")
	
	for _, id := range ls.GetLayerIDs() {
		layer, _ := ls.GetLayer(id)
		file.WriteString(fmt.Sprintf("- %s (Z-Order: %d)\n", layer.Name, layer.ZOrder))
	}
	
	file.WriteString("\nASCII Art Layout:\n")
	file.WriteString(result.String())
	
	fmt.Println("\nsns_mobile_ui_spec.txtに仕様書を出力しました")
	fmt.Println("\n特徴:")
	fmt.Println("- モバイル縦画面のSNSアプリレイアウト")
	fmt.Println("- 投稿フィード、ストーリーズ、ナビゲーション")
	fmt.Println("- 日本語テキスト対応")
	fmt.Println("- 絵文字でリアルなUI表現")
}