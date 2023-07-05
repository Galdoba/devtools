package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/devtools/cli/sheet"
)

func main() {
	sht := sheet.NewSheet(sheet.NewList("List 1", 3, 5))
	sht.List("List 1").Print()
	fmt.Print("R1C3 => ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")

	sht.List("List 1").Cell(1, 3).Write(input)
	fmt.Print("R2C3 => ")
	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")

	sht.List("List 1").Cell(2, 3).Write(input)
	fmt.Println("=========================================")
	sht.List("List 1").Print()
}

/*
List 1:
R0C0| R1 | R2 | R3 | R4 |
=========================
 C1 |    |    |    |    |
----+----+----+----+----|
 C2 |    |    |    |    |
----+----+----+----+----|
 C3 |    |    |    |    |
----+----+----+----+----|
 C4 |    |    |    |    |
=========================



*/
