package sheet

const (
	Mode_View              = 0
	Mode_NavigateByCELLS   = 1
	Mode_NavigateByROWS    = 2
	Mode_NavigateByCOLOMNS = 3
)

type sheet struct {
	lists      map[string]List
	activeList string
	buffer     string
	mode       int
}

func NewSheet(lists ...List) *sheet {
	sh := sheet{}
	sh.lists = make(map[string]List)
	for _, list := range lists {
		sh.lists[list.Name()] = list

	}
	return &sh
}

func (sh *sheet) List(name string) List {
	return sh.lists[name]
}
