package table

import "fmt"

type Cell struct {
	PosCol  int
	PosRow  int
	Content string
	// FGcol       byte
	// BGcol       byte
	// Independent bool              //есть ли связи с другими ячейками
	// ParentX     bool              //родитель слева
	// ParentY     bool              //родитель сверху
	// Action      map[string]string //при получении допустимого тригера (key) запускаем действие (val)
	err error
}

func NewCell() *Cell {
	return &Cell{}
}

//WithPosition - sets col and ros as cell coordinates
func (c *Cell) WithPosition(col, row int) *Cell {
	if c.err != nil {
		return c
	}
	if col < 0 {
		c.err = fmt.Errorf("cell column must be > 0")
		return c
	}
	if row < 0 {
		c.err = fmt.Errorf("cell row must be > 0")
		return c
	}
	////
	c.PosCol = col
	c.PosRow = row
	return c
}

//WithContent - sets content
func (c *Cell) WithContent(cnt string) *Cell {
	if c.err != nil {
		return c
	}
	////
	c.Content = cnt
	return c
}

func (c *Cell) String() string {
	return c.Content
}
