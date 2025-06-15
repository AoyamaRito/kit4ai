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

func drawProgressBar(c *canvas.Canvas, x, y, width int, percent int) {
	filled := (width * percent) / 100
	
	c.ReplaceChar(x, y, '[')
	c.ReplaceChar(x+width+1, y, ']')
	
	for i := 1; i <= width; i++ {
		if i <= filled {
			c.ReplaceChar(x+i, y, '=')
		} else {
			c.ReplaceChar(x+i, y, '-')
		}
	}
}

func writeText(c *canvas.Canvas, x, y int, text string) {
	textRunes := []rune(text)
	for i, r := range textRunes {
		c.ReplaceChar(x+i, y, r)
	}
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
	
	// テキストレイヤー (最前面)
	textID := ls.AddLayerWithName("テキストレイヤー")
	textLayer, _ := ls.GetLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// ナビゲーションテキスト
	writeText(textLayer.Canvas, 2, 1, "管理ダッシュボード                                      ログアウト")
	
	// カードのタイトルと数値
	writeText(textLayer.Canvas, 4, 5, "売上")
	writeText(textLayer.Canvas, 4, 7, "¥1,250,000")
	writeText(textLayer.Canvas, 4, 8, "+12.5%")
	
	writeText(textLayer.Canvas, 24, 5, "ユーザー")
	writeText(textLayer.Canvas, 24, 7, "2,348")
	writeText(textLayer.Canvas, 24, 8, "+5.2%")
	
	writeText(textLayer.Canvas, 44, 5, "注文")
	writeText(textLayer.Canvas, 44, 7, "1,429")
	writeText(textLayer.Canvas, 44, 8, "+8.1%")
	
	writeText(textLayer.Canvas, 64, 5, "評価")
	writeText(textLayer.Canvas, 64, 7, "4.8/5.0")
	writeText(textLayer.Canvas, 64, 8, "+0.2")
	
	// グラフタイトル
	writeText(textLayer.Canvas, 4, 13, "月別売上推移")
	
	// 簡単なグラフ表現
	writeText(textLayer.Canvas, 4, 15, "1月 [=======---] 70%")
	writeText(textLayer.Canvas, 4, 16, "2月 [==========-] 85%")
	writeText(textLayer.Canvas, 4, 17, "3月 [===========] 95%")
	writeText(textLayer.Canvas, 4, 18, "4月 [============] 100%")
	writeText(textLayer.Canvas, 4, 19, "5月 [===========-] 90%")
	writeText(textLayer.Canvas, 4, 20, "6月 [============] 100%")
	
	// アクティビティ
	writeText(textLayer.Canvas, 54, 13, "最近のアクティビティ")
	writeText(textLayer.Canvas, 54, 15, "• 新規注文 #1234")
	writeText(textLayer.Canvas, 54, 16, "• ユーザー登録")
	writeText(textLayer.Canvas, 54, 17, "• 商品レビュー投稿")
	writeText(textLayer.Canvas, 54, 18, "• 支払い完了")
	writeText(textLayer.Canvas, 54, 19, "• 配送開始")
	writeText(textLayer.Canvas, 54, 20, "• 問い合わせ受信")
	
	// フッターテキスト
	writeText(textLayer.Canvas, 2, 27, "© 2024 MyCompany. All rights reserved.")
	
	fmt.Println("eコマース管理ダッシュボード UI")
	fmt.Println("===============================")
	
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n完成したダッシュボード:")
	fmt.Println(result.String())
	
	file, err := os.Create("dashboard_ui_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("E-Commerce Dashboard UI Specification\n")
	file.WriteString("=====================================\n\n")
	file.WriteString("Components:\n")
	file.WriteString("- Top Navigation Bar\n")
	file.WriteString("- 4 Statistics Cards (Sales, Users, Orders, Rating)\n")
	file.WriteString("- Monthly Sales Chart Area\n")
	file.WriteString("- Recent Activity Panel\n")
	file.WriteString("- Footer\n\n")
	file.WriteString("ASCII Art Layout:\n")
	file.WriteString(result.String())
	
	fmt.Println("\ndashboard_ui_output.txtに出力しました")
}