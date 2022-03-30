package main

import "fmt"

func main() {
	birdFactory, _ := CreateAnimalFactory(BIRD)

	duck, _ := birdFactory.NewAnimal(DUCK)
	goose, _ := birdFactory.NewAnimal(GOOSE)
	printDetails(duck)
	printDetails(goose)

	mammalFactory, _ := CreateAnimalFactory(MAMMAL)

	dog, _ := mammalFactory.NewAnimal(DOG)
	printDetails(dog)

}

func printDetails(a Animal) {
	fmt.Printf("Animal: %s, Sound: %s", a.GetAnimalName(), a.MakeSound())
	fmt.Println()
}
