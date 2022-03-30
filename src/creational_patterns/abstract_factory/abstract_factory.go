package main

import (
	"fmt"
)

const (
	BIRD   = 1
	MAMMAL = 2
)

func CreateAnimalFactory(animalFactoryType int) (AnimalFactory, error) {
	switch animalFactoryType {
	case BIRD:
		return new(BirdFactory), nil
	case MAMMAL:
		return new(MammalFactory), nil
	default:
		return nil, fmt.Errorf("Wrong factory type passed")
	}
}
