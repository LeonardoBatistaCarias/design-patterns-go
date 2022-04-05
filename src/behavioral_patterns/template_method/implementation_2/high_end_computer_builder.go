package main

import "fmt"

type HighEndComputerBuilder struct {
	ComputerBuilder
}

func (h *HighEndComputerBuilder) AddMotherboard() {
	h.computerParts["Motherboard"] = "High-end Motherboard"
}

func (h *HighEndComputerBuilder) SetupMotherboard() {
	h.motherboardSetupStatus = append(h.motherboardSetupStatus, "Screwing the high-end motherboard to the case.")
	h.motherboardSetupStatus = append(h.motherboardSetupStatus, "Pluging in the power supply connectors.")

	for _, motherboardPart := range h.motherboardSetupStatus {
		fmt.Println(motherboardPart)
	}
}

func (h *HighEndComputerBuilder) AddProcessor() {
	h.computerParts["Processor"] = "High-end Processor"
}

func (h *HighEndComputerBuilder) GetComputerParts() map[string]string {
	return h.computerParts
}
