package battleship

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Assumption: ship is 3 units long and 1 unit wide

// PopulateGrid takes a 2 dimensional slice as input and populates
// the 2D slice with the parameters passed to it 
// currently only populates a 5x5 grid
func PopulateGrid(grid [][]int, i, j, k, l, m int) {

	switch i {
	case -1:
		// do nothing here.
	case 0:
		grid[0][0] = 1
		grid[1][0] = 1
		grid[2][0] = 1
	case 1:
		grid[1][0] = 1
		grid[2][0] = 1
		grid[3][0] = 1
	case 2:
		grid[2][0] = 1
		grid[3][0] = 1
		grid[4][0] = 1
	default:
		// do nothing here.
	}

	switch j {
	case -1:
		// do nothing here.
	case 0:
		grid[0][1] = 1
		grid[1][1] = 1
		grid[2][1] = 1

	case 1:
		grid[1][1] = 1
		grid[2][1] = 1
		grid[3][1] = 1
	case 2:
		grid[2][1] = 1
		grid[3][1] = 1
		grid[4][1] = 1
	default:
		// do nothing here.
	}

	switch k {
	case -1:
	// do nothing here.

	case 0:
		grid[0][2] = 1
		grid[1][2] = 1
		grid[2][2] = 1
	case 1:
		grid[1][2] = 1
		grid[2][2] = 1
		grid[3][2] = 1
	case 2:
		grid[2][2] = 1
		grid[3][2] = 1
		grid[4][2] = 1
	default:
		// do nothing here.
	}

	switch l {
	case -1:
		// do nothing here.
	case 0:
		grid[0][3] = 1
		grid[1][3] = 1
		grid[2][3] = 1
	case 1:
		grid[1][3] = 1
		grid[2][3] = 1
		grid[3][3] = 1
	case 2:
		grid[2][3] = 1
		grid[3][3] = 1
		grid[4][3] = 1
	default:
		// do nothing here.
	}

	switch m {
	case -1:
		// do nothing here.
	case 0:
		grid[0][4] = 1
		grid[1][4] = 1
		grid[2][4] = 1

	case 1:
		grid[1][4] = 1
		grid[2][4] = 1
		grid[3][4] = 1
	case 2:
		grid[2][4] = 1
		grid[3][4] = 1
		grid[4][4] = 1
	default:
		// do nothing here.
	}

}

// IsShipSunk  returns true if ship is sunk and false otherwise
// 2D slice and the column in which to check is passed as input
// currently works only for a 5x5 grid
func IsShipSunk(grid [][]int, column int) bool {
	return (grid[0][column] == 0 &&
		grid[1][column] == 0 &&
		grid[2][column] == 0 &&
		grid[3][column] == 0 &&
		grid[4][column] == 0)

}

// NumberOfGridsOccupied returns the number of grids that are non empty
// A value of 0 would indicate that all ships in grid have been totally 
// destroyed 
func NumberOfGridsOccupied(grid [][]int) (numgrids int) {
        sizeofgrid := len(grid)
	for i := 0; i < sizeofgrid; i++ {
		for j := 0; j < sizeofgrid; j++ {
			if grid[i][j] == 1 {
				numgrids++
			}
		}
	}
	return
}

// PrintGrid prints the 2D slice passed to it
func PrintGrid(grid [][]int) {
        sizeofgrid := len(grid)
	fmt.Println("+--------------------+")
	for i := 0; i < sizeofgrid; i++ {
		for j := 0; j < sizeofgrid; j++ {
			fmt.Print("| ", grid[i][j], " ")
		}
		fmt.Println(" |")
	}
	fmt.Println("+--------------------+")
	fmt.Println()
}

// Read2Ints reads 2 integers from the command line 
func Read2Ints(sizeofgrid int64) (int, int) {

	reader := bufio.NewReader(os.Stdin)

	var x, y int64
	var text string
	var err error

	valid := false
	for valid == false {
		fmt.Print("Enter 2 integers separated by space: ")
		text, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		nums := strings.Fields(text)
		if len(nums) == 2 {

			x, err = strconv.ParseInt(nums[0], 0, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			y, err = strconv.ParseInt(nums[1], 0, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			// coordinates cannot be negative
			if x < 0 || y < 0 ||
				x > sizeofgrid-1 ||
				y > sizeofgrid-1 {
				fmt.Println("Invalid input")
			} else {
				fmt.Println("Valid Input")
				valid = true
			}

		} else {
			fmt.Println("Invalid input")
		}
	}
	// ParseInt returns only int64. Convert int64 to int
	return int(x), int(y)
}

// Read5Ints reads 5 integers from the command line 
func Read5Ints(sizeofgrid int64) (int, int, int, int, int) {

	reader := bufio.NewReader(os.Stdin)

	var i, j, k, l, m int64
	var text string
	var err error

	valid := false
	for valid == false {
		fmt.Print("Enter 5 integers separated by space: ")
		text, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		nums := strings.Fields(text)
		if len(nums) == 5 {

			i, err = strconv.ParseInt(nums[0], 0, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			j, err = strconv.ParseInt(nums[1], 0, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			k, err = strconv.ParseInt(nums[2], 0, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			l, err = strconv.ParseInt(nums[3], 0, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			m, err = strconv.ParseInt(nums[4], 0, 64)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			// ship has a size of 3 and has to be contained within the grid
			if i < -1 || j < -1 || k < -1 || l < -1 || m < -1 ||
				i > sizeofgrid-3 || j > sizeofgrid-3 || k > sizeofgrid-3 ||
				l > sizeofgrid-3 || m > sizeofgrid-3 {
				fmt.Println("Invalid input")
			} else {
				fmt.Println("Valid Input")
				valid = true
			}

		} else {
			fmt.Println("Invalid input")
		}
	}
	// ParseInt returns only int64. Convert int64 to int
	return int(i), int(j), int(k), int(l), int(m)
}
