package directory

import (
	"fmt"
	"testing"
)

func TestReadDir(t *testing.T) {
	i := 0
	for _, tr := range Tree("d:\\IN\\IN_testInput\\") {
		files := ListFiles(tr)
		for _, fl := range files {
			i++
			fmt.Println(i, fl)
		}
	}
}
