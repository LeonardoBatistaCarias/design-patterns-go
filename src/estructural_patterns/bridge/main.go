package main

import "fmt"

func main() {

	triangle := &Triangle{Shape{&Red{}}}

	fmt.Println(triangle.Draw())

	square := &Square{Shape{&Blue{}}}

	fmt.Println(square.Draw())
}
