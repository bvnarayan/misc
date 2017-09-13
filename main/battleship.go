package main

import (
	"fmt"
	"strconv"

        "github.com/bvnarayan/misc/battleship"
)

const (
// SIZEOFGRID IS the grid size that the game is played on
	SIZEOFGRID = 5
)

// get the 5 coordinates that represents the battleship location and
// print the 5x5 grid

func main() {
	// 5 integers that represent coordinate of ship
	var i, j, k, l, m int
	// coordinates that players use for attacking
	var x, y int

	// this map represents coordinates that each player attacks
	p1Attacks := make(map[string]int)
	p2Attacks := make(map[string]int)

	p1 := make([][]int, SIZEOFGRID)
	p2 := make([][]int, SIZEOFGRID)

	for n := 0; n < SIZEOFGRID; n++ {
		p1[n] = make([]int, SIZEOFGRID)
		p2[n] = make([]int, SIZEOFGRID)
	}

	fmt.Println("Enter Player1 (P1) ship coordinates")
	i, j, k, l, m = battleship.Read5Ints(SIZEOFGRID)
	battleship.PopulateGrid(p1, i, j, k, l, m)

	fmt.Println("Enter Player2 (P2) ship coordinates")
	i, j, k, l, m = battleship.Read5Ints(SIZEOFGRID)
	battleship.PopulateGrid(p2, i, j, k, l, m)

	fmt.Println("Player 1 grid")
	battleship.PrintGrid(p1)

	fmt.Println("Player 2 grid")
	battleship.PrintGrid(p2)

	fmt.Printf("Number of grids occupied for P1 %d\n", battleship.NumberOfGridsOccupied(p1))
	fmt.Printf("Number of grids occupied for P2 %d\n", battleship.NumberOfGridsOccupied(p2))

	state := "START"

	for state != "END" {
		fmt.Println("P1 enter coordinates you want to attack")
		//	fmt.Scanf("%d %d\n", &x, &y)
		x, y = battleship.Read2Ints(SIZEOFGRID)

		coordinates := "x:" + strconv.Itoa(x) + "y:" + strconv.Itoa(y)
		if _, ok := p1Attacks[coordinates]; ok {
			fmt.Println("Already Taken")
		} else {
			p1Attacks[coordinates] = 1
			if p2[x][y] == 1 {
				fmt.Println("Hit!!!")
				p2[x][y] = 0
				if battleship.IsShipSunk(p2, y) {
					fmt.Println("Sunk a ship!!!")
				}
			} else {
				fmt.Println("Miss!!!")
			}
		}
		if 0 == battleship.NumberOfGridsOccupied(p2) {
			fmt.Println("P1 wins")
			state = "END"
			break
		}

		fmt.Println("P2 enter coordinates you want to attack")
		x, y = battleship.Read2Ints(SIZEOFGRID)
		//fmt.Scanf("%d %d\n", &x, &y)
		coordinates = "x:" + strconv.Itoa(x) + "y:" + strconv.Itoa(y)
		if _, ok := p2Attacks[coordinates]; ok {
			fmt.Println("Already Taken")
		} else {
			p2Attacks[coordinates] = 1
			if p1[x][y] == 1 {
				fmt.Println("Hit!!!")
				p1[x][y] = 0
				if battleship.IsShipSunk(p1, y) {
					fmt.Println("Sunk a ship!!!")
				}
			} else {
				fmt.Println("Miss!!!")
			}
		}
		if 0 == battleship.NumberOfGridsOccupied(p1) {
			fmt.Println("P2 wins")
			state = "END"
		}
	}
}
