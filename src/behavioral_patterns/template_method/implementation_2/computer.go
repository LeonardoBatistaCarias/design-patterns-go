package main

type Computer struct {
	computerParts map[string]string
}

func (c *Computer) getComputerParts() map[string]string {
	return c.computerParts
}
