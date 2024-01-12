package grid

import (
	"testing"
)

func TestCoord(t *testing.T) {
	testX := []int{-2, -1, 0, 1, 2, 3}
	testY := []int{-2, -1, 0, 1, 2, 3}
	offX := []int{-2, -1, 0, 1, 2}
	offY := []int{-2, -1, 0, 1, 2}
	//outCoord := 4
	for _, x := range testX {
		for _, y := range testY {
			//fmt.Println("======TEST", x, y)
			c := NewCoord(x, y)
			cx, cy := c.Coordinates()
			cErr := c.CoordError()
			if cx < 0 && cErr == nil {
				t.Errorf("expected to have error")
			}
			if cy < 0 && cErr == nil {
				t.Errorf("expected to have error")
			}
			for _, ox := range offX {
				for _, oy := range offY {
					//fmt.Println("======NEIB", x+ox, y+oy)
					//nx, ny := x+ox, y+oy
					// n := &coord{}
					// n.x, n.y = nx, ny
					// if nx >= outCoord || ny >= outCoord {
					// 	n.err = fmt.Errorf("out of grid")
					// 	continue
					// }

					n := NewCoord(x+ox, y+oy)

					err := c.SubjucateTo(n)
					if err != nil {
						t.Errorf("subjugation err: %v", err.Error())
					}
					//fmt.Println(c, n, c.Neighbouring(n))
					switch c.Neighbouring(n) {
					default:
					case NeighbourError, BaseCoordError, NeighbourNot:
						continue
					}

					//		fmt.Println("======SUCCESS NEIB", x+ox, y+oy)

				}
			}

		}
	}
}
