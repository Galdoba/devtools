package main

import (
	"fmt"

	"github.com/Galdoba/devtools/decidion/operator"
)

func main() {

	if !operator.Confirm("test 2") {
		fmt.Println("bad")
	}
	fmt.Println("good	")
}
