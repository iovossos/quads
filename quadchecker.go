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

	// Step 2: Extract the command-line used to invoke quadC, which will contain the dimensions
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Unable to extract dimensions.")
		os.Exit(1)
	}

	// Extract the last executed command from os.Args
	// Assuming the dimensions were passed like this: ./quadC <dimX> <dimY>
	command := os.Args[0]       // The command is ./quadC
	dimX := os.Args[len(os.Args)-2] // Extract width
	dimY := os.Args[len(os.Args)-1] // Extract height

	// Step 3: Define the quad executables to compare against
	quadExecutables := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	matches := []string{} // To store the matches

	// Step 4: Loop through each quad executable and compare outputs
	for _, quad := range quadExecutables {
		// Call the quad executable with the same dimensions that were piped in
		cmd := exec.Command(quad, dimX, dimY) // Use dimensions from input
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
