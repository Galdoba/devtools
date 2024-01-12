package square

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	sq := coord{1, 2}
	fmt.Println(sq.Vals())
	fmt.Println(sq.String())
	fmt.Println(sq.GridType())
	fmt.Println(sq.Validate())
}
