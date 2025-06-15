package main

import (
	"fmt"
	"kit4ai/pkg/canvas"
)

func main() {
	fmt.Println("横幅100文字テスト")
	fmt.Println("=================")
	
	// 現在のデフォルト値を確認
	fmt.Printf("現在のデフォルト幅: %d\n", canvas.DefaultWidth)
	fmt.Printf("現在のデフォルト高さ: %d\n", canvas.DefaultHeight)
	
	// 100文字幅での問題を確認
	fmt.Println("\n100文字幅での問題チェック:")
	
	// 1. 表示問題
	fmt.Println("1. 表示範囲の問題:")
	fmt.Println("   - 多くのターミナルは80文字が標準")
	fmt.Println("   - 100文字は横スクロールが発生する可能性")
	fmt.Println("   - プリンタのA4用紙に収まらない")
	
	// 2. 互換性問題
	fmt.Println("2. 互換性の問題:")
	fmt.Println("   - レガシーシステムは80文字前提")
	fmt.Println("   - 一部のテキストエディタで折り返し")
	fmt.Println("   - メール本文での表示崩れ")
	
	// 3. 実際のテスト
	fmt.Println("3. 実際の100文字ライン:")
	line100 := ""
	for i := 0; i < 100; i++ {
		line100 += fmt.Sprintf("%d", i%10)
	}
	fmt.Println(line100)
	fmt.Println("↑この行が画面内に収まっているかチェック")
	
	// 4. ByteCanvasで100文字テスト
	fmt.Println("\n4. ByteCanvasでの100文字テスト:")
	
	// カスタム幅のByteCanvas作成
	bc := &canvas.ByteCanvas{
		Width:  100,
		Height: 20,
		Grid:   make([][]byte, 20),
	}
	
	for i := range bc.Grid {
		bc.Grid[i] = make([]byte, 100)
		for j := range bc.Grid[i] {
			bc.Grid[i][j] = ' '
		}
	}
	
	// 100文字幅のボックス描画
	bc.DrawBox(0, 0, 99, 10)
	bc.WriteBytes(2, 2, "100文字幅のテストボックス")
	bc.WriteBytes(2, 4, "Left Side")
	bc.WriteBytes(50, 4, "Center")
	bc.WriteBytes(90, 4, "Right")
	
	// 位置マーカー
	for i := 0; i < 100; i += 10 {
		marker := fmt.Sprintf("%d", i/10)
		bc.WriteBytes(i, 8, marker)
	}
	
	fmt.Println("100文字幅ボックス:")
	fmt.Println(bc.String())
	
	// 5. 推奨案
	fmt.Println("\n5. 推奨案:")
	fmt.Println("   - 標準版: 80文字（互換性重視）")
	fmt.Println("   - ワイド版: 100文字（詳細表示用）")
	fmt.Println("   - 設定可能にして用途別使い分け")
	
	// 6. 実用例での比較
	fmt.Println("\n6. 実用例での比較:")
	
	// 80文字版
	bc80 := canvas.NewByteCanvas()
	bc80.DrawBox(0, 0, 79, 8)
	bc80.WriteBytes(2, 2, "Standard 80-char Dashboard")
	bc80.WriteBytes(2, 4, "Users: 1,234")
	bc80.WriteBytes(20, 4, "Sales: $5,678")
	bc80.WriteBytes(40, 4, "Orders: 89")
	bc80.WriteBytes(60, 4, "Status: OK")
	
	fmt.Println("80文字版:")
	fmt.Println(bc80.String())
	
	// 100文字版
	bc100 := &canvas.ByteCanvas{
		Width:  100,
		Height: 10,
		Grid:   make([][]byte, 10),
	}
	
	for i := range bc100.Grid {
		bc100.Grid[i] = make([]byte, 100)
		for j := range bc100.Grid[i] {
			bc100.Grid[i][j] = ' '
		}
	}
	
	bc100.DrawBox(0, 0, 99, 8)
	bc100.WriteBytes(2, 2, "Wide 100-char Dashboard")
	bc100.WriteBytes(2, 4, "Users: 1,234")
	bc100.WriteBytes(20, 4, "Sales: $5,678")
	bc100.WriteBytes(40, 4, "Orders: 89")
	bc100.WriteBytes(60, 4, "Status: OK")
	bc100.WriteBytes(78, 4, "More Data: Available")
	
	fmt.Println("\n100文字版:")
	fmt.Println(bc100.String())
	
	fmt.Println("\n結論:")
	fmt.Println("- 100文字は技術的に問題なし")
	fmt.Println("- 表示環境によっては見切れる可能性")
	fmt.Println("- 用途に応じて選択するのがベスト")
}