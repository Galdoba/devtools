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

type Border interface {
	Header(List) string
	HorisontalDelimiter(List) string
	Vertical() string
}

func (b *border) SetDefault() {
	b.header = "="
	b.tail = "="
	b.vertical = "|"
	b.horisontal = "-"
	b.crossing = "+"
}

func (b *border) Header(l List) string {
	colData := l.ColWidths()
	hd := ""
	if b.header == "" {
		return hd
	}
	for _, col := range colData {
		hd += b.header
		for i := 0; i < col; i++ {
			hd += b.header
		}
	}
	hd += b.header
	return hd
}

func (b *border) HorisontalDelimiter(l List) string {
	hd := ""
	if b.horisontal == "" && b.crossing == "" {
		return hd
	}
	hor := b.horisontal
	if hor == "" {
		hor = " "
	}
	cros := b.crossing
	if cros == "" {
		cros = " "
	}
	for _, col := range l.ColWidths() {
		hd += cros
		for i := 0; i < col; i++ {
			hd += hor
		}
	}
	hd += cros
	return hd
}

func (b *border) Vertical() string {
	return b.vertical
}
