package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings" // Add this import
)

func main() {
	// Step 1: Read the input from stdin (the piped output from one of the quad executables)
	input := new(bytes.Buffer)
	scanner := bufio.NewScanner(os.Stdin)

	var rowCount int
	var columnCount int

	// Step 2: Read the piped output and calculate the dimensions (number of rows and columns)
	for scanner.Scan() {
		line := scanner.Text()
		input.WriteString(line + "\n")

		// Set the column count based on the length of the first line
		if rowCount == 0 {
			columnCount = len(line)
		}
		rowCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Step 3: Print the calculated dimensions
	fmt.Printf("Calculated dimensions: %d x %d\n", columnCount, rowCount)

	// Step 4: Define the quad executables to compare against
	quadExecutables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	matches := []string{} // To store the matches

	// Step 5: Loop through each quad executable and compare outputs
	for _, quad := range quadExecutables {
		// Call the quad executable with the calculated dimensions
		cmd := exec.Command(quad, fmt.Sprintf("%d", columnCount), fmt.Sprintf("%d", rowCount))
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run %s: %v\n", quad, err)
			continue
		}

		// Compare the output of the quad executable with the piped input
		if out.String() == input.String() {
			// Add the match in the required format
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad, columnCount, rowCount))
		}
	}

	// Step 6: Output the matches in the desired format
	if len(matches) > 0 {
		fmt.Println(strings.Join(matches, " || "))
	} else {
		fmt.Println("No match found with any quad.")
	}
}
