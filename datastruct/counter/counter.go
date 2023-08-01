package counter

type counter struct {
	val  int
	name string
}

func New(name string) *counter {
	c := counter{}
	c.name = name
	return &c
}

type Counter interface {
	Name() string
	Increase()
	IncreaseBy(int)
	Decrease()
	DecreaseBy(int)
	Value() int
}

func (c *counter) Name() string {
	return c.name
}

func (c *counter) Increase() {
	c.val++
}

func (c *counter) IncreaseBy(n int) {
	c.val = c.val + n
}

func (c *counter) Decrease() {
	c.val--
}

func (c *counter) DecreaseBy(n int) {
	c.val = c.val - n
}

func (c *counter) Value() int {
	return c.val
}

func (c *counter) Reset() {
	c.val = 0
}
