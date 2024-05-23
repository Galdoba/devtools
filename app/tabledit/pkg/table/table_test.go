package table

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestTable(t *testing.T) {
	tab := New()
	tab.SetSize(3, 3)
	tab.SetContent(0, 1, "111")
	tab.SetContent(1, 1, "2222")
	tab.SetContent(1, 0, "24")
	tab.SetContent(2, 2, "33")
	fmt.Println(tab.String())
	fmt.Println(tab.Widths())
	fmt.Println(tab.Cell(2, 2).String())
	wd := tab.Widths()

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("114")).
		Align(lipgloss.Center).
		Border(lipgloss.DoubleBorder(), true).
		BorderForeground(lipgloss.Color("42"))

	fmt.Println(style.Render("Hello, kitty"))
	allCells := tab.Cells()
	line := ""
	for i, cell := range allCells[0] {
		text := cell
		for len(strings.Split(text, "")) < wd[i] {
			text += " "
		}
		line += text + style.GetBorderStyle().Left
	}
	fmt.Println(style.Render(line))
}
