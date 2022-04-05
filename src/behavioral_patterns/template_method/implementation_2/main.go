package main

import "fmt"

func main() {

	highEndComputerBuilder := &HighEndComputerBuilder{
		ComputerBuilder{
			computerParts:          map[string]string{},
			motherboardSetupStatus: []string{},
		},
	}
	highEndComputerBuilder.IComputerBuilder = highEndComputerBuilder

	computer := highEndComputerBuilder.buildComputer()

	fmt.Println(computer.getComputerParts())

}
