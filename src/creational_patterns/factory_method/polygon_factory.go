package main

import (
	"fmt"
)

func GetPolygon(numberOfSides int) (Polygon, error) {

	switch numberOfSides {
	case 3:
		return NewTriangle(), nil
	case 4:
		return NewSquare(), nil
	case 5:
		return NewPentagon(), nil
	}

	return nil, fmt.Errorf("Wrong polygon type passed")
}
