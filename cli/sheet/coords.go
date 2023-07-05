package sheet

type coord struct {
	r int
	c int
}

// type Coord interface {
// 	Current() (int, int)
// 	Up() (int, int)
// 	Down() (int, int)
// 	Left() (int, int)
// 	Right() (int, int)
// 	Valid(int, int) bool
// }

func NewCoord(r, c int) coord {
	return coord{r, c}
}

func (c *coord) Current() (int, int) {
	return c.r, c.c
}

func (c *coord) Up() (int, int) {
	return c.r - 1, c.c
}

func (c *coord) Down() (int, int) {
	return c.r + 1, c.c
}

func (c *coord) Left() (int, int) {
	return c.r, c.c - 1
}

func (c *coord) Right() (int, int) {
	return c.r, c.c + 1
}

func (c *coord) Valid(maxRow int, maxCol int) bool {
	if c.r < 1 {
		return false
	}
	if c.r > maxRow {
		return false
	}
	if c.c < 1 {
		return false
	}
	if c.c > maxCol {
		return false
	}
	return true
}
