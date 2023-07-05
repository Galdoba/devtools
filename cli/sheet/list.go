package sheet

import (
	"fmt"
	"strings"
)

type list struct {
	name      string
	maxRow    int
	maxCol    int
	cell      map[coord]Cell
	border    *border
	rowHeight []int
	colWidth  []int
}

func NewList(name string, mR, mC int) *list {
	l := list{}
	l.name = name
	l.maxRow = mR
	l.maxCol = mC
	for len(l.rowHeight) < l.maxRow {
		l.rowHeight = append(l.rowHeight, 1)
	}
	for len(l.colWidth) < l.maxCol {
		l.colWidth = append(l.colWidth, 1)
	}
	l.cell = make(map[coord]Cell)
	l.border = NewBorder()
	l.border.SetDefault()
	for r := 1; r <= l.maxRow; r++ {

		for c := 1; c <= l.maxCol; c++ {
			cell := NewCell(r, c)
			l.cell[NewCoord(r, c)] = &cell

		}
	}
	return &l
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

func (l *list) Print() {
	l.Update()
	vertdel := l.border.vertical
	for r := 1; r <= l.maxRow; r++ {
		row := ""
		for c := 1; c <= l.maxCol; c++ {
			row += vertdel
			crd := NewCoord(r, c)
			data := l.cell[crd].Read()
			for len(data) < l.colWidth[c-1] {
				data += " "
			}
			row += data
			if c == l.maxCol {
				row += vertdel
			}
		}
		fmt.Println(row)
	}
}

func (l *list) Update() {
	for r := 0; r < l.maxRow; r++ {
		for c := 0; c < l.maxCol; c++ {
			data := l.cell[NewCoord(r+1, c+1)].Read()
			if len(data) > l.colWidth[c] {
				l.colWidth[c] = len(data)
			}
			if len(strings.Split(data, "\n")) > l.rowHeight[r] {
				l.rowHeight[r] = len(strings.Split(data, "\n"))
			}
		}
	}
}

//list1{R3:C2}

type Row struct {
	num int
	//height int
}

type Col struct {
	num   int
	width int
}
