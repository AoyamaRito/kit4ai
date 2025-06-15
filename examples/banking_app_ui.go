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

func drawShadowBox(c *canvas.Canvas, x1, y1, x2, y2 int) {
	// メインボックス
	drawBox(c, x1, y1, x2, y2)
	
	// 影効果
	for x := x1+1; x <= x2+1; x++ {
		if x < 80 && y2+1 < 100 {
			c.ReplaceChar(x, y2+1, '.')
		}
	}
	for y := y1+1; y <= y2+1; y++ {
		if x2+1 < 80 && y < 100 {
			c.ReplaceChar(x2+1, y, '.')
		}
	}
}

func drawProgressBar(c *canvas.Canvas, x, y, width, percent int) {
	filled := (width * percent) / 100
	for i := 0; i < width; i++ {
		if i < filled {
			c.ReplaceChar(x+i, y, '■')
		} else {
			c.ReplaceChar(x+i, y, '□')
		}
	}
}

func main() {
	fmt.Println("銀行アプリUI設計")
	fmt.Println("===============")
	
	ls := canvas.NewLayerSystem()
	
	// メインフレーム
	frameID := ls.AddLayerWithName("メインフレーム")
	frameLayer, _ := ls.GetLayer(frameID)
	ls.SetZOrder(frameID, 1)
	drawBox(frameLayer.Canvas, 0, 0, 79, 40)
	
	// ヘッダーエリア
	headerID := ls.AddLayerWithName("ヘッダー")
	headerLayer, _ := ls.GetLayer(headerID)
	ls.SetZOrder(headerID, 2)
	drawBox(headerLayer.Canvas, 1, 1, 78, 4)
	
	// 残高カード
	balanceID := ls.AddLayerWithName("残高カード")
	balanceLayer, _ := ls.GetLayer(balanceID)
	ls.SetZOrder(balanceID, 3)
	drawShadowBox(balanceLayer.Canvas, 3, 6, 38, 12)
	
	// クレジットカード
	creditID := ls.AddLayerWithName("クレジットカード")
	creditLayer, _ := ls.GetLayer(creditID)
	ls.SetZOrder(creditID, 4)
	drawShadowBox(creditLayer.Canvas, 42, 6, 77, 12)
	
	// 取引履歴エリア
	historyID := ls.AddLayerWithName("取引履歴")
	historyLayer, _ := ls.GetLayer(historyID)
	ls.SetZOrder(historyID, 5)
	drawBox(historyLayer.Canvas, 3, 15, 50, 30)
	
	// 投資ポートフォリオ
	portfolioID := ls.AddLayerWithName("投資ポートフォリオ")
	portfolioLayer, _ := ls.GetLayer(portfolioID)
	ls.SetZOrder(portfolioID, 6)
	drawBox(portfolioLayer.Canvas, 52, 15, 77, 24)
	
	// 資産推移グラフ
	chartID := ls.AddLayerWithName("資産推移")
	chartLayer, _ := ls.GetLayer(chartID)
	ls.SetZOrder(chartID, 7)
	drawBox(chartLayer.Canvas, 52, 26, 77, 35)
	
	// クイックアクション
	actionID := ls.AddLayerWithName("クイックアクション")
	actionLayer, _ := ls.GetLayer(actionID)
	ls.SetZOrder(actionID, 8)
	drawBox(actionLayer.Canvas, 3, 32, 50, 38)
	
	// フッター
	footerID := ls.AddLayerWithName("フッター")
	footerLayer, _ := ls.GetLayer(footerID)
	ls.SetZOrder(footerID, 9)
	drawBox(footerLayer.Canvas, 1, 39, 78, 39)
	
	// テキストレイヤー
	textID := ls.AddTextLayer("テキスト")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// ヘッダー
	textLayer.WriteTextWithWidth(3, 2, "🏦 みずほ銀行")
	textLayer.WriteTextRight(1, 2, 76, "田中太郎 様")
	textLayer.WriteTextCentered(1, 3, 77, "2024年06月15日 14:32")
	
	// 残高カード
	textLayer.WriteTextCentered(3, 7, 35, "普通預金残高")
	textLayer.WriteTextCentered(3, 9, 35, "¥1,247,850")
	textLayer.WriteTextWithWidth(5, 10, "前月比: +¥25,400 (+2.1%)")
	textLayer.WriteTextRight(3, 11, 35, "*** 口座番号: 1234567")
	
	// クレジットカード
	textLayer.WriteTextCentered(42, 7, 35, "クレジットカード")
	textLayer.WriteTextWithWidth(44, 8, "今月利用額: ¥89,250")
	textLayer.WriteTextWithWidth(44, 9, "利用可能額: ¥410,750")
	textLayer.WriteTextWithWidth(44, 10, "引き落とし日: 6/27")
	
	// 利用率プログレスバー
	drawProgressBar(creditLayer.Canvas, 44, 11, 30, 18)
	textLayer.WriteTextWithWidth(44, 11, "                              18%")
	
	// 取引履歴
	textLayer.WriteTextCentered(3, 16, 47, "最近の取引履歴")
	
	transactions := [][]string{
		{"06/15", "コンビニ決済", "-¥680"},
		{"06/14", "給与振込", "+¥285,000"},
		{"06/13", "電気料金", "-¥8,450"},
		{"06/12", "ATM出金", "-¥20,000"},
		{"06/11", "ネット通販", "-¥12,800"},
		{"06/10", "家賃振込", "-¥85,000"},
		{"06/09", "スーパー", "-¥3,250"},
		{"06/08", "カフェ", "-¥420"},
		{"06/07", "書籍購入", "-¥2,800"},
		{"06/06", "交通費", "-¥1,340"},
	}
	
	for i, tx := range transactions {
		if i < 10 {
			y := 18 + i
			textLayer.WriteTextWithWidth(5, y, tx[0])
			textLayer.WriteTextWithWidth(12, y, tx[1])
			textLayer.WriteTextRight(3, y, 47, tx[2])
		}
	}
	
	// 投資ポートフォリオ
	textLayer.WriteTextCentered(52, 16, 25, "投資ポートフォリオ")
	textLayer.WriteTextWithWidth(54, 18, "総評価額: ¥580,450")
	textLayer.WriteTextWithWidth(54, 19, "評価損益: +¥38,200")
	textLayer.WriteTextWithWidth(54, 20, "           (+7.0%)")
	
	stocks := [][]string{
		{"日本株", "65%"},
		{"米国株", "25%"},
		{"債券", "10%"},
	}
	
	for i, stock := range stocks {
		y := 22 + i
		textLayer.WriteTextWithWidth(54, y, stock[0])
		textLayer.WriteTextRight(52, y, 25, stock[1])
	}
	
	// 資産推移グラフ
	textLayer.WriteTextCentered(52, 27, 25, "6ヶ月資産推移")
	
	// 簡易グラフ
	months := []string{"1月", "2月", "3月", "4月", "5月", "6月"}
	heights := []int{3, 4, 2, 5, 4, 6}
	
	for i, month := range months {
		x := 54 + i*3
		textLayer.WriteTextWithWidth(x, 34, month)
		
		// 棒グラフ
		for h := 0; h < heights[i]; h++ {
			chartLayer.Canvas.ReplaceChar(x+1, 33-h, '|')
		}
	}
	
	// クイックアクション
	textLayer.WriteTextCentered(3, 33, 47, "クイックアクション")
	
	actions := []string{
		"💰 振込・送金", "📊 投資注文", "🏧 ATM検索", 
		"📱 家計簿", "💳 カード管理", "📞 サポート",
	}
	
	for i, action := range actions {
		x := 5 + (i%3)*15
		y := 35 + (i/3)
		textLayer.WriteTextWithWidth(x, y, action)
	}
	
	// フッター
	textLayer.WriteTextCentered(1, 39, 77, "🔒 安全な接続でご利用いただいています | ログアウト")
	
	fmt.Println("レイヤー構成:")
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n銀行アプリUI:")
	fmt.Println(result.String())
	
	// 仕様書出力
	file, err := os.Create("banking_app_ui_spec.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Banking App UI Design Specification\n")
	file.WriteString("====================================\n\n")
	file.WriteString("Application: みずほ銀行モバイルアプリ\n")
	file.WriteString("Screen Size: 80x41 characters\n\n")
	file.WriteString("Features:\n")
	file.WriteString("1. Header - Bank logo, user name, timestamp\n")
	file.WriteString("2. Account Balance Card - Current balance with trend\n")
	file.WriteString("3. Credit Card Info - Usage and available limit\n")
	file.WriteString("4. Transaction History - Recent 10 transactions\n")
	file.WriteString("5. Investment Portfolio - Stock allocation and performance\n")
	file.WriteString("6. Asset Trend Chart - 6-month asset growth visualization\n")
	file.WriteString("7. Quick Actions - Common banking operations\n")
	file.WriteString("8. Security Footer - SSL connection status\n\n")
	file.WriteString("Design Elements:\n")
	file.WriteString("- Shadow effects for card-like appearance\n")
	file.WriteString("- Progress bars for credit utilization\n")
	file.WriteString("- Simple bar chart for asset trends\n")
	file.WriteString("- Emoji icons for visual appeal\n")
	file.WriteString("- Right-aligned amounts for easy reading\n\n")
	file.WriteString("Security Features:\n")
	file.WriteString("- Masked account numbers (*** prefix)\n")
	file.WriteString("- Secure connection indicator\n")
	file.WriteString("- User identification in header\n\n")
	file.WriteString("Layer Structure:\n")
	
	for _, id := range ls.GetLayerIDs() {
		layer, _ := ls.GetLayer(id)
		file.WriteString(fmt.Sprintf("- %s (Z-Order: %d)\n", layer.Name, layer.ZOrder))
	}
	
	file.WriteString("\nASCII Art Layout:\n")
	file.WriteString(result.String())
	
	fmt.Println("\nbanking_app_ui_spec.txtに出力しました")
	fmt.Println("\n特徴:")
	fmt.Println("- 本格的な銀行アプリUI")
	fmt.Println("- 残高・クレジット・投資の一元管理")
	fmt.Println("- 取引履歴とリアルタイム資産推移")
	fmt.Println("- セキュリティ機能の表示")
	fmt.Println("- 直感的なクイックアクション")
}