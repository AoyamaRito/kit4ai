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
	ls := canvas.NewLayerSystem()
	
	// トップナビゲーション
	navID := ls.AddLayerWithName("トップナビ")
	navLayer, _ := ls.GetLayer(navID)
	ls.SetZOrder(navID, 1)
	drawBox(navLayer.Canvas, 0, 0, 79, 2)
	
	// 統計カード群
	card1ID := ls.AddLayerWithName("売上カード")
	card1Layer, _ := ls.GetLayer(card1ID)
	ls.SetZOrder(card1ID, 2)
	drawBox(card1Layer.Canvas, 0, 0, 18, 6)
	ls.MoveLayer(card1ID, 2, 4)
	
	card2ID := ls.AddLayerWithName("ユーザーカード")
	card2Layer, _ := ls.GetLayer(card2ID)
	ls.SetZOrder(card2ID, 3)
	drawBox(card2Layer.Canvas, 0, 0, 18, 6)
	ls.MoveLayer(card2ID, 22, 4)
	
	card3ID := ls.AddLayerWithName("注文カード")
	card3Layer, _ := ls.GetLayer(card3ID)
	ls.SetZOrder(card3ID, 4)
	drawBox(card3Layer.Canvas, 0, 0, 18, 6)
	ls.MoveLayer(card3ID, 42, 4)
	
	card4ID := ls.AddLayerWithName("評価カード")
	card4Layer, _ := ls.GetLayer(card4ID)
	ls.SetZOrder(card4ID, 5)
	drawBox(card4Layer.Canvas, 0, 0, 18, 6)
	ls.MoveLayer(card4ID, 62, 4)
	
	// グラフエリア
	chartID := ls.AddLayerWithName("グラフエリア")
	chartLayer, _ := ls.GetLayer(chartID)
	ls.SetZOrder(chartID, 6)
	drawBox(chartLayer.Canvas, 0, 0, 48, 12)
	ls.MoveLayer(chartID, 2, 12)
	
	// 最近のアクティビティ
	activityID := ls.AddLayerWithName("アクティビティ")
	activityLayer, _ := ls.GetLayer(activityID)
	ls.SetZOrder(activityID, 7)
	drawBox(activityLayer.Canvas, 0, 0, 26, 12)
	ls.MoveLayer(activityID, 52, 12)
	
	// フッター
	footerID := ls.AddLayerWithName("フッター")
	footerLayer, _ := ls.GetLayer(footerID)
	ls.SetZOrder(footerID, 8)
	drawBox(footerLayer.Canvas, 0, 0, 79, 2)
	ls.MoveLayer(footerID, 0, 26)
	
	// テキストレイヤー (最前面) - 全角対応版
	textID := ls.AddLayerWithName("テキストレイヤー")
	textLayer, _ := ls.GetLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// ナビゲーションテキスト（全角対応）
	textLayer.Canvas.WriteText(2, 1, "管理ダッシュボード")
	textLayer.Canvas.WriteTextRight(2, 1, 75, "ログアウト")
	
	// カードのタイトルと数値（全角対応）
	textLayer.Canvas.WriteText(4, 5, "売上")
	textLayer.Canvas.WriteText(4, 7, "¥1,250,000")
	textLayer.Canvas.WriteText(4, 8, "+12.5%")
	
	textLayer.Canvas.WriteText(24, 5, "ユーザー")
	textLayer.Canvas.WriteText(24, 7, "2,348")
	textLayer.Canvas.WriteText(24, 8, "+5.2%")
	
	textLayer.Canvas.WriteText(44, 5, "注文")
	textLayer.Canvas.WriteText(44, 7, "1,429")
	textLayer.Canvas.WriteText(44, 8, "+8.1%")
	
	textLayer.Canvas.WriteText(64, 5, "評価")
	textLayer.Canvas.WriteText(64, 7, "4.8/5.0")
	textLayer.Canvas.WriteText(64, 8, "+0.2")
	
	// グラフタイトル（全角対応）
	textLayer.Canvas.WriteText(4, 13, "月別売上推移")
	
	// 簡単なグラフ表現
	textLayer.Canvas.WriteText(4, 15, "1月 [=======---] 70%")
	textLayer.Canvas.WriteText(4, 16, "2月 [==========-] 85%")
	textLayer.Canvas.WriteText(4, 17, "3月 [===========] 95%")
	textLayer.Canvas.WriteText(4, 18, "4月 [============] 100%")
	textLayer.Canvas.WriteText(4, 19, "5月 [===========-] 90%")
	textLayer.Canvas.WriteText(4, 20, "6月 [============] 100%")
	
	// アクティビティ（全角対応）
	textLayer.Canvas.WriteText(54, 13, "最近のアクティビティ")
	textLayer.Canvas.WriteText(54, 15, "• 新規注文 #1234")
	textLayer.Canvas.WriteText(54, 16, "• ユーザー登録")
	textLayer.Canvas.WriteText(54, 17, "• 商品レビュー投稿")
	textLayer.Canvas.WriteText(54, 18, "• 支払い完了")
	textLayer.Canvas.WriteText(54, 19, "• 配送開始")
	textLayer.Canvas.WriteText(54, 20, "• 問い合わせ受信")
	
	// フッターテキスト
	textLayer.Canvas.WriteText(2, 27, "© 2024 MyCompany. All rights reserved.")
	
	fmt.Println("修正版: 全角対応eコマース管理ダッシュボード")
	fmt.Println("======================================")
	
	// 文字幅テスト表示
	fmt.Println("文字幅テスト:")
	testTexts := []string{"管理", "ダッシュボード", "Hello", "こんにちは"}
	for _, text := range testTexts {
		width := canvas.GetTextWidth(text)
		fmt.Printf("'%s' -> %d文字幅\n", text, width)
	}
	
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n修正された合成結果:")
	fmt.Println(result.String())
	
	file, err := os.Create("corrected_dashboard_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Corrected E-Commerce Dashboard UI (Full-width Japanese Support)\n")
	file.WriteString("===============================================================\n\n")
	file.WriteString("Character Width Test Results:\n")
	
	for _, text := range testTexts {
		width := canvas.GetTextWidth(text)
		file.WriteString(fmt.Sprintf("'%s' -> %d chars width\n", text, width))
	}
	
	file.WriteString("\nASCII Art Layout:\n")
	file.WriteString(result.String())
	
	fmt.Println("\ncorrected_dashboard_output.txtに出力しました")
}