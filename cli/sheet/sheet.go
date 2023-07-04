package sheet

type sheet struct {
}

type list struct {
	maxRow int
	maxCol int
	cell   map[coord]Cell
}

func (l *list) Cell(r, c int) Cell {
	return l.cell[*NewCoord(r, c)]
}

//list1{R3:C2}
