package autodoc

import (
	"fmt"
	"testing"
)

func TestParsing(t *testing.T) {
	vs, err := stateFromString("1.6.19")
	fmt.Println(err)
	fmt.Println(vs)
}
