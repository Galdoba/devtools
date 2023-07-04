package sheet

type coord struct {
	r int
	c int
}

func NewCoord(r, c int) *coord {
	return &coord{r, c}
}

type Coords interface {
	RowCol() (int, int)
	//String() string
}

func (c *coord) RowCol() (int, int) {
	return c.r, c.c
}
