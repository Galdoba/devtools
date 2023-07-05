package sheet

const (
	Mode_View              = 0
	Mode_NavigateByCELLS   = 1
	Mode_NavigateByROWS    = 2
	Mode_NavigateByCOLOMNS = 3
)

type sheet struct {
	lists      map[string]*list
	activeList string
	buffer     string
	mode       int
}

func NewSheet(lists ...*list) *sheet {
	sh := sheet{}
	sh.lists = make(map[string]*list)
	for _, list := range lists {
		sh.lists[list.Name()] = list

	}
	return &sh
}

func (sh *sheet) List(name string) *list {
	return sh.lists[name]
}
