package sheet

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/muesli/termenv"
)

type list struct {
	name      string
	maxRow    int
	maxCol    int
	cell      map[coord]Cell
	border    *border
	rowHeight []int
	colWidth  []int
	writters  []io.Writer
}

func NewList(name string, mR, mC int) *list {
	l := list{}
	l.name = name
	l.maxRow = mR
	l.maxCol = mC
	l.writters = append(l.writters, os.Stdout)
	for len(l.rowHeight) < l.maxRow {
		l.rowHeight = append(l.rowHeight, 1)
	}
	for len(l.colWidth) < l.maxCol {
		l.colWidth = append(l.colWidth, 1)
	}
	l.cell = make(map[coord]Cell)
	l.border = NewBorder()
	l.border.SetDefault()
	for r := 0; r < l.maxRow; r++ {

		for c := 0; c < l.maxCol; c++ {
			cell := NewCell(r, c)
			l.cell[NewCoord(r, c)] = &cell

		}
	}
	return &l
}

type List interface {
	Cell(int, int) Cell
	CellsMatrix() [][]Cell
	Update()
	Border() Border
	ColWidths() []int
	Name() string
	Size() (int, int)
}

func (l *list) Cell(r, c int) Cell {
	coord := NewCoord(r, c)
	if !coord.Valid(r, c) {
		return nil
	}
	return l.cell[coord]
}

func (l *list) Name() string {
	return l.name
}

func (l *list) CellsMatrix() [][]Cell {
	cells := [][]Cell{}
	for r := 0; r < l.maxRow; r++ {
		row := []Cell{}
		for c := 0; c < l.maxCol; c++ {
			row = append(row, l.Cell(r, c))
		}
		cells = append(cells, row)
	}
	return cells
}

func Print(l List) {
	restoreConsole, err := termenv.EnableVirtualTerminalProcessing(termenv.DefaultOutput())
	if err != nil {
		panic(err)
	}
	defer restoreConsole()
	//p := termenv.ColorProfile()

	l.Update()
	vertdel := l.Border().Vertical()
	if vertdel == "" {
		vertdel = " "
	}

	for _, row := range l.CellsMatrix() {
		for c, cell := range row {
			text := AlignText(cell)
			for len(text) < l.ColWidths()[c] {
				text += " "
			}
			switch cell.Cursor() {
			case true:
				fmt.Printf("%s ", termenv.String(text).Foreground(termenv.ANSIBlack).Background(termenv.ANSIWhite))
			case false:
				fmt.Printf("%s ", termenv.String(text).Foreground(termenv.ANSIWhite).Background(termenv.ANSIBlack))
			}

			//fmt.Print(text, " ")
		}
		fmt.Print("\n")
	}
}

func (l *list) Update() {
	for r := 0; r < l.maxRow; r++ {
		for c := 0; c < l.maxCol; c++ {
			data := l.cell[NewCoord(r, c)].Read()
			if len(data) > l.colWidth[c] {
				l.colWidth[c] = len(data)
			}
			if len(strings.Split(data, "\n")) > l.rowHeight[r] {
				l.rowHeight[r] = len(strings.Split(data, "\n"))
			}
		}
	}
}

func (l *list) ColWidths() []int {
	return l.colWidth
}

func (l *list) Border() Border {
	return l.border
}

func (l *list) Size() (int, int) {
	return l.maxCol, l.maxRow
}

//list1{R3:C2}

/*
[fg:red;bg:black]
*/

type Row struct {
	num int
	//height int
}

type Col struct {
	num   int
	width int
}

func DebugFill(l List) {
	fl := 0
	width, height := l.Size()
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if l.Cell(r, c) == nil {
				continue
			}
			l.Cell(r, c).Write(fmt.Sprintf("filler %v", fl))
			fl++
		}
	}
}
