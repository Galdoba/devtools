package sheet

type cell struct {
	content string
}

type Cell interface {
	Read() string
	Write(string)
}

func (c *cell) Read() string {
	return c.content
}

func (c *cell) Write(s string) {
	c.content = s
}
