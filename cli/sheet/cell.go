package sheet

type cell struct {
	crd      coord
	content  string
	selected bool
	cursor   bool
}

func NewCell(r, c int) cell {
	crd := NewCoord(r, c)
	return cell{crd, "", false, false}
}

type Cell interface {
	Read() string
	Write(string)
	Coord() coord
}

func (c *cell) Read() string {
	return c.content
}

func (c *cell) Write(s string) {
	c.content = s
}

func (c *cell) Coord() coord {
	return c.crd
}
