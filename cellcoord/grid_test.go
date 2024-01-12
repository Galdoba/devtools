package coordinates

import (
	"fmt"
	"testing"

	"github.com/Galdoba/devtools/cellcoord/spreadsheet"
	"github.com/Galdoba/devtools/cellcoord/square"
)

func TestGrid(t *testing.T) {
	gr, err := NewGrid(square.GridType)
	fmt.Println(gr, err)
	errAdd := gr.AddCoords(1, 1)
	errAdd = gr.AddCoords(1, 2)
	fmt.Println(gr, errAdd)

	gr2, err := NewGrid(spreadsheet.GridType)
	gr2.AddCoords(1, 1)

	s1, _ := square.NewCoordinates(1, 1)
	s2, _ := square.NewCoordinates(2, 1)
	s3, _ := square.NewCoordinates(1, 1)
	sp1, _ := spreadsheet.NewCoordinates(1, 1)
	fmt.Println(Equal(s1, s2))
	fmt.Println(Equal(s1, s3))
	fmt.Println(Equal(s1, sp1))
	fmt.Println(gr2, errAdd)
	s4 := Add(s1, s2)
	fmt.Println(s4)

}
