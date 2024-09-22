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
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Step 2: Extract dimensions by scanning through the input
	// The first line of input is expected to contain dimensions like "1x1" or something recognizable
	firstLine := strings.TrimSpace(input.String())
	dimensions := strings.Fields(firstLine) // Assumes first line contains dimensions or relevant info

	if len(dimensions) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Unable to extract dimensions from input.")
		os.Exit(1)
	}

	// Capture dimensions from the input
	dimX := dimensions[0]
	dimY := dimensions[1]

	// Step 3: Define the quad executables to compare against
	quadExecutables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	matches := []string{} // To store the matches

	// Step 4: Loop through each quad executable and compare outputs
	for _, quad := range quadExecutables {
		// Call the quad executable with the same dimensions that were piped in
		cmd := exec.Command(quad, dimX, dimY) // Use extracted dimensions
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
