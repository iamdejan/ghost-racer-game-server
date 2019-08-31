package modelimpl

type circuit struct {
	circuitID uint64
}

func (c *circuit) ID() uint64 {
	return c.circuitID
}