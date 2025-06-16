package canvas

// SetByteAt sets a single byte at the specified position
func (bc *ByteCanvas) SetByteAt(x, y int, b byte) {
	if x >= 0 && x < bc.Width && y >= 0 && y < bc.Height {
		bc.Grid[y][x] = b
	}
}