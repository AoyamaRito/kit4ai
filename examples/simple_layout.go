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
	
	// ヘッダー
	headerLayerID := ls.AddLayer()
	headerLayer, _ := ls.GetLayer(headerLayerID)
	ls.SetZOrder(headerLayerID, 1)
	drawBox(headerLayer.Canvas, 0, 0, 79, 4)
	
	// サイドバー
	sidebarLayerID := ls.AddLayer()
	sidebarLayer, _ := ls.GetLayer(sidebarLayerID)
	ls.SetZOrder(sidebarLayerID, 2)
	drawBox(sidebarLayer.Canvas, 0, 5, 19, 25)
	
	// メインエリア
	mainLayerID := ls.AddLayer()
	mainLayer, _ := ls.GetLayer(mainLayerID)
	ls.SetZOrder(mainLayerID, 3)
	drawBox(mainLayer.Canvas, 20, 5, 79, 25)
	
	// フッター
	footerLayerID := ls.AddLayer()
	footerLayer, _ := ls.GetLayer(footerLayerID)
	ls.SetZOrder(footerLayerID, 4)
	drawBox(footerLayer.Canvas, 0, 26, 79, 30)
	
	result := ls.Composite()
	
	file, err := os.Create("simple_layout_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	_, err = file.WriteString(result.String())
	if err != nil {
		fmt.Printf("ファイル書き込みエラー: %v\n", err)
		return
	}
	
	fmt.Println("シンプルなレイアウトをsimple_layout_output.txtに出力しました")
	fmt.Printf("使用レイヤー数: %d\n", ls.GetLayerCount())
}