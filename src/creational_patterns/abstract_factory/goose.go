package main

type Goose struct {
}

func (g *Goose) GetAnimalName() string {
	return "Goose"
}

func (g *Goose) MakeSound() string {
	return "Squack"
}
