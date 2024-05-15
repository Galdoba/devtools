package grid

import "fmt"

type coord struct {
	x      int
	y      int
	err    error
	parent Coord
}

const (
	BaseCoordError = "base coord error"
	NeighbourNot   = "not a neighbour"
	NeighbourRight = "neighbour right"
	NeighbourDown  = "neighbour down"
	NeighbourLeft  = "neighbour left"
	NeighbourUp    = "neighbour up"
	NeighbourError = "neighbour error"
)

func NewCoord(x, y int) Coord {
	c := coord{}
	c.x = x
	c.y = y
	if c.x < 0 {
		c.err = fmt.Errorf("invalid x coordinate: %v is lower than 0", c.x)
	}
	if c.y < 0 {
		c.err = fmt.Errorf("invalid y coordinate: %v is lower than 0", c.y)
	}
	return &c
}

type Coord interface {
	Coordinates() (int, int)
	Neighbouring(Coord) string
	SubjucateTo(Coord) error
	Parent() Coord
	CoordError() error
}

func (c *coord) Coordinates() (int, int) {
	return c.x, c.y
}

func (c *coord) Neighbouring(neib Coord) string {
	if c.err != nil {
		return BaseCoordError
	}
	if neib.CoordError() != nil {
		return NeighbourError
	}
	nX, nY := neib.Coordinates()
	if c.x-nX == -1 && c.y-nY == 0 {
		return NeighbourRight
	}
	if c.x-nX == 0 && c.y-nY == -1 {
		return NeighbourDown
	}
	if c.x-nX == 1 && c.y-nY == 0 {
		return NeighbourLeft
	}
	if c.x-nX == 0 && c.y-nY == 1 {
		return NeighbourUp
	}
	return NeighbourNot
}

func (c *coord) SubjucateTo(n Coord) error {
	nv := c.Neighbouring(n)
	switch nv {
	default:
		return fmt.Errorf("can't subjugate to %v", nv)
	case NeighbourLeft, NeighbourUp:
	}

	c.parent = n
	return nil
}

func (c *coord) Parent() Coord {
	p := c.parent
	if p == nil {
		return c
	}
	return p.Parent()
}

func (c *coord) CoordError() error {
	return c.err
}

func FormRelations(coords ...Coord) error {
	coordMap := make(map[string]int)
	for _, c := range coords {
		x, y := c.Coordinates()
		key := fmt.Sprintf("%v:%v", x, y)
		if coordMap[key] != 0 {
			return fmt.Errorf("can't form reltions: coords duplicated (%v)", key)
		}
		coordMap[key]++
	}
	if len(coordMap)%2 != 0 {
		return fmt.Errorf("can't form reltions: must have even sets if coords")
	}
	return nil
}
