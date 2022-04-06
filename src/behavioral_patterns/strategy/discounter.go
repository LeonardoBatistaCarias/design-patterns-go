package main

type Discounter interface {
	ApplyDiscount(amount float32) float32
}
