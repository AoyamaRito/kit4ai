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

func drawDoubleBox(c *canvas.Canvas, x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		c.ReplaceChar(x, y1, '=')
		c.ReplaceChar(x, y2, '=')
	}
	
	for y := y1; y <= y2; y++ {
		c.ReplaceChar(x1, y, '#')
		c.ReplaceChar(x2, y, '#')
	}
	
	c.ReplaceChar(x1, y1, '#')
	c.ReplaceChar(x2, y1, '#')
	c.ReplaceChar(x1, y2, '#')
	c.ReplaceChar(x2, y2, '#')
}

func drawHealthBar(c *canvas.Canvas, x, y, width, percent int) {
	c.ReplaceChar(x, y, '[')
	c.ReplaceChar(x+width+1, y, ']')
	
	filled := (width * percent) / 100
	for i := 1; i <= width; i++ {
		if i <= filled {
			c.ReplaceChar(x+i, y, '█')
		} else {
			c.ReplaceChar(x+i, y, '░')
		}
	}
}

func drawMiniMap(c *canvas.Canvas, x1, y1, x2, y2 int) {
	// ミニマップの枠
	drawBox(c, x1, y1, x2, y2)
	
	// 地形表現
	for y := y1+1; y < y2; y++ {
		for x := x1+1; x < x2; x++ {
			if (x+y)%3 == 0 {
				c.ReplaceChar(x, y, '.')
			} else if (x+y)%5 == 0 {
				c.ReplaceChar(x, y, '^')
			} else {
				c.ReplaceChar(x, y, ' ')
			}
		}
	}
	
	// プレイヤー位置
	c.ReplaceChar(x1+3, y1+2, '@')
}

func main() {
	fmt.Println("RPGゲームUI設計")
	fmt.Println("===============")
	
	ls := canvas.NewLayerSystem()
	
	// メインウィンドウ
	mainID := ls.AddLayerWithName("メインウィンドウ")
	mainLayer, _ := ls.GetLayer(mainID)
	ls.SetZOrder(mainID, 1)
	drawDoubleBox(mainLayer.Canvas, 0, 0, 79, 35)
	
	// ゲーム画面エリア
	gameAreaID := ls.AddLayerWithName("ゲーム画面")
	gameAreaLayer, _ := ls.GetLayer(gameAreaID)
	ls.SetZOrder(gameAreaID, 2)
	drawBox(gameAreaLayer.Canvas, 2, 2, 55, 25)
	
	// ステータスパネル
	statusID := ls.AddLayerWithName("ステータスパネル")
	statusLayer, _ := ls.GetLayer(statusID)
	ls.SetZOrder(statusID, 3)
	drawBox(statusLayer.Canvas, 57, 2, 77, 15)
	
	// インベントリ
	inventoryID := ls.AddLayerWithName("インベントリ")
	inventoryLayer, _ := ls.GetLayer(inventoryID)
	ls.SetZOrder(inventoryID, 4)
	drawBox(inventoryLayer.Canvas, 57, 17, 77, 25)
	
	// ミニマップ
	minimapID := ls.AddLayerWithName("ミニマップ")
	minimapLayer, _ := ls.GetLayer(minimapID)
	ls.SetZOrder(minimapID, 5)
	drawMiniMap(minimapLayer.Canvas, 57, 27, 77, 33)
	
	// チャットログ
	chatID := ls.AddLayerWithName("チャットログ")
	chatLayer, _ := ls.GetLayer(chatID)
	ls.SetZOrder(chatID, 6)
	drawBox(chatLayer.Canvas, 2, 27, 55, 33)
	
	// ゲーム内容描画（簡易ダンジョン）
	dungeonID := ls.AddLayerWithName("ダンジョン")
	dungeonLayer, _ := ls.GetLayer(dungeonID)
	ls.SetZOrder(dungeonID, 10)
	
	// ダンジョンの壁
	for x := 3; x < 55; x++ {
		for y := 3; y < 25; y++ {
			if x == 3 || x == 54 || y == 3 || y == 24 {
				dungeonLayer.Canvas.ReplaceChar(x, y, '#')
			} else if (x%5 == 0 && y%3 == 0) {
				dungeonLayer.Canvas.ReplaceChar(x, y, 'o') // 柱
			} else {
				dungeonLayer.Canvas.ReplaceChar(x, y, '.')
			}
		}
	}
	
	// プレイヤーキャラクター
	dungeonLayer.Canvas.ReplaceChar(10, 10, '@')
	
	// モンスター
	dungeonLayer.Canvas.ReplaceChar(20, 8, 'D') // ドラゴン
	dungeonLayer.Canvas.ReplaceChar(30, 15, 'G') // ゴブリン
	dungeonLayer.Canvas.ReplaceChar(45, 12, 'S') // スライム
	
	// 宝箱
	dungeonLayer.Canvas.ReplaceChar(50, 20, '$')
	
	// テキストレイヤー（全角対応）
	textID := ls.AddTextLayer("テキスト")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// メインタイトル
	textLayer.WriteTextCentered(0, 1, 79, "魔王の城 - 地下5階")
	
	// ステータス情報
	textLayer.WriteTextSimple(59, 3, "プレイヤー情報")
	textLayer.WriteTextSimple(59, 4, "===============")
	textLayer.WriteTextSimple(59, 5, "勇者アレックス")
	textLayer.WriteTextSimple(59, 6, "Lv.25  経験値:12450")
	
	// HPバー描画
	drawHealthBar(statusLayer.Canvas, 58, 7, 15, 85)
	textLayer.WriteTextSimple(59, 8, "HP: 340/400")
	
	// MPバー描画
	drawHealthBar(statusLayer.Canvas, 58, 9, 15, 60)
	textLayer.WriteTextSimple(59, 10, "MP: 120/200")
	
	textLayer.WriteTextSimple(59, 12, "攻撃力: 85")
	textLayer.WriteTextSimple(59, 13, "防御力: 62")
	textLayer.WriteTextSimple(59, 14, "素早さ: 73")
	
	// インベントリ
	textLayer.WriteTextSimple(59, 18, "アイテム")
	textLayer.WriteTextSimple(59, 19, "========")
	textLayer.WriteTextSimple(59, 20, "鋼の剣     x1")
	textLayer.WriteTextSimple(59, 21, "回復薬     x5")
	textLayer.WriteTextSimple(59, 22, "魔法の石   x2")
	textLayer.WriteTextSimple(59, 23, "金貨      x250")
	
	// ミニマップラベル
	textLayer.WriteTextCentered(57, 26, 20, "ミニマップ")
	
	// チャットログ
	textLayer.WriteTextSimple(4, 28, "システムメッセージ")
	textLayer.WriteTextSimple(4, 29, "> ドラゴンが現れた！")
	textLayer.WriteTextSimple(4, 30, "> 勇者の攻撃！ 45ダメージ")
	textLayer.WriteTextSimple(4, 31, "> ドラゴンの炎攻撃！")
	textLayer.WriteTextSimple(4, 32, "> 25ダメージを受けた")
	
	// 操作説明
	textLayer.WriteTextCentered(0, 34, 79, "WASD:移動 SPACE:攻撃 I:インベントリ ESC:メニュー")
	
	fmt.Println("レイヤー構成:")
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\nRPGゲームUI:")
	fmt.Println(result.String())
	
	// ゲーム仕様書として出力
	file, err := os.Create("rpg_game_ui_spec.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("RPG Game UI Design Specification\n")
	file.WriteString("=================================\n\n")
	file.WriteString("Game: 魔王の城 (Demon King's Castle)\n")
	file.WriteString("Screen Resolution: 80x36 characters\n\n")
	file.WriteString("UI Components:\n")
	file.WriteString("1. Main Game Area (54x24) - Dungeon view with player, monsters, items\n")
	file.WriteString("2. Status Panel (20x14) - Player stats, HP/MP bars, level info\n")
	file.WriteString("3. Inventory Panel (20x9) - Item list with quantities\n")
	file.WriteString("4. Mini Map (20x7) - Terrain overview with player position\n")
	file.WriteString("5. Chat Log (54x7) - System messages and battle log\n")
	file.WriteString("6. Control Instructions - Key bindings display\n\n")
	file.WriteString("Game Elements:\n")
	file.WriteString("- Player: @ symbol\n")
	file.WriteString("- Dragon: D symbol\n")
	file.WriteString("- Goblin: G symbol\n")
	file.WriteString("- Slime: S symbol\n")
	file.WriteString("- Treasure: $ symbol\n")
	file.WriteString("- Walls: # symbol\n")
	file.WriteString("- Floor: . symbol\n")
	file.WriteString("- Pillars: o symbol\n\n")
	file.WriteString("Health/Mana Bars:\n")
	file.WriteString("- Filled: █ character\n")
	file.WriteString("- Empty: ░ character\n")
	file.WriteString("- Brackets: [ ] characters\n\n")
	file.WriteString("Layer Structure:\n")
	
	for _, id := range ls.GetLayerIDs() {
		layer, _ := ls.GetLayer(id)
		file.WriteString(fmt.Sprintf("- %s (Z-Order: %d)\n", layer.Name, layer.ZOrder))
	}
	
	file.WriteString("\nASCII Art Layout:\n")
	file.WriteString(result.String())
	
	fmt.Println("\nrpg_game_ui_spec.txtに出力しました")
	fmt.Println("\n特徴:")
	fmt.Println("- 本格的なRPGゲームUI")
	fmt.Println("- リアルタイム戦闘ログ")
	fmt.Println("- ステータス表示（HP/MPバー）")
	fmt.Println("- ミニマップ機能")
	fmt.Println("- インベントリ管理")
	fmt.Println("- ダンジョン探索画面")
}