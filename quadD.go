package piscine

import "fmt"

func QuadD(x, y int) {
	// Check if x and y are positive, otherwise do nothing
	if x <= 0 || y <= 0 {
		return
	}

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
					fmt.Print("A") // bottom-left corner
				} else if j == x-1 {
					fmt.Print("C") // bottom-right corner
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
