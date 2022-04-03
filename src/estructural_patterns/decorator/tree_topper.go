package main

type TreeTopper struct {
	Tree ChristmasTree
}

func (t *TreeTopper) Decorate() string {
	return t.Tree.Decorate() + decorateWithTreeTopper()
}

func decorateWithTreeTopper() string {
	return " with TreeTopper"
}
