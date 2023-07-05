package sheet

type border struct {
	header     string
	tail       string
	vertical   string
	horisontal string
	crossing   string
}

func NewBorder() *border {
	brd := border{}
	brd.header = "="
	brd.tail = "="
	brd.vertical = "|"
	brd.horisontal = "-"
	brd.crossing = "+"
	return &brd
}

func (b *border) SetDefault() {
	b.header = "="
	b.tail = "="
	b.vertical = "|"
	b.horisontal = "-"
	b.crossing = "+"
}
