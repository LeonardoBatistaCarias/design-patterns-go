package main

type IComputerBuilder interface {
	AddMotherboard()
	SetupMotherboard()
	AddProcessor()
	GetComputerParts() map[string]string
}

type ComputerBuilder struct {
	computerParts          map[string]string `default:map[string]string{}`
	motherboardSetupStatus []string          `default:[]string{}`
	IComputerBuilder
}

func (c *ComputerBuilder) buildComputer() Computer {
	c.AddMotherboard()
	c.SetupMotherboard()
	c.AddProcessor()

	return c.getComputer()
}

func (c *ComputerBuilder) GetMotherBoardSetupStatus() []string {
	return c.motherboardSetupStatus
}

func (c *ComputerBuilder) GetComputerParts() map[string]string {
	return c.computerParts
}

func (c *ComputerBuilder) getComputer() Computer {
	return Computer{c.GetComputerParts()}
}
