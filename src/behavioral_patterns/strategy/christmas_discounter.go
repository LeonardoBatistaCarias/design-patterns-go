package main

type ChristmasDiscounter struct {
}

func (c *ChristmasDiscounter) ApplyDiscount(amount float32) float32 {
	return amount * 0.9
}
