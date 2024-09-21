package main

import "fmt"

func QuadA(x, y int) {
	// Check if x and y are positive, otherwise do nothing
	if x <= 0 || y <= 0 {
		return
	}

	// Loop through the rows (y times)
	for i := 0; i < y; i++ {
		// Loop through the columns (x times)
		for j := 0; j < x; j++ {
			// First row and last row
			if i == 0 || i == y-1 {
				if j == 0 || j == x-1 {
					fmt.Print("o") // corners
				} else {
					fmt.Print("-") // top and bottom edges
				}
			} else {
				if j == 0 || j == x-1 {
					fmt.Print("|") // left and right edges
				} else {
					fmt.Print(" ") // interior
				}
			}
		}
		fmt.Println() // New line at the end of each row
	}
}
