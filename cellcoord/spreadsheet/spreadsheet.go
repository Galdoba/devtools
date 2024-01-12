package spreadsheet

import (
	"fmt"
	"strings"
)

const (
	GridType    = "Spreadsheet"
	DirectUP    = 221
	DirectRIGHT = 222
	DirectDOWN  = 223
	DirectLEFT  = 224
)

// coord (spreadsheet) is coordinates object for grid used in spreadsheet tables
// satisfies Coords (interface) of parent packege
type coord struct {
	row      int
	col      int
	cellName string
	cellRC   string
}

// NewCoordinates - creates Spreadsheet coords
func NewCoordinates(vals ...int) (coord, error) {
	if len(vals) != 2 {
		return coord{}, fmt.Errorf("can't create %v coords: vals provided %v (expect 2)", GridType, len(vals))
	}
	c := coord{}
	c.row = vals[0]
	c.col = vals[1]
	c.cellName = fmt.Sprintf("%v%v", int2Str(c.col), c.row)
	c.cellRC = fmt.Sprintf("R%vC%v", c.row, c.col)
	return c, nil
}

func SetCoords(row, col int) coord {
	return coord{row, col, fmt.Sprintf("%v%v", int2Str(col), row), fmt.Sprintf("R%vC%v", row, col)}
}

// GridType - returns grid composition this coord belongs to
func (c coord) GridType() string {
	return GridType
}

func (c coord) Directions() map[int][]int {
	direction := make(map[int][]int)
	direction[DirectUP] = []int{-1, 0}
	direction[DirectRIGHT] = []int{0, 1}
	direction[DirectDOWN] = []int{1, 0}
	direction[DirectLEFT] = []int{0, -1}
	return direction
}

func (c coord) Validate() error {
	if c.col < 1 {
		return fmt.Errorf("can't have column number less than 1")
	}
	if c.row < 1 {
		return fmt.Errorf("can't have row number less than 1")
	}
	if c.cellRC != fmt.Sprintf("R%vC%v", c.row, c.col) {
		return fmt.Errorf("cell RC name is incorrect")
	}
	if c.cellName != fmt.Sprintf("%v%v", int2Str(c.col), c.row) {
		return fmt.Errorf("cell Name is incorrect")
	}
	return nil
}

func (c coord) Vals() []int {
	return []int{c.col, c.row}
}

func (c coord) String() string {
	return fmt.Sprintf("%v:(%v;%v)", GridType, c.row, c.col)
}

// local
func (c coord) Name() string {
	return c.cellName
}

func (c coord) NameAlt() string {
	return c.cellRC
}

//helpers

func int2Str(i int) string {
	total := i
	out := []string{}
	v := total / 26
	vv := v / 26
	s := total % 26
	out = append(out, i2v(vv))
	out = append(out, i2v(v))
	out = append(out, i2v(s))
	return strings.Join(out, "")

}

func i2v(i int) string {
	switch i {
	default:
		return "?"
	case 0:
		return ""
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	case 5:
		return "E"
	case 6:
		return "F"
	case 7:
		return "G"
	case 8:
		return "H"
	case 9:
		return "I"
	case 10:
		return "J"
	case 11:
		return "K"
	case 12:
		return "L"
	case 13:
		return "M"
	case 14:
		return "N"
	case 15:
		return "O"
	case 16:
		return "P"
	case 17:
		return "Q"
	case 18:
		return "R"
	case 19:
		return "S"
	case 20:
		return "T"
	case 21:
		return "U"
	case 22:
		return "V"
	case 23:
		return "W"
	case 24:
		return "X"
	case 25:
		return "Y"
	case 26:
		return "Z"
	}
}
