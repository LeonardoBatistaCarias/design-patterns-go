package main

type Triangle struct {
	Shape
}

func (t *Triangle) Draw() string {
	return "Triangle draw. " + t.Color.Fill()
}
