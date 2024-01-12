package square

import (
	"fmt"
)

const (
	AxisX    = "X"
	AxisY    = "Y"
	GridType = "Square"
	DirectN  = 211
	DirectE  = 212
	DirectS  = 213
	DirectW  = 214
)

type coord struct {
	x int
	y int
}

func NewCoordinates(vals ...int) (coord, error) {
	if len(vals) != 2 {
		return coord{}, fmt.Errorf("can't create new square coords: %v axis provided", len(vals))
	}
	return coord{vals[0], vals[1]}, nil
}

func SetCoord(x, y int) coord {
	return coord{x, y}
}

func (c coord) GridType() string {
	return GridType
}

func (c coord) Validate() error {
	return nil
}

func (c coord) Vals() []int {
	return []int{c.x, c.y}
}

func (c coord) Directions() map[int][]int {
	directions := make(map[int][]int)
	directions[DirectN] = []int{0, 1}
	directions[DirectE] = []int{1, 0}
	directions[DirectS] = []int{0, -1}
	directions[DirectW] = []int{-1, 0}
	return directions
}

func (c coord) String() string {
	return fmt.Sprintf("%v:(%v;%v)", c.GridType(), c.x, c.y)
}
