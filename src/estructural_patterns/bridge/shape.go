package main

type Shape struct {
	Color Color
}

func (s *Shape) Draw() string {
	return "Square draw. " + s.Color.Fill()
}
