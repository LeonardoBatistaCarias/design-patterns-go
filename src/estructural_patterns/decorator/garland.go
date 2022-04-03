package main

type Garland struct {
	Tree ChristmasTree
}

func (g *Garland) Decorate() string {
	return g.Tree.Decorate() + decorateWithGarland()
}

func decorateWithGarland() string {
	return " with Garland"
}
