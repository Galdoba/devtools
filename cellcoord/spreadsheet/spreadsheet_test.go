package spreadsheet

import (
	"fmt"
	"testing"
)

func TestSpreadShhet(t *testing.T) {
	sq, err := NewCoordinates(1, 2)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err.Error()))
	}
	fmt.Println(sq.Vals())
	fmt.Println(sq.String())
	fmt.Println(sq.GridType())
	err = sq.Validate()
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err.Error()))
	}
	fmt.Println(sq.cellName, sq.cellRC)
}
