package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

	// Step 3: Check if the output contains an error message starting with "Usage:" or invalid input
	if columnCount == 0 || rowCount == 0 || input.Len() == 0 || strings.HasPrefix(input.String(), "Usage:") {
		fmt.Fprintln(os.Stderr, "Error: Invalid input or arguments. Ensure that the quad function is used with proper dimensions.")
		os.Exit(1)
	}

	// Step 4: Print the calculated dimensions
	fmt.Printf("Calculated dimensions: %d x %d\n", columnCount, rowCount)

	// Step 5: Define the quad executables to compare against
	quadExecutables := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	// Step 6: Adjust for Windows (add .exe extension)
	if runtime.GOOS == "windows" {
		for i := range quadExecutables {
			quadExecutables[i] += ".exe"
		}
	}

	matches := []string{} // To store the matches

	// Step 7: Loop through each quad executable and compare outputs
	for _, quad := range quadExecutables {
		// Call the quad executable with the calculated dimensions
		cmd := exec.Command("./"+quad, fmt.Sprintf("%d", columnCount), fmt.Sprintf("%d", rowCount))
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

	// Step 8: Output the matches in the desired format
	if len(matches) > 0 {
		fmt.Println(strings.Join(matches, " || "))
	} else {
		fmt.Println("No match found with any quad.")
	}
}
