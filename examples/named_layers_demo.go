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
	
	headerID := ls.AddLayerWithName("ヘッダー")
	headerLayer, _ := ls.GetLayer(headerID)
	ls.SetZOrder(headerID, 1)
	drawBox(headerLayer.Canvas, 0, 0, 30, 3)
	
	sidebarID := ls.AddLayerWithName("サイドバー")
	sidebarLayer, _ := ls.GetLayer(sidebarID)
	ls.SetZOrder(sidebarID, 2)
	drawBox(sidebarLayer.Canvas, 0, 0, 15, 10)
	ls.MoveLayer(sidebarID, 0, 4)
	
	mainID := ls.AddLayerWithName("メインコンテンツ")
	mainLayer, _ := ls.GetLayer(mainID)
	ls.SetZOrder(mainID, 3)
	drawBox(mainLayer.Canvas, 0, 0, 25, 8)
	ls.MoveLayer(mainID, 17, 6)
	
	modalID := ls.AddLayerWithName("モーダル")
	modalLayer, _ := ls.GetLayer(modalID)
	ls.SetZOrder(modalID, 10)
	drawBox(modalLayer.Canvas, 0, 0, 20, 5)
	ls.MoveLayer(modalID, 10, 8)
	
	fmt.Println("レイヤー名とオフセット機能のデモ")
	fmt.Println("===============================")
	
	ls.ListLayers()
	
	result := ls.Composite()
	
	fmt.Println("\n合成結果:")
	fmt.Println(result.String())
	
	file, err := os.Create("named_layers_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	file.WriteString("Named Layers Demo Output\n")
	file.WriteString("========================\n\n")
	file.WriteString("Layer Information:\n")
	
	for _, id := range ls.GetLayerIDs() {
		name, zorder, offsetX, offsetY, _ := ls.GetLayerInfo(id)
		file.WriteString(fmt.Sprintf("ID: %d, Name: %s, Z-Order: %d, Offset: (%d,%d)\n", 
			id, name, zorder, offsetX, offsetY))
	}
	
	file.WriteString("\nComposite Result:\n")
	file.WriteString(result.String())
	
	fmt.Println("\nnamed_layers_output.txtに出力しました")
}