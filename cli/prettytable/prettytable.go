package prettytable

import (
	"errors"
	"fmt"
	"time"

	"github.com/Galdoba/TR_Dynasty/cli/features"
)

//fmt.Println("terminal.GetSize(int(os.Stdout.Fd()))")
//fmt.Println(terminal.GetSize(int(os.Stdout.Fd())))

type pTable struct {
	vals   [][]string
	colLen []int
}

type PrettyTable interface {
	PTPrint()
}

func New() pTable {
	return pTable{}
}

func From(data [][]string) pTable {
	pt := pTable{}
	pt.vals = data
	return pt
}

func ColLen(pt pTable) []int {
	var colLen []int
	var l int
	for i := 0; i < TotalCols(pt); i++ {
		colLen = append(colLen, 0)
	}
	for _, row := range pt.vals {
		for c, cell := range row {
			l = len(cell)
			if l > colLen[c] {
				colLen[c] = l
			}
		}
	}
	return colLen
}

func TotalCols(pt pTable) int {
	if len(pt.vals) > 0 {
		return len(pt.vals[0])
	}
	return 0
}

func (pt pTable) PTPrint() {
	tbl := ""
	colLen := ColLen(pt)
	for _, row := range pt.vals {
		tbl += "|"
		for c, cell := range row {
			tbl += cell
			filled := len(cell)
			for filled < colLen[c] {
				tbl += " "
				filled++
			}
			//if c != len(row)-1 {
			tbl += " | "
			//}
		}
		tbl += "\n"
	}
	fmt.Print(tbl)
}

func (pt pTable) PTPrintSlow(msDelay time.Duration) {
	tbl := ""
	colLen := ColLen(pt)
	for _, row := range pt.vals {
		tbl += "|"
		for c, cell := range row {
			tbl += cell
			filled := len(cell)
			for filled < colLen[c] {
				tbl += " "
				filled++
			}
			//if c != len(row)-1 {
			tbl += " | "
			//}
		}
		tbl += "\n"
	}
	//fmt.Print(tbl)
	features.TypingSlowly(tbl, msDelay)
}

func InsertSeparatorRow(pt pTable, r int) pTable {
	newPt := pTable{}
	newPt.colLen = pt.colLen
	for i, oldRow := range pt.vals {
		if i == r {
			newPt.vals = append(newPt.vals, separatorFor(pt))
		}
		newPt.vals = append(newPt.vals, oldRow)
	}
	return newPt
}

func separatorFor(pt pTable) []string {
	var row []string
	colLen := ColLen(pt)
	for i := range colLen {
		cell := ""
		for len(cell) < colLen[i] {
			cell += "-"
		}
		row = append(row, cell)
	}

	return row
}

func (pt *pTable) AddRow(row []string) {
	pt.vals = append(pt.vals, row)
	return
}

func (pt pTable) Value(r, c int) (string, error) {
	if r >= len(pt.vals) || len(pt.vals) == 0 {
		return "", errors.New("Row number out of bounds")
	}
	if c >= len(pt.vals[0]) {
		return "", errors.New("Column number out of bounds")
	}
	return pt.vals[r][c], nil
}
