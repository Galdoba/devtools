package sheet

import (
	"fmt"

	"github.com/Galdoba/devtools/text"
)

const (
	aligh_left    = "ALeft"
	aligh_right   = "ARight"
	aligh_center  = "ACenter"
	TOGGLE_CURSOR = "Toggle_Cursor"
)

type cell struct {
	crd       coord
	content   string
	aligh     string
	width     int
	height    int
	selected  bool
	cursor    bool
	bold      bool
	faint     bool
	italic    bool
	underline bool
	crossout  bool
}

func NewCell(r, c int) cell {
	crd := NewCoord(r, c)
	cl := cell{}
	cl.crd = crd
	cl.aligh = aligh_left
	return cl
}

type Cell interface {
	Read() string
	Write(string)
	Coord() coord
	Size() (int, int)
	Alignment() string
	Cursor() bool
	Action(string) error
}

func (c *cell) Read() string {
	return c.content
}

func (c *cell) Write(s string) {
	s = text.EnsureASCII(s)
	c.content = s
}

func (c *cell) Coord() coord {
	return c.crd
}

func (c *cell) Alignment() string {
	return c.aligh
}

func (c *cell) Cursor() bool {
	return c.cursor
}

func (c *cell) Size() (int, int) {
	return c.width, c.height
}

func AlignText(cell Cell) string {
	width, _ := cell.Size()
	text := cell.Read()
	switch cell.Alignment() {
	case aligh_left:
		for len(text) < width {
			text += " "
		}
	case aligh_right:
		for len(text) < width {
			text = " " + text
		}
	case aligh_center:
		for len(text) < width {
			switch len(text) % 2 {
			case 0:
				text = " " + text
			case 1:
				text += " "
			}
		}
	}
	return text
}

func (c *cell) Action(action string) error {
	switch action {
	default:
		return fmt.Errorf("unknown action '%v'", action)
	case TOGGLE_CURSOR:
		c.cursor = !c.cursor
	}
	return nil
}
