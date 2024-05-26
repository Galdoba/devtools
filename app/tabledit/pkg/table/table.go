package table

import (
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/devtools/csvp"
)

type table struct {
	cols         int
	rows         int
	initComplete bool
	borders      string
	Data         [][]*Cell
}

type Table interface {
	SetSize(int, int) error
	AddRow()
	AddCol()
	SetContent(int, int, string) error
	Cell(int, int) *Cell
	Widths() []int
	String() string
	Cells() [][]string
	At(int, int) string
	Rows() int
	Columns() int
}

func New() *table {
	t := table{}
	t.borders = "|"
	return &t
}

func (t *table) AddRow() {
	t.Data = append(t.Data, []*Cell{})
}

func (t *table) AddCol() {
	for i, row := range t.Data {
		row = append(row, &Cell{})
		t.Data[i] = row
	}
}

func (t *table) SetSize(col, row int) error {
	if t.initComplete {
		return fmt.Errorf("can't set size after initiation")
	}
	if col < 0 {
		return fmt.Errorf("can't set columns less than zero")
	}
	if row < 0 {
		return fmt.Errorf("can't set rows less than zero")
	}
	for len(t.Data) < row {
		newRow := []*Cell{}
		for len(newRow) < col {
			newRow = append(newRow, &Cell{})
		}
		t.Data = append(t.Data, newRow)
	}
	t.initComplete = true
	return nil
}

func (t *table) SetContent(row, col int, cont string) error {
	if col > len(t.Data[0]) {
		fmt.Println("err 1")
		return fmt.Errorf("bad col '%v' for table (cols:%v rows:%v)", col, len(t.Data[0]), len(t.Data))
	}
	if col < 0 {
		fmt.Println("err 2")
		return fmt.Errorf("bad col '%v' for table (cols:%v rows:%v)", col, len(t.Data[0]), len(t.Data))
	}
	if row > len(t.Data)-1 {
		fmt.Println("err 3")
		return fmt.Errorf("bad row '%v' for table (cols:%v rows:%v)", row, len(t.Data[0]), len(t.Data))
	}
	if row < 0 {
		fmt.Println("err 4")
		return fmt.Errorf("bad row '%v' for table (cols:%v rows:%v)", row, len(t.Data[0]), len(t.Data))
	}

	t.Data[row][col].Content = cont
	return nil
}

func (t *table) Widths() []int {
	maxColWidth := []int{}
	for _, cell := range t.Data[0] {
		maxColWidth = append(maxColWidth, len(strings.Split(cell.Content, "")))
	}
	for _, row := range t.Data {
		for c, cell := range row {
			colWidth := len(strings.Split(cell.Content, ""))
			if colWidth > maxColWidth[c] {
				maxColWidth[c] = colWidth
			}
		}
	}
	return maxColWidth
}

func (t *table) Cell(col, row int) *Cell {
	badCell := &Cell{}
	badCell.err = fmt.Errorf("bad cell position [cols:%v rows:%v] (cols:%v rows:%v)", col, row, len(t.Data[0]), len(t.Data))
	if row > len(t.Data)-1 {
		return badCell
	}
	if row < 0 {
		return badCell
	}
	if col > len(t.Data[0])-1 {
		return badCell
	}
	if col < 0 {
		return badCell
	}

	return t.Data[row][col]

}
func (t *table) String() string {
	s := ""
	maxColWidth := t.Widths()
	for _, row := range t.Data {
		for c, cell := range row {
			text := cell.Content
			for len(strings.Split(text, "")) < maxColWidth[c] {
				text += " "
			}
			s += t.borders + text
		}
		s += t.borders + "\n"
	}
	return strings.TrimSuffix(s, "\n")
}

func (t *table) Cells() [][]string {
	cls := [][]string{}
	for _, row := range t.Data {
		rw := []string{}
		for _, cl := range row {
			rw = append(rw, cl.Content)
		}
		cls = append(cls, rw)
	}
	return cls
}

func (t *table) At(row, col int) string {
	return t.Data[row][col].Content
}

func (t *table) Rows() int {
	return len(t.Data)
}

func (t *table) Columns() int {
	return len(t.Widths())
}

//////////////////////

func ImportCSV(path string) (*table, error) {
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("bad path: %v", err.Error())
	}
	text := string(bt)
	e, err := csvp.FromString(text)
	if err != nil {
		return nil, fmt.Errorf("csv conversion failed: %v", err.Error())
	}
	entries := e.Entries()
	rows := len(entries)
	cols := len(entries[0].Fields())
	tab := New()
	if err := tab.SetSize(cols, rows); err != nil {
		return nil, fmt.Errorf("can't set size: %v", err.Error())
	}
	for r, entry := range entries {
		for c, field := range entry.Fields() {
			if err := tab.SetContent(r, c, field); err != nil {
				return nil, fmt.Errorf("can't set cell content: %v", err.Error())
			}
		}
	}
	return tab, nil
}
