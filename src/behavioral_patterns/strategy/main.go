package main

import "fmt"

func main() {
	orderWithEasterDiscounter := &Order{1000, &EasterDiscounter{}}
	fmt.Printf("The total amount: $%f ", orderWithEasterDiscounter.ApplyDiscount(orderWithEasterDiscounter.Amount))

	fmt.Println()

	orderWithChristmasDiscounter := &Order{1000, &ChristmasDiscounter{}}
	fmt.Printf("The total amount: $%f ", orderWithChristmasDiscounter.ApplyDiscount(orderWithChristmasDiscounter.Amount))

}
