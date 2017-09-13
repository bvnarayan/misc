package battleship

import (
	"testing"
)

const (
	SIZEOFGRID = 5
)

// get the 5 coordinates that represents the battleship location and
// print the 5x5 grid

func TestAll(t *testing.T) {

	p1 := make([][]int, SIZEOFGRID)

	for n := 0; n < SIZEOFGRID; n++ {
		p1[n] = make([]int, SIZEOFGRID)
	}

	PopulateGrid(p1, -1, -1, -1, -1, -1)
	PrintGrid(p1)

	numGridsOccupied := NumberOfGridsOccupied(p1)

	if numGridsOccupied != 0 {
		t.Fatalf("Number of grids occupied is not 0")
	}

	status := IsShipSunk(p1, 0)

	if status != true {
		t.Fatalf("P1 has unexpected ships")
	}
}
