package directory

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	i := 0
	//for _, tr := range Tree(`/home/galdoba/Videos/`) {
	root := `/home/galdoba/Videos`
	root, err := filepath.Abs(root)
	if err != nil {
		fmt.Println(err.Error())
	}

	dir, files, err := List(root)

	if err != nil {
		fmt.Println(err.Error())
	}
	for _, fl := range files {
		if _, err := os.Stat(dir + fl); err == nil {
			// path/to/whatever exists
			fmt.Println("No Err")
		} else if errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does *not* exist
			fmt.Println("   " + err.Error())
		} else {
			// Schrodinger: file may or may not exist. See err for details.
			fmt.Println("   " + err.Error())
			// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence

		}

		fmt.Println(i, fl)
	}
	//}
}

func TestTree(t *testing.T) {

	root := `/home/galdoba/Videos`
	tree := Tree(root)
	for i, branch := range tree {
		fmt.Println(i, branch)
	}
}
