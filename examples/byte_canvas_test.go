package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
)

func main() {
	fmt.Println("8bit単位処理テスト")
	fmt.Println("==================")
	
	// ByteCanvasを使用
	bc := canvas.NewByteCanvas()
	
	// 基本ボックス描画
	bc.DrawBox(5, 3, 25, 10)
	bc.DrawBox(30, 5, 50, 12)
	
	// テキスト配置（8bit単位）
	bc.WriteBytes(7, 5, "ASCII Box")
	bc.WriteBytes(32, 7, "Byte Canvas")
	
	// 日本語文字列（UTF-8バイト列として）
	bc.WriteBytes(7, 7, "Hello")
	bc.WriteBytes(7, 8, "World")
	
	// 罫線文字のテスト（UTF-8として）
	bc.WriteBytes(10, 15, "+----------+")
	bc.WriteBytes(10, 16, "| Safe Box |")
	bc.WriteBytes(10, 17, "+----------+")
	
	// 位置マーカー
	for i := 0; i < 60; i += 10 {
		marker := fmt.Sprintf("%d", i/10)
		bc.WriteBytes(i, 1, marker)
	}
	
	fmt.Println("ByteCanvas結果:")
	fmt.Println(bc.String())
	
	// runeベースとの比較
	fmt.Println("\n=== 比較テスト ===")
	
	// 従来のrune版
	ls := canvas.NewLayerSystem()
	
	boxID := ls.AddLayerWithName("ボックス")
	boxLayer, _ := ls.GetLayer(boxID)
	ls.SetZOrder(boxID, 1)
	
	// 手動ボックス描画
	for x := 5; x <= 25; x++ {
		boxLayer.Canvas.ReplaceChar(x, 3, '-')
		boxLayer.Canvas.ReplaceChar(x, 10, '-')
	}
	for y := 3; y <= 10; y++ {
		boxLayer.Canvas.ReplaceChar(5, y, '|')
		boxLayer.Canvas.ReplaceChar(25, y, '|')
	}
	boxLayer.Canvas.ReplaceChar(5, 3, '+')
	boxLayer.Canvas.ReplaceChar(25, 3, '+')
	boxLayer.Canvas.ReplaceChar(5, 10, '+')
	boxLayer.Canvas.ReplaceChar(25, 10, '+')
	
	textID := ls.AddTextLayer("テキスト")
	textLayer, _ := ls.GetTextLayer(textID)
	ls.SetZOrder(textID, 100)
	
	textLayer.WriteTextWithWidth(7, 5, "Rune Version")
	
	result := ls.Composite()
	
	fmt.Println("Rune版結果:")
	fmt.Println(result.String())
	
	fmt.Println("\n比較結果:")
	fmt.Println("- ByteCanvas: 8bit単位で確実な配置")
	fmt.Println("- RuneCanvas: Unicode対応だが位置ずれリスク")
	fmt.Println("- 用途に応じて使い分けが重要")
}