package main

type Square struct {
	Name string
	Side int
}

func (s *Square) GetType() string {
	return s.Name
}

func (s *Square) GetSides() int {
	return s.Side
}

func NewSquare() Polygon {
	return &Square{
		Name: "Square",
		Side: 4,
	}
}
