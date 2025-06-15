package canvas

import (
	"unicode"
)

// 全角文字かどうかを判定する関数
func IsFullWidth(r rune) bool {
	return unicode.Is(unicode.Han, r) ||          // 漢字
		unicode.Is(unicode.Hiragana, r) ||        // ひらがな
		unicode.Is(unicode.Katakana, r) ||        // カタカナ
		(r >= 0xFF01 && r <= 0xFF5E) ||          // 全角英数字記号
		(r >= 0x3000 && r <= 0x303F)             // 全角スペース・句読点
}

// 文字の表示幅を取得
func GetCharWidth(r rune) int {
	if IsFullWidth(r) {
		return 2
	}
	return 1
}

// テキストの総表示幅を計算
func GetTextWidth(text string) int {
	width := 0
	for _, r := range []rune(text) {
		width += GetCharWidth(r)
	}
	return width
}

// 全角対応のテキスト書き込み関数
func (c *Canvas) WriteText(x, y int, text string) error {
	currentX := x
	textRunes := []rune(text)
	
	for _, r := range textRunes {
		// 境界チェック
		if currentX >= c.Width || y < 0 || y >= c.Height {
			break
		}
		
		if IsFullWidth(r) {
			// 全角文字は2文字分のスペースを使用
			if currentX < c.Width {
				c.ReplaceChar(currentX, y, r)
				currentX++
			}
			// 次の位置もこの文字の一部として扱う（空白で上書きしない）
			if currentX < c.Width {
				// 全角文字の右半分は特別な処理をしない（既存の文字をそのまま残す）
				currentX++
			}
		} else {
			// 半角文字は1文字分
			c.ReplaceChar(currentX, y, r)
			currentX++
		}
	}
	
	return nil
}

// 指定した幅内にテキストを中央寄せで配置
func (c *Canvas) WriteTextCentered(x, y, width int, text string) error {
	textWidth := GetTextWidth(text)
	if textWidth > width {
		// テキストが幅を超える場合は左寄せ
		return c.WriteText(x, y, text)
	}
	
	// 中央寄せの位置を計算
	offset := (width - textWidth) / 2
	return c.WriteText(x+offset, y, text)
}

// 指定した幅内にテキストを右寄せで配置
func (c *Canvas) WriteTextRight(x, y, width int, text string) error {
	textWidth := GetTextWidth(text)
	if textWidth > width {
		// テキストが幅を超える場合は左寄せ
		return c.WriteText(x, y, text)
	}
	
	// 右寄せの位置を計算
	offset := width - textWidth
	return c.WriteText(x+offset, y, text)
}