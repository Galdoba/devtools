package text

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	input := "0123456789"
	output := SetLength(input, 4)
	fmt.Println(input, "|")
	fmt.Println(output, "|")
}
