package main

import (
	"fmt"
)

const (
	DUCK  = 1
	GOOSE = 2
)

type BirdFactory struct{}

func (b *BirdFactory) NewAnimal(birdType int) (Animal, error) {
	switch birdType {
	case DUCK:
		return new(Duck), nil
	case GOOSE:
		return new(Goose), nil
	default:
		return nil, fmt.Errorf("Wrong bird type passed")
	}
}
