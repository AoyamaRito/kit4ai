package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
)

func drawBox(c *canvas.Canvas, x1, y1, x2, y2 int, title string) {
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
	
	if title != "" {
		titleRunes := []rune(title)
		for i, r := range titleRunes {
			if x1+2+i <= x2-1 {
				c.ReplaceChar(x1+2+i, y1+1, r)
			}
		}
	}
}

func main() {
	fmt.Println("Web Layout Demo - Using Layer System")
	fmt.Println("====================================")
	
	ls := canvas.NewLayerSystem()
	
	headerLayerID := ls.AddLayer()
	headerLayer, _ := ls.GetLayer(headerLayerID)
	ls.SetZOrder(headerLayerID, 1)
	drawBox(headerLayer.Canvas, 0, 0, 79, 4, "ヘッダー - ナビゲーション")
	
	sidebarLayerID := ls.AddLayer()
	sidebarLayer, _ := ls.GetLayer(sidebarLayerID)
	ls.SetZOrder(sidebarLayerID, 2)
	drawBox(sidebarLayer.Canvas, 0, 5, 19, 25, "サイドバー")
	
	menuItems := []string{"ホーム", "プロフィール", "設定", "ログアウト"}
	for i, item := range menuItems {
		itemRunes := []rune(item)
		for j, r := range itemRunes {
			if 2+j <= 17 {
				sidebarLayer.Canvas.ReplaceChar(2+j, 7+i*2, r)
			}
		}
	}
	
	mainLayerID := ls.AddLayer()
	mainLayer, _ := ls.GetLayer(mainLayerID)
	ls.SetZOrder(mainLayerID, 3)
	drawBox(mainLayer.Canvas, 20, 5, 79, 25, "メインコンテンツエリア")
	
	contentItems := []string{
		"投稿タイトル１：最新のお知らせ",
		"投稿タイトル２：重要な更新情報",
		"投稿タイトル３：新機能のご紹介",
	}
	
	for i, item := range contentItems {
		itemRunes := []rune(item)
		for j, r := range itemRunes {
			if 22+j <= 77 {
				mainLayer.Canvas.ReplaceChar(22+j, 8+i*3, r)
			}
		}
		
		for x := 22; x <= 77; x++ {
			mainLayer.Canvas.ReplaceChar(x, 9+i*3, '-')
		}
	}
	
	footerLayerID := ls.AddLayer()
	footerLayer, _ := ls.GetLayer(footerLayerID)
	ls.SetZOrder(footerLayerID, 4)
	drawBox(footerLayer.Canvas, 0, 26, 79, 30, "フッター - コピーライト情報")
	
	modalLayerID := ls.AddLayer()
	modalLayer, _ := ls.GetLayer(modalLayerID)
	ls.SetZOrder(modalLayerID, 10)
	
	for y := 10; y <= 20; y++ {
		for x := 25; x <= 55; x++ {
			modalLayer.Canvas.ReplaceChar(x, y, ' ')
		}
	}
	
	drawBox(modalLayer.Canvas, 25, 10, 55, 20, "モーダルダイアログ")
	
	modalText := []string{
		"この操作を実行しますか？",
		"",
		"[はい] [いいえ]",
	}
	
	for i, text := range modalText {
		textRunes := []rune(text)
		startX := 27
		if i == 2 {
			startX = 35
		}
		for j, r := range textRunes {
			if startX+j <= 53 {
				modalLayer.Canvas.ReplaceChar(startX+j, 13+i, r)
			}
		}
	}
	
	result := ls.Composite()
	
	fmt.Println("\n完成したWebレイアウト:")
	fmt.Println(result.String())
	
	fmt.Printf("\n使用レイヤー数: %d\n", ls.GetLayerCount())
	fmt.Printf("レイヤーID: %v\n", ls.GetLayerIDs())
	
	fmt.Println("\n各レイヤーの内容:")
	fmt.Println("================")
	
	layers := []struct {
		id   int
		name string
	}{
		{headerLayerID, "ヘッダー"},
		{sidebarLayerID, "サイドバー"},
		{mainLayerID, "メイン"},
		{footerLayerID, "フッター"},
		{modalLayerID, "モーダル"},
	}
	
	for _, layer := range layers {
		fmt.Printf("\n%s レイヤー (ID: %d):\n", layer.name, layer.id)
		layerObj, _ := ls.GetLayer(layer.id)
		layerContent := layerObj.Canvas.String()
		if layerContent == "" {
			fmt.Println("(空のレイヤー)")
		} else {
			fmt.Println(layerContent)
		}
	}
}