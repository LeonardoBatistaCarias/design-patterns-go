package main

type Dog struct {
}

func (d *Dog) GetAnimalName() string {
	return "Dog"
}

func (d *Dog) MakeSound() string {
	return "Au au"
}
