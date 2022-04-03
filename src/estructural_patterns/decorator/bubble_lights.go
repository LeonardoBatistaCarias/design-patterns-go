package main

type BubbleLights struct {
	Tree ChristmasTree
}

func (b *BubbleLights) Decorate() string {
	return b.Tree.Decorate() + decorateWithBubbleLights()
}

func decorateWithBubbleLights() string {
	return " with BubbleLights"
}
