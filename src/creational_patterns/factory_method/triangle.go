package main

type Triangle struct {
	Name string
	Side int
}

func (t *Triangle) GetType() string {
	return t.Name
}

func (t *Triangle) GetSides() int {
	return t.Side
}

func NewTriangle() Polygon {
	return &Triangle{
		Name: "Triangle",
		Side: 3,
	}
}
