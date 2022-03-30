package main

import "fmt"

const (
	DOG = 1
)

type MammalFactory struct{}

func (m *MammalFactory) NewAnimal(mammalType int) (Animal, error) {
	switch mammalType {
	case DOG:
		return new(Dog), nil
	default:
		return nil, fmt.Errorf("Wrong mammal type passed")
	}
}
