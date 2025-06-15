package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
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
	fmt.Println("修正版レイヤー合成テスト")
	fmt.Println("=======================")
	
	ls := canvas.NewLayerSystem()
	
	// レイヤー1: 背景ボックス（Z-Order: 1）
	box1ID := ls.AddLayerWithName("背景ボックス")
	box1Layer, _ := ls.GetLayer(box1ID)
	ls.SetZOrder(box1ID, 1)
	drawBox(box1Layer.Canvas, 5, 5, 25, 15)
	
	// レイヤー2: 重複ボックス（Z-Order: 2）
	box2ID := ls.AddLayerWithName("重複ボックス")
	box2Layer, _ := ls.GetLayer(box2ID)
	ls.SetZOrder(box2ID, 2)
	drawBox(box2Layer.Canvas, 15, 10, 35, 20)
	
	// レイヤー3: テキストレイヤー（Z-Order: 100 - 最上位）
	textID := ls.AddTextLayer("テキスト")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	// テキスト配置
	textLayer.WriteTextWithWidth(7, 7, "背景ボックス")
	textLayer.WriteTextWithWidth(17, 12, "重複エリア")
	textLayer.WriteTextWithWidth(25, 17, "前面ボックス")
	
	// 中央寄せテスト
	textLayer.WriteTextCentered(5, 10, 20, "中央寄せ")
	
	fmt.Println("レイヤー構成:")
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n合成結果:")
	fmt.Println(result.String())
	
	fmt.Println("\n合成ルール確認:")
	fmt.Println("1. 80文字x100行の配列を作成")
	fmt.Println("2. Z-Order順（小→大）で下層から上層へ合成")
	fmt.Println("3. スペース位置 -> 新しい文字を配置")
	fmt.Println("4. 文字がある位置 -> 上位レイヤー優先で上書き")
	fmt.Println("5. 新しい文字がスペース -> 既存文字を保持")
	
	// 各レイヤー単体表示
	fmt.Println("\n=== 各レイヤー単体 ===")
	
	fmt.Println("\n背景ボックス（Z-Order: 1）:")
	fmt.Println(box1Layer.Canvas.String())
	
	fmt.Println("\n重複ボックス（Z-Order: 2）:")
	fmt.Println(box2Layer.Canvas.String())
	
	fmt.Println("\nテキストレイヤー（Z-Order: 100）:")
	fmt.Println(textLayer.Canvas.String())
}