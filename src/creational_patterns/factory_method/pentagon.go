package main

type Pentagon struct {
	Name string
	Side int
}

func (p *Pentagon) GetType() string {
	return p.Name
}

func (p *Pentagon) GetSides() int {
	return p.Side
}

func NewPentagon() Polygon {
	return &Pentagon{
		Name: "Pentagon",
		Side: 5,
	}
}
