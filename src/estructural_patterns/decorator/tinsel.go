package main

type Tinsel struct {
	Tree ChristmasTree
}

func (t *Tinsel) Decorate() string {
	return t.Tree.Decorate() + decorateWithTinsel()
}

func decorateWithTinsel() string {
	return " with Tinsel"
}
