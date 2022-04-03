package main

import "fmt"

func main() {
	ferrari := &Ferrari{}
	fmt.Println(ferrari.GetSpeed())

	ferrariMovableAdapter := &MovableAdapterImpl{ferrari}
	fmt.Println(ferrariMovableAdapter.GetSpeed())
}
