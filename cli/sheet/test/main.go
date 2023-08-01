// /*
// List 1:
// R0C0| R1 | R2 | R3 | R4 |
// =========================
//  C1 |    |    |    |    |
// ----+----+----+----+----|
//  C2 |    |    |    |    |
// ----+----+----+----+----|
//  C3 |    |    |    |    |
// ----+----+----+----+----|
//  C4 |    |    |    |    |
// =========================

// */
package main

import (
	"fmt"

	"github.com/Galdoba/devtools/cli/sheet"
)

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {

	lst := sheet.NewList("List 1", 3, 5)
	sht := sheet.NewSheet(lst)
	sheet.DebugFill(sht.List("List 1"))
	//sht.List("List 1").Print()
	sheet.Print(sht.List("List 1"))
	fmt.Println("")
	sht.List("List 1").Cell(1, 2).Action(sheet.TOGGLE_CURSOR)
	sht.List("List 1").Cell(1, 1).Write("===!!!!!======!")
	sheet.Print(sht.List("List 1"))
}
