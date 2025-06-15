package canvas

import (
	"errors"
	"strings"
)

const (
	DefaultWidth  = 80
	DefaultHeight = 100
)

type Canvas struct {
	Width  int
	Height int
	Grid   [][]rune
}

func NewCanvas() *Canvas {
	c := &Canvas{
		Width:  DefaultWidth,
		Height: DefaultHeight,
		Grid:   make([][]rune, DefaultHeight),
	}
	
	for i := range c.Grid {
		c.Grid[i] = make([]rune, DefaultWidth)
		for j := range c.Grid[i] {
			c.Grid[i][j] = ' '
		}
	}
	
	return c
}

func (c *Canvas) ReplaceChar(x, y int, char rune) error {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return errors.New("coordinates out of bounds")
	}
	
	c.Grid[y][x] = char
	return nil
}

func (c *Canvas) GetChar(x, y int) (rune, error) {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return 0, errors.New("coordinates out of bounds")
	}
	
	return c.Grid[y][x], nil
}

func (c *Canvas) Clear() {
	for i := range c.Grid {
		for j := range c.Grid[i] {
			c.Grid[i][j] = ' '
		}
	}
}

func (c *Canvas) String() string {
	lines := make([]string, 0, c.Height)
	
	lastNonEmptyLine := -1
	for i := 0; i < c.Height; i++ {
		line := string(c.Grid[i])
		if strings.TrimSpace(line) != "" {
			lastNonEmptyLine = i
		}
	}
	
	for i := 0; i <= lastNonEmptyLine; i++ {
		lines = append(lines, string(c.Grid[i]))
	}
	
	if len(lines) == 0 {
		return ""
	}
	
	return strings.Join(lines, "\n")
}