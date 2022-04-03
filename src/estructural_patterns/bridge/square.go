package main

type Square struct {
	Shape
}

func (s *Square) Draw() string {
	return "Square draw. " + s.Color.Fill()
}
