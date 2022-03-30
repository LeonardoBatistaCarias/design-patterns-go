package main

type Duck struct {
}

func (d *Duck) GetAnimalName() string {
	return "Duck"
}

func (d *Duck) MakeSound() string {
	return "Quack"
}
