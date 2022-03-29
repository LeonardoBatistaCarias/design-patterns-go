package main

import (
	"fmt"
)

func main() {
	triangle, _ := GetPolygon(3)
	square, _ := GetPolygon(4)
	pentagon, _ := GetPolygon(5)

	printDetails(triangle)
	printDetails(square)
	printDetails(pentagon)
}

func printDetails(p Polygon) {
	fmt.Printf("Polygon Form: %s, Sides: %d", p.GetType(), p.GetSides())
	fmt.Println()
}
