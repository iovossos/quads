package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Step 1: Read the input from stdin (the piped output from one of the quad executables)
	input := new(bytes.Buffer)
	scanner := bufio.NewScanner(os.Stdin)

	// Read the entire input
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Step 2: Capture the dimensions from the environment variable (if set by the piped executable)
	dimensions := os.Getenv("QUAD_DIMENSIONS")
	if dimensions == "" {
		fmt.Fprintln(os.Stderr, "Error: Unable to extract dimensions from the input.")
		os.Exit(1)
	}

	// Split the dimensions (should be two numbers, e.g., "1 1")
	dim := strings.Fields(dimensions)
	if len(dim) != 2 {
		fmt.Fprintln(os.Stderr, "Error: Invalid dimensions format.")
		os.Exit(1)
	}
	dimX := dim[0]
	dimY := dim[1]

	// Step 3: Define the quad executables to compare against
	quadExecutables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	pipedQuad := strings.Split(os.Args[0], "/")[1] // Get the name of the quad program piped

	matches := []string{} // To store the matches

	// Step 4: Loop through each quad executable and compare outputs
	for _, quad := range quadExecutables {
		// Skip comparing the piped quad with itself
		if quad == pipedQuad {
			continue
		}

		// Call the quad executable with the same dimensions that were piped in
		cmd := exec.Command(quad, dimX, dimY)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run %s: %v\n", quad, err)
			continue
		}

		// Compare the output of the quad executable with the piped input
		if out.String() == input.String() {
			// Add the match in the required format
			matches = append(matches, fmt.Sprintf("[%s] [%s] [%s]", quad, dimX, dimY))
		}
	}

	// Step 5: Output the matches in the desired format
	if len(matches) > 0 {
		fmt.Println(strings.Join(matches, " || "))
	} else {
		fmt.Println("No match found with any quad.")
	}
}
