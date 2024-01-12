package coordinates

import (
	"fmt"

	"github.com/Galdoba/devtools/cellcoord/spreadsheet"
	"github.com/Galdoba/devtools/cellcoord/square"
)

type Grid struct {
	gridtype string
	coords   map[string]Coords
}

// Coord - common functions to control cursor on different grids layout
type Coords interface {
	Directions() map[int][]int
	GridType() string
	Validate() error
	Vals() []int
	String() string
}

func NewGrid(gridtype string) (*Grid, error) {
	gr := Grid{}
	gr.gridtype = gridtype
	switch gridtype {
	default:
		return nil, fmt.Errorf("can't create new grid: unknown type '%v'", gridtype)
	case square.GridType:
	case spreadsheet.GridType:
	}
	gr.coords = make(map[string]Coords)
	return &gr, nil
}

func (gr *Grid) AddCoords(vals ...int) error {
	switch gr.gridtype {
	default:
		return fmt.Errorf("can't add coords: unknown gridtype '%v'", gr.gridtype)
	case square.GridType:
		crd, err := square.NewCoordinates(vals...)
		if err != nil {
			return fmt.Errorf("can't add coords: %v", err.Error())
		}
		if err = gr.assertAvailability(crd); err != nil {
			return fmt.Errorf("can't add coords: %v", err.Error())
		}
		gr.coords[crd.String()] = crd
	case spreadsheet.GridType:
		crd, err := spreadsheet.NewCoordinates(vals...)
		if err != nil {
			return fmt.Errorf("can't add coords: %v", err.Error())
		}
		if err = gr.assertAvailability(crd); err != nil {
			return fmt.Errorf("can't add coords: %v", err.Error())
		}
		gr.coords[crd.String()] = crd
	}
	return nil
}

func (gr *Grid) assertAvailability(c Coords) error {
	code := c.String()
	if _, ok := gr.coords[code]; ok {
		return fmt.Errorf("coords %v already exist", code)
	}
	return nil
}

//COMMONS

func AxisNames(gridType string) []string {
	switch gridType {
	default:
		return nil
	case square.GridType:
		return []string{square.AxisX, square.AxisY}
	}
}

func Equal(c1, c2 Coords) bool {
	if c1.GridType() != c2.GridType() {
		return false
	}
	vals1 := c1.Vals()
	vals2 := c2.Vals()
	if len(vals1) != len(vals2) {
		return false
	}
	for i, v1 := range vals1 {
		if vals2[i] != v1 {
			return false
		}
	}
	return true
}

func Add(a, b Coords) Coords {
	vals1 := a.Vals()
	vals2 := b.Vals()
	switch a.GridType() {
	case square.GridType:
		sqrc := square.SetCoord(vals1[0]+vals2[0], vals1[1]+vals2[1])
		return sqrc
	case spreadsheet.GridType:
		spcrd := spreadsheet.SetCoords(vals1[0]+vals1[1], vals1[1]+vals2[1])
		return spcrd
	}
	return nil
}

func Neighbors(c Coords) []Coords {
	return []Coords{c}
}
