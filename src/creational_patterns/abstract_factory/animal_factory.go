package main

type AnimalFactory interface {
	NewAnimal(int) (Animal, error)
}
