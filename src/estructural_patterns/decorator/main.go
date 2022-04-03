package main

import (
	"fmt"
)

func main() {

	christmasTreeWithTinselGarlandBubbleLightsAndTreeTopper := &Tinsel{
		Tree: &Garland{Tree: &BubbleLights{Tree: &TreeTopper{Tree: &ChristmasTreeImpl{}}}},
	}

	fmt.Println(christmasTreeWithTinselGarlandBubbleLightsAndTreeTopper.Decorate())
}
