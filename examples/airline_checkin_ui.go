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

func drawRoundedBox(c *canvas.Canvas, x1, y1, x2, y2 int) {
	for x := x1+1; x < x2; x++ {
		c.ReplaceChar(x, y1, '-')
		c.ReplaceChar(x, y2, '-')
	}
	
	for y := y1+1; y < y2; y++ {
		c.ReplaceChar(x1, y, '|')
		c.ReplaceChar(x2, y, '|')
	}
	
	c.ReplaceChar(x1, y1, '.')
	c.ReplaceChar(x2, y1, '.')
	c.ReplaceChar(x1, y2, '`')
	c.ReplaceChar(x2, y2, '\'')
}

func drawPlane(c *canvas.Canvas, x, y int) {
	// 簡易飛行機アイコン
	c.ReplaceChar(x, y, '-')
	c.ReplaceChar(x+1, y, '=')
	c.ReplaceChar(x+2, y, '✈')
	c.ReplaceChar(x+3, y, '=')
	c.ReplaceChar(x+4, y, '-')
}

func drawSeatMap(c *canvas.Canvas, x1, y1 int) {
	// 座席マップ表現
	seats := []string{
		"A B C   D E F",
		"[■][■][■] [□][□][□] 1",
		"[□][■][■] [□][□][■] 2", 
		"[■][□][□] [■][■][□] 3",
		"[□][□][■] [□][■][■] 4",
		"[■][■][□] [□][□][□] 5",
	}
	
	for i, row := range seats {
		for j, char := range []rune(row) {
			c.ReplaceChar(x1+j, y1+i, char)
		}
	}
}

func main() {
	fmt.Println("航空会社チェックインキオスクUI設計")
	fmt.Println("================================")
	
	ls := canvas.NewLayerSystem()
	
	// メインフレーム
	frameID := ls.AddLayerWithName("メインフレーム")
	frameLayer, _ := ls.GetLayer(frameID)
	ls.SetZOrder(frameID, 1)
	drawBox(frameLayer.Canvas, 0, 0, 79, 45)
	
	// ヘッダーエリア
	headerID := ls.AddLayerWithName("ヘッダー")
	headerLayer, _ := ls.GetLayer(headerID)
	ls.SetZOrder(headerID, 2)
	drawBox(headerLayer.Canvas, 1, 1, 78, 5)
	
	// フライト情報カード
	flightID := ls.AddLayerWithName("フライト情報")
	flightLayer, _ := ls.GetLayer(flightID)
	ls.SetZOrder(flightID, 3)
	drawRoundedBox(flightLayer.Canvas, 3, 7, 38, 18)
	
	// 搭乗券プレビュー
	boardingID := ls.AddLayerWithName("搭乗券")
	boardingLayer, _ := ls.GetLayer(boardingID)
	ls.SetZOrder(boardingID, 4)
	drawBox(boardingLayer.Canvas, 42, 7, 77, 23)
	
	// 座席選択エリア
	seatID := ls.AddLayerWithName("座席選択")
	seatLayer, _ := ls.GetLayer(seatID)
	ls.SetZOrder(seatID, 5)
	drawBox(seatLayer.Canvas, 3, 20, 38, 35)
	
	// ステータスバー
	statusID := ls.AddLayerWithName("ステータス")
	statusLayer, _ := ls.GetLayer(statusID)
	ls.SetZOrder(statusID, 6)
	drawBox(statusLayer.Canvas, 3, 37, 77, 41)
	
	// アクションボタンエリア
	actionID := ls.AddLayerWithName("アクション")
	actionLayer, _ := ls.GetLayer(actionID)
	ls.SetZOrder(actionID, 7)
	drawBox(actionLayer.Canvas, 42, 25, 77, 35)
	
	// フッター
	footerID := ls.AddLayerWithName("フッター")
	footerLayer, _ := ls.GetLayer(footerID)
	ls.SetZOrder(footerID, 8)
	drawBox(footerLayer.Canvas, 1, 43, 78, 44)
	
	// テキストレイヤー
	textID := ls.AddTextLayer("テキスト")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// ヘッダー
	textLayer.WriteTextWithWidth(3, 2, "✈ スカイライン航空")
	textLayer.WriteTextRight(1, 2, 76, "セルフチェックイン")
	textLayer.WriteTextCentered(1, 3, 77, "出発まで残り 2時間 15分")
	textLayer.WriteTextCentered(1, 4, 77, "2024年06月15日 16:45")
	
	// フライト情報
	textLayer.WriteTextCentered(3, 8, 35, "フライト情報")
	textLayer.WriteTextWithWidth(5, 10, "便名: SL1234")
	textLayer.WriteTextWithWidth(5, 11, "出発: 東京羽田 (HND)")
	textLayer.WriteTextWithWidth(5, 12, "到着: 大阪関西 (KIX)")
	textLayer.WriteTextWithWidth(5, 13, "出発時刻: 19:00")
	textLayer.WriteTextWithWidth(5, 14, "到着時刻: 20:15")
	textLayer.WriteTextWithWidth(5, 15, "搭乗ゲート: A12")
	textLayer.WriteTextWithWidth(5, 16, "機材: Boeing 737-800")
	
	// 飛行機アイコン
	drawPlane(flightLayer.Canvas, 25, 10)
	
	// 搭乗券プレビュー
	textLayer.WriteTextCentered(42, 8, 35, "搭乗券プレビュー")
	textLayer.WriteTextWithWidth(44, 10, "┌──────────────────────────────┐")
	textLayer.WriteTextWithWidth(44, 11, "│ スカイライン航空 SL1234         │")
	textLayer.WriteTextWithWidth(44, 12, "│                              │")
	textLayer.WriteTextWithWidth(44, 13, "│ 田中太郎 様                    │")
	textLayer.WriteTextWithWidth(44, 14, "│ TANAKA TARO                  │")
	textLayer.WriteTextWithWidth(44, 15, "│                              │")
	textLayer.WriteTextWithWidth(44, 16, "│ HND → KIX   19:00 → 20:15    │")
	textLayer.WriteTextWithWidth(44, 17, "│ ゲート: A12  座席: 12A        │")
	textLayer.WriteTextWithWidth(44, 18, "│ エコノミークラス              │")
	textLayer.WriteTextWithWidth(44, 19, "│                              │")
	textLayer.WriteTextWithWidth(44, 20, "│ 搭乗開始: 18:30              │")
	textLayer.WriteTextWithWidth(44, 21, "│ QR: ████████                │")
	textLayer.WriteTextWithWidth(44, 22, "└──────────────────────────────┘")
	
	// 座席選択
	textLayer.WriteTextCentered(3, 21, 35, "座席選択")
	textLayer.WriteTextWithWidth(5, 23, "■ 選択済み  □ 空席  ❌ 不可")
	
	// 座席マップを描画
	drawSeatMap(seatLayer.Canvas, 8, 25)
	
	textLayer.WriteTextWithWidth(5, 32, "現在選択: 12A (窓側)")
	textLayer.WriteTextWithWidth(5, 33, "追加料金: ¥0")
	
	// アクションボタン
	textLayer.WriteTextCentered(42, 26, 35, "操作メニュー")
	
	buttons := []string{
		"✓ チェックイン完了",
		"💺 座席を変更", 
		"🍽️ 機内食選択",
		"👥 同行者追加",
		"📧 搭乗券をメール送信",
		"🖨️ 搭乗券を印刷",
	}
	
	for i, button := range buttons {
		y := 28 + i
		textLayer.WriteTextCentered(42, y, 35, button)
	}
	
	// ステータスバー
	textLayer.WriteTextWithWidth(5, 38, "チェックイン状況:")
	textLayer.WriteTextWithWidth(5, 39, "✓ 本人確認完了  ✓ 座席選択完了  ⏳ 手荷物確認")
	
	// プログレスバー
	for i := 0; i < 60; i++ {
		char := '='
		if i > 40 {
			char = '-'
		}
		statusLayer.Canvas.ReplaceChar(15 + i, 40, char)
	}
	textLayer.WriteTextWithWidth(5, 40, "進捗: [========================================------------] 67%")
	
	// フッター
	textLayer.WriteTextCentered(1, 43, 77, "💳 支払い不要 | ❓ ヘルプ | 🔄 言語変更 | 🚪 キャンセル")
	
	fmt.Println("レイヤー構成:")
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n航空会社チェックインキオスクUI:")
	fmt.Println(result.String())
	
	// 仕様書出力
	file, err := os.Create("airline_checkin_ui_spec.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Airline Check-in Kiosk UI Design Specification\n")
	file.WriteString("===============================================\n\n")
	file.WriteString("Application: スカイライン航空 セルフチェックインキオスク\n")
	file.WriteString("Screen Size: 80x46 characters\n")
	file.WriteString("Target Users: 航空会社の乗客（セルフサービス）\n\n")
	file.WriteString("Features:\n")
	file.WriteString("1. Header - Airline branding, countdown timer, current time\n")
	file.WriteString("2. Flight Information - Flight details with aircraft icon\n")
	file.WriteString("3. Boarding Pass Preview - Real-time boarding pass display\n")
	file.WriteString("4. Seat Selection - Interactive seat map with availability\n")
	file.WriteString("5. Action Menu - Check-in completion and additional services\n")
	file.WriteString("6. Progress Status - Step-by-step completion tracking\n")
	file.WriteString("7. Footer - Help, language options, cancel\n\n")
	file.WriteString("UI Components:\n")
	file.WriteString("- Rounded boxes for modern card-like appearance\n")
	file.WriteString("- Seat map with visual seat availability (■□❌)\n")
	file.WriteString("- Progress bar for completion tracking\n")
	file.WriteString("- QR code placeholder for boarding pass\n")
	file.WriteString("- Emoji icons for intuitive navigation\n\n")
	file.WriteString("Accessibility Features:\n")
	file.WriteString("- Large, clear text for easy reading\n")
	file.WriteString("- High contrast visual elements\n")
	file.WriteString("- Multilingual support option\n")
	file.WriteString("- Touch-friendly button spacing\n\n")
	file.WriteString("Business Logic:\n")
	file.WriteString("- Real-time flight information display\n")
	file.WriteString("- Seat availability checking\n")
	file.WriteString("- Boarding pass generation\n")
	file.WriteString("- Email/print options for boarding pass\n")
	file.WriteString("- Additional service selection (meals, baggage)\n\n")
	file.WriteString("Layer Structure:\n")
	
	for _, id := range ls.GetLayerIDs() {
		layer, _ := ls.GetLayer(id)
		file.WriteString(fmt.Sprintf("- %s (Z-Order: %d)\n", layer.Name, layer.ZOrder))
	}
	
	file.WriteString("\nASCII Art Layout:\n")
	file.WriteString(result.String())
	
	fmt.Println("\nairline_checkin_ui_spec.txtに出力しました")
	fmt.Println("\n特徴:")
	fmt.Println("- 実用的な航空会社チェックインキオスク")
	fmt.Println("- 座席選択とリアルタイム搭乗券プレビュー")
	fmt.Println("- 直感的なタッチスクリーンUI")
	fmt.Println("- 多言語対応とアクセシビリティ考慮")
	fmt.Println("- プログレス表示で進捗を明確化")
}