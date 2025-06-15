package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
	"os"
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
	
	result := ls.Composite()
	
	file, err := os.Create("web_layout_output.txt")
	if err != nil {
		fmt.Printf("ファイル作成エラー: %v\n", err)
		return
	}
	defer file.Close()
	
	_, err = file.WriteString("ASCII Art Web Layout Specification\n")
	if err != nil {
		fmt.Printf("ファイル書き込みエラー: %v\n", err)
		return
	}
	
	_, err = file.WriteString("====================================\n\n")
	if err != nil {
		fmt.Printf("ファイル書き込みエラー: %v\n", err)
		return
	}
	
	_, err = file.WriteString(result.String())
	if err != nil {
		fmt.Printf("ファイル書き込みエラー: %v\n", err)
		return
	}
	
	_, err = file.WriteString("\n\n")
	if err != nil {
		fmt.Printf("ファイル書き込みエラー: %v\n", err)
		return
	}
	
	_, err = file.WriteString(fmt.Sprintf("Generated using %d layers\n", ls.GetLayerCount()))
	if err != nil {
		fmt.Printf("ファイル書き込みエラー: %v\n", err)
		return
	}
	
	_, err = file.WriteString(fmt.Sprintf("Layer IDs: %v\n", ls.GetLayerIDs()))
	if err != nil {
		fmt.Printf("ファイル書き込みエラー: %v\n", err)
		return
	}
	
	fmt.Println("ASCII artをweb_layout_output.txtに出力しました")
	fmt.Printf("使用レイヤー数: %d\n", ls.GetLayerCount())
}