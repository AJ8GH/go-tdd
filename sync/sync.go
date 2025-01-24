package sync

type Counter struct {
	val int
}

func (c *Counter) Inc() {
	c.val++
}

func (c *Counter) Value() int {
	return c.val
}
