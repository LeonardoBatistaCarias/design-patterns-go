package main

import "fmt"

type ExpensiveObjectImpl struct {
}

func (e *ExpensiveObjectImpl) Process() {
	fmt.Println("Processing complete.")
}

func heavyInitialConfiguration() {
	fmt.Println("Loading initial configuration...")
}

func NewExpensiveObjectImpl() ExpensiveObject {
	heavyInitialConfiguration()

	return &ExpensiveObjectImpl{}
}
