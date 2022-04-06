package main

type EasterDiscounter struct {
}

func (e *EasterDiscounter) ApplyDiscount(amount float32) float32 {
	return amount * 0.5
}
