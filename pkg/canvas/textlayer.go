package canvas

import (
	"unicode"
)

// TextLayer は全角文字対応の専用テキストレイヤー
type TextLayer struct {
	*Canvas
}

// NewTextLayer creates a new text layer with full-width character support
func NewTextLayer() *TextLayer {
	return &TextLayer{
		Canvas: NewCanvas(),
	}
}

// IsFullWidth 全角文字かどうかを判定
func (tl *TextLayer) IsFullWidth(r rune) bool {
	// Unicode罫線文字は半角として扱う（表示安定性のため）
	if r >= 0x2500 && r <= 0x257F {
		return false // Box Drawing文字は半角扱い
	}
	
	return unicode.Is(unicode.Han, r) ||
		unicode.Is(unicode.Hiragana, r) ||
		unicode.Is(unicode.Katakana, r) ||
		(r >= 0xFF01 && r <= 0xFF5E) ||
		(r >= 0x3000 && r <= 0x303F)
}

// GetCharWidth 文字の表示幅を取得
func (tl *TextLayer) GetCharWidth(r rune) int {
	if tl.IsFullWidth(r) {
		return 2
	}
	return 1
}

// GetTextWidth テキストの総表示幅を計算
func (tl *TextLayer) GetTextWidth(text string) int {
	width := 0
	for _, r := range []rune(text) {
		width += tl.GetCharWidth(r)
	}
	return width
}

// WriteTextWithWidth 全角文字対応のテキスト書き込み（連続配置、空白禁止）
func (tl *TextLayer) WriteTextWithWidth(x, y int, text string) error {
	currentX := x
	textRunes := []rune(text)
	
	for _, r := range textRunes {
		// 境界チェック
		if currentX >= tl.Width || y < 0 || y >= tl.Height {
			break
		}
		
		// 全角・半角に関係なく連続配置（空白を作らない）
		tl.ReplaceChar(currentX, y, r)
		currentX += 1 // 常に1つずつ進める
	}
	
	return nil
}

// WriteTextSimple 従来通りの1文字=1グリッド配置（互換性用）
func (tl *TextLayer) WriteTextSimple(x, y int, text string) error {
	for i, r := range []rune(text) {
		if x+i >= tl.Width || y < 0 || y >= tl.Height {
			break
		}
		tl.ReplaceChar(x+i, y, r)
	}
	return nil
}

// WriteTextCentered 指定した幅内にテキストを中央寄せで配置
func (tl *TextLayer) WriteTextCentered(x, y, width int, text string) error {
	textWidth := tl.GetTextWidth(text)
	if textWidth > width {
		return tl.WriteTextWithWidth(x, y, text)
	}
	
	offset := (width - textWidth) / 2
	return tl.WriteTextWithWidth(x+offset, y, text)
}

// WriteTextRight 指定した幅内にテキストを右寄せで配置
func (tl *TextLayer) WriteTextRight(x, y, width int, text string) error {
	textWidth := tl.GetTextWidth(text)
	if textWidth > width {
		return tl.WriteTextWithWidth(x, y, text)
	}
	
	offset := width - textWidth
	return tl.WriteTextWithWidth(x+offset, y, text)
}

// FillBox ボックス内を指定文字で埋める（背景用）
func (tl *TextLayer) FillBox(x1, y1, x2, y2 int, char rune) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			if x >= 0 && x < tl.Width && y >= 0 && y < tl.Height {
				tl.ReplaceChar(x, y, char)
			}
		}
	}
}