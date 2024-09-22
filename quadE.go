package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if enough arguments are provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./quadE x y")
		return
	}

	// Convert the arguments to integers
	x, errX := strconv.Atoi(os.Args[1])
	y, errY := strconv.Atoi(os.Args[2])

	// Check if the arguments are valid integers
	if errX != nil || errY != nil || x <= 0 || y <= 0 {
		fmt.Println("Error: Invalid input. Both x and y must be positive integers.")
		return
	}

	// Set the dimensions as an environment variable
	os.Setenv("QUAD_DIMENSIONS", fmt.Sprintf("%d %d", x, y))

	// Loop through the rows (y times)
	for i := 0; i < y; i++ {
		// Loop through the columns (x times)
		for j := 0; j < x; j++ {
			// First row
			if i == 0 {
				if j == 0 {
					fmt.Print("A") // top-left corner
				} else if j == x-1 {
					fmt.Print("C") // top-right corner
				} else {
					fmt.Print("B") // top edge
				}
			} else if i == y-1 {
				// Last row
				if j == 0 {
					fmt.Print("C") // bottom-left corner
				} else if j == x-1 {
					fmt.Print("A") // bottom-right corner
				} else {
					fmt.Print("B") // bottom edge
				}
			} else {
				// Middle rows
				if j == 0 || j == x-1 {
					fmt.Print("B") // left and right edges
				} else {
					fmt.Print(" ") // interior
				}
			}
		}
		fmt.Println() // New line at the end of each row
	}
}
