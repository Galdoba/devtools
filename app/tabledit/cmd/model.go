package cmd

import (
	"github.com/Galdoba/devtools/app/tabledit/pkg/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	liptab "github.com/charmbracelet/lipgloss/table"
)

type model struct {
	table table.Table
	ltab  *liptab.Table
	mode  string
}

type cursor struct {
	row int
	col int
}

func NewModel() (model, error) {
	m := model{}
	tab, err := table.ImportCSV(cfg.DefaultTablePath())
	if err != nil {
		return m, err
	}
	m.table = tab

	m.ltab = liptab.New()
	m.ltab.Headers([]string{"11", "11", "11", "11", "11"}...)
	return m, err
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "m", "ь":
			m = m.toggleMode()
			return m, nil
		}

	}

	return m, nil
}

var styleHeader = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("66")).
	BorderBackground(lipgloss.Color("15"))

var styleCell = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("80")).
	BorderBackground(lipgloss.Color("26"))

func (m model) View() string {
	if m.table == nil {
		return "no table"
	}

	header := lipgloss.JoinHorizontal(lipgloss.Left, "file: "+cfg.DefaultTablePath())

	s := header + "\n"
	m.ltab.ClearRows()

	strDat := liptab.NewStringData([]string{"Source ", "TYPE ", "Description ", "OMIT ", "Comment "})
	strDat.Append([]string{"Source ", "TYPE ", "Description ", "OMIT ", "Comment "})
	for _, data := range m.table.Cells() {
		strDat.Append(data)
	}
	m.ltab.Data(strDat)
	//m.ltab.Width(totalW(m.table))

	s += m.ltab.Render() + "\n"
	//s += m.table.String() + "\n"
	s += "MODE: " + m.mode + "\n"

	//s += lipgloss.JoinHorizontal(lipgloss.Left, "111\n2 2 ", "3333\n44", "#55")

	return s
}

func totalW(t table.Table) int {
	w := 1
	for _, wc := range t.Widths() {
		w += wc + 3
	}
	return w
}

func (m model) toggleMode() model {
	switch m.mode {
	case "EDIT":
		m.mode = "SELECTION"
	default:
		m.mode = "EDIT"
	}
	return m
}

/*
s file: [path/to/file]
s +-------------------+-----+---+
s |coment             | ID  | № |
s +-------------------+-----+---+
d | data 1            |     |   |   OR  d | no data           |     |   |
d | data 2            |     |   |
d | data 3            |     |   |
d | data 4            |     |   |
s +-------------------+-----+---+
s MODE: mode

*/
